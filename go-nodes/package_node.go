// Package main demonstrates the ast.Package node.
// The ast.Package node is created when parsing an entire directory of Go files
// as a package using parser.ParseDir(), not parser.ParseFile().
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// AST Nodes Covered:
// - *ast.Package - Created by parser.ParseDir() when parsing a directory

func main() {
	fmt.Println("=== package_node.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")
	fmt.Println()

	// The ast.Package node represents a Go package consisting of multiple files
	fmt.Println("ast.Package Node:")
	fmt.Println("  The *ast.Package type represents a set of source files")
	fmt.Println("  that collectively form a Go package.")
	fmt.Println()

	// Demonstrate creating an ast.Package by parsing the current directory
	fmt.Println("  ✓ Creating ast.Package by parsing directory:")

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("    Error getting working directory: %v\n", err)
		return
	}

	// Parse the directory as a package
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("    Error parsing directory: %v\n", err)
		return
	}

	// Display information about parsed packages
	for pkgName, pkg := range pkgs {
		// This demonstrates the ast.Package node
		nodeType := fmt.Sprintf("%T", pkg)
		fmt.Printf("    Package node type: %s\n", nodeType)

		// Verify it's actually an *ast.Package
		if _, ok := interface{}(pkg).(*ast.Package); ok {
			fmt.Println("    ✓ Confirmed: This is an *ast.Package node")
		}
		fmt.Printf("    Package name: %s\n", pkgName)
		fmt.Printf("    Number of files: %d\n", len(pkg.Files))

		// Show files in the package
		fmt.Println("    Files in package:")
		for fileName := range pkg.Files {
			fmt.Printf("      - %s\n", fileName)
		}
		fmt.Println()

		// ast.Package structure:
		// type Package struct {
		//     Name    string             // package name
		//     Scope   *Scope             // package scope across all files
		//     Imports map[string]*Object // map of package id -> package object
		//     Files   map[string]*File   // Go source files by filename
		// }

		// Access package scope
		if pkg.Scope != nil {
			fmt.Printf("    Package scope contains %d objects\n", len(pkg.Scope.Objects))
		}

		// Access imports
		fmt.Printf("    Package imports: %d packages\n", len(pkg.Imports))
	}

	fmt.Println("\nKey Points:")
	fmt.Println("  • ast.Package represents an entire package (multiple files)")
	fmt.Println("  • Created by parser.ParseDir(), not parser.ParseFile()")
	fmt.Println("  • Contains a map of filenames to *ast.File nodes")
	fmt.Println("  • Has a unified scope across all files in the package")
	fmt.Println("  • Tracks imports at the package level")
	fmt.Println()
	fmt.Println("Summary: *ast.Package node exercised")
	fmt.Println("========================================")
}
