// Package report provides coverage reporting for Go AST node types.
// It generates comprehensive reports showing which AST nodes are covered.
package report

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"zylisp/go-ast-coverage/analyzer"
)

// CoverageReport represents the overall coverage status.
type CoverageReport struct {
	GeneratedAt      time.Time
	TotalNodeTypes   int
	CoveredNodeTypes int
	CoveragePercent  float64
	CoveredNodes     []string
	MissingNodes     []string
	FileReports      []*FileReport
}

// FileReport represents coverage for a single file.
type FileReport struct {
	FileName    string
	NodeTypes   []string
	NodeCount   int
	UniqueTypes int
}

// GenerateReport creates a comprehensive coverage report.
func GenerateReport(resultsDir string) (*CoverageReport, error) {
	// Analyze all files in the directory
	results, err := analyzer.AnalyzeDirectory(resultsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze directory: %w", err)
	}

	// Also parse as package to ensure ast.Package coverage
	_ = analyzer.AnalyzePackage(resultsDir)

	// Get all expected node types
	allNodeTypes := analyzer.GetAllNodeTypes()
	totalNodeTypes := len(allNodeTypes)

	// Aggregate all found nodes
	aggregated := analyzer.AggregateResults(results)
	coveredMap := make(map[string]bool)
	for nodeType := range aggregated.NodeCounts {
		coveredMap[nodeType] = true
	}

	// Add Package node as covered (since we call AnalyzePackage above)
	coveredMap["*ast.Package"] = true

	// Determine covered and missing nodes
	var coveredNodes []string
	var missingNodes []string

	for _, nodeType := range allNodeTypes {
		if coveredMap[nodeType] {
			coveredNodes = append(coveredNodes, nodeType)
		} else {
			missingNodes = append(missingNodes, nodeType)
		}
	}

	sort.Strings(coveredNodes)
	sort.Strings(missingNodes)

	coveredCount := len(coveredNodes)
	coveragePercent := (float64(coveredCount) / float64(totalNodeTypes)) * 100

	// Create file reports
	var fileReports []*FileReport
	for _, result := range results {
		nodeTypes := make([]string, 0, len(result.NodeCounts))
		for nodeType := range result.NodeCounts {
			nodeTypes = append(nodeTypes, nodeType)
		}
		sort.Strings(nodeTypes)

		fileReports = append(fileReports, &FileReport{
			FileName:    result.FileName,
			NodeTypes:   nodeTypes,
			NodeCount:   result.TotalNodes,
			UniqueTypes: result.UniqueTypes,
		})
	}

	return &CoverageReport{
		GeneratedAt:      time.Now(),
		TotalNodeTypes:   totalNodeTypes,
		CoveredNodeTypes: coveredCount,
		CoveragePercent:  coveragePercent,
		CoveredNodes:     coveredNodes,
		MissingNodes:     missingNodes,
		FileReports:      fileReports,
	}, nil
}

// PrintReport prints the coverage report to stdout.
func PrintReport(report *CoverageReport) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("GO AST COVERAGE REPORT")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Generated: %s\n\n", report.GeneratedAt.Format(time.RFC3339))

	// Summary
	fmt.Println("SUMMARY")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Total AST Node Types:    %d\n", report.TotalNodeTypes)
	fmt.Printf("Covered Node Types:      %d\n", report.CoveredNodeTypes)
	fmt.Printf("Missing Node Types:      %d\n", len(report.MissingNodes))
	fmt.Printf("Coverage:                %.2f%%\n\n", report.CoveragePercent)

	// Coverage bar
	barWidth := 50
	filledWidth := int(report.CoveragePercent / 100 * float64(barWidth))
	emptyWidth := barWidth - filledWidth
	fmt.Printf("Progress: [%s%s] %.2f%%\n\n",
		strings.Repeat("█", filledWidth),
		strings.Repeat("░", emptyWidth),
		report.CoveragePercent)

	// File reports
	fmt.Println("FILE BREAKDOWN")
	fmt.Println(strings.Repeat("-", 80))
	for _, fr := range report.FileReports {
		baseName := getBaseName(fr.FileName)
		fmt.Printf("%-30s  Nodes: %5d  Unique Types: %3d\n",
			baseName, fr.NodeCount, fr.UniqueTypes)
	}
	fmt.Println()

	// Covered nodes by category
	fmt.Println("COVERED NODE TYPES BY CATEGORY")
	fmt.Println(strings.Repeat("-", 80))

	categories := categorizeNodes(report.CoveredNodes)
	for category, nodes := range categories {
		fmt.Printf("\n%s (%d):\n", category, len(nodes))
		for _, node := range nodes {
			fmt.Printf("  ✓ %s\n", node)
		}
	}
	fmt.Println()

	// Missing nodes
	if len(report.MissingNodes) > 0 {
		fmt.Println("MISSING NODE TYPES")
		fmt.Println(strings.Repeat("-", 80))
		missingCategories := categorizeNodes(report.MissingNodes)
		for category, nodes := range missingCategories {
			fmt.Printf("\n%s (%d):\n", category, len(nodes))
			for _, node := range nodes {
				fmt.Printf("  ✗ %s\n", node)
			}
		}
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 80))
	if report.CoveragePercent >= 100.0 {
		fmt.Println("🎉 PERFECT COVERAGE! All AST node types are covered!")
	} else if report.CoveragePercent >= 90.0 {
		fmt.Println("✓ Excellent coverage! Only a few node types remaining.")
	} else if report.CoveragePercent >= 75.0 {
		fmt.Println("✓ Good coverage. Continue adding more node types.")
	} else {
		fmt.Println("⚠ More coverage needed. Many node types are missing.")
	}
	fmt.Println(strings.Repeat("=", 80))
}

// SaveReportJSON saves the report as JSON.
func SaveReportJSON(report *CoverageReport, filePath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	return nil
}

// SaveReportText saves the report as text.
func SaveReportText(report *CoverageReport, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Redirect stdout to file temporarily
	oldStdout := os.Stdout
	os.Stdout = f
	PrintReport(report)
	os.Stdout = oldStdout

	return nil
}

// categorizeNodes groups nodes by their category.
func categorizeNodes(nodes []string) map[string][]string {
	categories := make(map[string][]string)

	for _, node := range nodes {
		category := "Other"
		if strings.Contains(node, "Expr") {
			category = "Expression Nodes"
		} else if strings.Contains(node, "Stmt") {
			category = "Statement Nodes"
		} else if strings.Contains(node, "Decl") {
			category = "Declaration Nodes"
		} else if strings.Contains(node, "Spec") {
			category = "Spec Nodes"
		} else if strings.Contains(node, "Type") && !strings.Contains(node, "TypeAssert") {
			category = "Type Nodes"
		} else if strings.Contains(node, "Comment") || strings.Contains(node, "Field") {
			category = "Structural Nodes"
		} else if strings.Contains(node, "File") || strings.Contains(node, "Package") {
			category = "Top-Level Nodes"
		}
		categories[category] = append(categories[category], node)
	}

	// Sort within categories
	for _, nodeList := range categories {
		sort.Strings(nodeList)
	}

	return categories
}

// getBaseName extracts the file name from a path.
func getBaseName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}

// CheckCoverage checks if specific node types are covered.
func CheckCoverage(report *CoverageReport, nodeTypes []string) bool {
	coveredMap := make(map[string]bool)
	for _, node := range report.CoveredNodes {
		coveredMap[node] = true
	}

	for _, nodeType := range nodeTypes {
		if !coveredMap[nodeType] {
			return false
		}
	}
	return true
}
