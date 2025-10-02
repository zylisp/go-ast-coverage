// Package analyzer provides utilities for analyzing Go AST nodes.
// It can parse Go source files and enumerate all AST node types found.
package analyzer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"sort"
)

// NodeCount tracks the count of each AST node type.
type NodeCount struct {
	Type  string
	Count int
}

// AnalysisResult contains the results of AST analysis.
type AnalysisResult struct {
	FileName   string
	NodeCounts map[string]int
	TotalNodes int
	UniqueTypes int
}

// AnalyzeFile parses a Go source file and returns analysis results.
func AnalyzeFile(filePath string) (*AnalysisResult, error) {
	// Read the file
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}

	// Count nodes
	nodeCounts := make(map[string]int)
	totalNodes := 0

	ast.Inspect(file, func(n ast.Node) bool {
		if n != nil {
			nodeType := fmt.Sprintf("%T", n)
			nodeCounts[nodeType]++
			totalNodes++
		}
		return true
	})

	return &AnalysisResult{
		FileName:    filePath,
		NodeCounts:  nodeCounts,
		TotalNodes:  totalNodes,
		UniqueTypes: len(nodeCounts),
	}, nil
}

// PrintAnalysis prints the analysis results in a human-readable format.
func PrintAnalysis(result *AnalysisResult) {
	fmt.Printf("\n=== AST Analysis: %s ===\n", result.FileName)
	fmt.Printf("Total nodes: %d\n", result.TotalNodes)
	fmt.Printf("Unique node types: %d\n\n", result.UniqueTypes)

	// Sort by node type name
	var counts []NodeCount
	for nodeType, count := range result.NodeCounts {
		counts = append(counts, NodeCount{Type: nodeType, Count: count})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Type < counts[j].Type
	})

	fmt.Println("Node type distribution:")
	for _, nc := range counts {
		fmt.Printf("  %-40s %5d\n", nc.Type, nc.Count)
	}
	fmt.Println("========================================")
}

// GetAllNodeTypes returns a list of all AST node types defined in go/ast package.
func GetAllNodeTypes() []string {
	// This is a comprehensive list of all AST node types
	return []string{
		// Expression nodes
		"*ast.BadExpr",
		"*ast.Ident",
		"*ast.Ellipsis",
		"*ast.BasicLit",
		"*ast.FuncLit",
		"*ast.CompositeLit",
		"*ast.ParenExpr",
		"*ast.SelectorExpr",
		"*ast.IndexExpr",
		"*ast.IndexListExpr", // Go 1.18+ generics
		"*ast.SliceExpr",
		"*ast.TypeAssertExpr",
		"*ast.CallExpr",
		"*ast.StarExpr",
		"*ast.UnaryExpr",
		"*ast.BinaryExpr",
		"*ast.KeyValueExpr",

		// Statement nodes
		"*ast.BadStmt",
		"*ast.DeclStmt",
		"*ast.EmptyStmt",
		"*ast.LabeledStmt",
		"*ast.ExprStmt",
		"*ast.SendStmt",
		"*ast.IncDecStmt",
		"*ast.AssignStmt",
		"*ast.GoStmt",
		"*ast.DeferStmt",
		"*ast.ReturnStmt",
		"*ast.BranchStmt",
		"*ast.BlockStmt",
		"*ast.IfStmt",
		"*ast.CaseClause",
		"*ast.SwitchStmt",
		"*ast.TypeSwitchStmt",
		"*ast.CommClause",
		"*ast.SelectStmt",
		"*ast.ForStmt",
		"*ast.RangeStmt",

		// Declaration nodes
		"*ast.BadDecl",
		"*ast.GenDecl",
		"*ast.FuncDecl",

		// Spec nodes
		"*ast.ImportSpec",
		"*ast.ValueSpec",
		"*ast.TypeSpec",

		// Other important nodes
		"*ast.File",
		"*ast.Package",
		"*ast.Comment",
		"*ast.CommentGroup",
		"*ast.Field",
		"*ast.FieldList",

		// Type nodes
		"*ast.ArrayType",
		"*ast.StructType",
		"*ast.FuncType",
		"*ast.InterfaceType",
		"*ast.MapType",
		"*ast.ChanType",
	}
}

// CompareWithExpected compares analysis results with expected node types.
func CompareWithExpected(result *AnalysisResult, expectedTypes []string) {
	fmt.Printf("\n=== Coverage Check: %s ===\n", result.FileName)

	found := make(map[string]bool)
	for nodeType := range result.NodeCounts {
		found[nodeType] = true
	}

	missing := []string{}
	for _, expected := range expectedTypes {
		if !found[expected] {
			missing = append(missing, expected)
		}
	}

	if len(missing) == 0 {
		fmt.Println("✓ All expected node types found!")
	} else {
		fmt.Printf("⚠ Missing %d node types:\n", len(missing))
		for _, nodeType := range missing {
			fmt.Printf("  - %s\n", nodeType)
		}
	}
	fmt.Println("========================================")
}

// AnalyzeDirectory analyzes all Go files in a directory.
func AnalyzeDirectory(dirPath string) ([]*AnalysisResult, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var results []*AnalysisResult
	for _, entry := range entries {
		if !entry.IsDir() && len(entry.Name()) > 3 && entry.Name()[len(entry.Name())-3:] == ".go" {
			filePath := dirPath + "/" + entry.Name()
			result, err := AnalyzeFile(filePath)
			if err != nil {
				fmt.Printf("Warning: failed to analyze %s: %v\n", filePath, err)
				continue
			}
			results = append(results, result)
		}
	}

	return results, nil
}

// AggregateResults combines multiple analysis results into one.
func AggregateResults(results []*AnalysisResult) *AnalysisResult {
	aggregated := &AnalysisResult{
		FileName:   "Aggregated",
		NodeCounts: make(map[string]int),
	}

	for _, result := range results {
		for nodeType, count := range result.NodeCounts {
			aggregated.NodeCounts[nodeType] += count
			aggregated.TotalNodes += count
		}
	}

	aggregated.UniqueTypes = len(aggregated.NodeCounts)
	return aggregated
}

// GetNodeTypeName returns the AST node type name for any node.
func GetNodeTypeName(n ast.Node) string {
	if n == nil {
		return "<nil>"
	}
	return reflect.TypeOf(n).String()
}
