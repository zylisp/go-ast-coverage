// Package main is the orchestrator for the Go AST coverage test suite.
// It runs all test files and generates comprehensive coverage reports.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"zylisp/go-ast-coverage/ast-analyzer"
	"zylisp/go-ast-coverage/ast-generator"
	"zylisp/go-ast-coverage/coverage-report"
)

// Configuration flags
var (
	runTests       = flag.Bool("run", false, "Run all test files")
	analyze        = flag.Bool("analyze", false, "Analyze AST nodes in test files")
	generateReport = flag.Bool("report", false, "Generate coverage report")
	generateAST    = flag.Bool("generate", false, "Generate AST files from go-nodes")
	saveJSON       = flag.Bool("json", false, "Save report as JSON")
	verbose        = flag.Bool("verbose", false, "Verbose output")
	all            = flag.Bool("all", false, "Run all tests, analyze, and generate report")
)

func main() {
	flag.Parse()

	// If no flags, default to all
	if !*runTests && !*analyze && !*generateReport && !*all {
		*all = true
	}

	if *all {
		*runTests = true
		*analyze = true
		*generateReport = true
	}

	fmt.Println("=== Go AST Coverage Test Suite ===")
	fmt.Println()

	astNodesDir := "go-nodes"

	// Run test files
	if *runTests {
		fmt.Println("Running test files...")
		if err := runTestFiles(astNodesDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error running tests: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}

	// Analyze AST nodes
	if *analyze {
		fmt.Println("Analyzing AST nodes...")
		if err := analyzeFiles(astNodesDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error analyzing files: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}

	// Generate AST files
	if *generateAST {
		fmt.Println("Generating AST files...")
		if err := generateASTFiles(astNodesDir, "ast-nodes"); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating AST files: %v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}

	// Generate coverage report
	if *generateReport {
		fmt.Println("Generating coverage report...")
		if err := generateCoverageReport(astNodesDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating report: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("\n✓ All tasks completed successfully!")
}

// runTestFiles executes all Go files in the ast-nodes directory.
func runTestFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	executedCount := 0
	failedCount := 0

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		fmt.Printf("Running %s...\n", file.Name())

		cmd := exec.Command("go", "run", filePath)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("  ✗ FAILED: %v\n", err)
			if *verbose {
				fmt.Printf("Output:\n%s\n", string(output))
			}
			failedCount++
		} else {
			if *verbose {
				fmt.Printf("Output:\n%s\n", string(output))
			} else {
				fmt.Printf("  ✓ Success\n")
			}
			executedCount++
		}
	}

	fmt.Printf("\nExecution Summary: %d succeeded, %d failed\n", executedCount, failedCount)

	if failedCount > 0 {
		return fmt.Errorf("%d file(s) failed to execute", failedCount)
	}

	return nil
}

// analyzeFiles analyzes all Go files and prints AST statistics.
func analyzeFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	var allResults []*analyzer.AnalysisResult

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		result, err := analyzer.AnalyzeFile(filePath)
		if err != nil {
			fmt.Printf("Warning: failed to analyze %s: %v\n", file.Name(), err)
			continue
		}

		if *verbose {
			analyzer.PrintAnalysis(result)
		}

		allResults = append(allResults, result)
	}

	// Print aggregated statistics
	if len(allResults) > 0 {
		aggregated := analyzer.AggregateResults(allResults)
		fmt.Println("\n=== Aggregated Statistics ===")
		fmt.Printf("Total files analyzed: %d\n", len(allResults))
		fmt.Printf("Total AST nodes: %d\n", aggregated.TotalNodes)
		fmt.Printf("Unique node types: %d\n", aggregated.UniqueTypes)
		fmt.Println()

		// Show top 10 most common node types
		type nodeCount struct {
			nodeType string
			count    int
		}
		var counts []nodeCount
		for nodeType, count := range aggregated.NodeCounts {
			counts = append(counts, nodeCount{nodeType, count})
		}

		// Sort by count (descending)
		for i := 0; i < len(counts); i++ {
			for j := i + 1; j < len(counts); j++ {
				if counts[j].count > counts[i].count {
					counts[i], counts[j] = counts[j], counts[i]
				}
			}
		}

		fmt.Println("Top 10 most common node types:")
		for i := 0; i < 10 && i < len(counts); i++ {
			fmt.Printf("  %d. %-40s %5d\n", i+1, counts[i].nodeType, counts[i].count)
		}
		fmt.Println()
	}

	// Parse directory as package to exercise ast.Package node
	fmt.Println("Analyzing directory as package for ast.Package coverage:")
	if err := analyzer.AnalyzePackage(dir); err != nil {
		fmt.Printf("Warning: failed to analyze package: %v\n", err)
	}
	fmt.Println()

	return nil
}

// generateCoverageReport generates and displays the coverage report.
func generateCoverageReport(dir string) error {
	rep, err := report.GenerateReport(dir)
	if err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	// Print report to stdout
	report.PrintReport(rep)

	// Save JSON if requested
	if *saveJSON {
		jsonPath := "coverage-report.json"
		if err := report.SaveReportJSON(rep, jsonPath); err != nil {
			fmt.Printf("Warning: failed to save JSON report: %v\n", err)
		} else {
			fmt.Printf("\n✓ JSON report saved to: %s\n", jsonPath)
		}
	}

	// Save text report
	textPath := "coverage-report.txt"
	if err := report.SaveReportText(rep, textPath); err != nil {
		fmt.Printf("Warning: failed to save text report: %v\n", err)
	} else {
		fmt.Printf("✓ Text report saved to: %s\n", textPath)
	}

	return nil
}

// generateASTFiles generates AST representation files from Go source files.
func generateASTFiles(inDir, outDir string) error {
	if err := generator.WriteASTFiles(inDir, outDir); err != nil {
		return fmt.Errorf("failed to generate AST files: %w", err)
	}
	fmt.Printf("✓ AST files written to: %s\n", outDir)
	return nil
}
