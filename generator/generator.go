// Package generator provides utilities for generating AST archive files.
// It parses Go source files and creates .asta (AST Archive) files that preserve
// source code and can reconstruct complete AST with Scope/Object references.
package generator

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"zylisp/go-ast-coverage/archive"
)

// WriteASTFiles generates AST archive files for all Go files in the input directory.
// It reads .go files from inDir and writes .asta (AST Archive) files to outDir.
func WriteASTFiles(inDir, outDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Read all files from input directory
	entries, err := os.ReadDir(inDir)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	filesProcessed := 0
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		inPath := filepath.Join(inDir, entry.Name())
		outPath := filepath.Join(outDir, strings.TrimSuffix(entry.Name(), ".go")+".asta")

		if err := generateASTFile(inPath, outPath); err != nil {
			fmt.Printf("Warning: failed to generate AST for %s: %v\n", entry.Name(), err)
			continue
		}

		filesProcessed++
		fmt.Printf("  ✓ Generated %s\n", filepath.Base(outPath))
	}

	if filesProcessed == 0 {
		return fmt.Errorf("no Go files processed")
	}

	fmt.Printf("\nGenerated %d AST files\n", filesProcessed)
	return nil
}

// generateASTFile parses a single Go file and creates an AST archive.
func generateASTFile(inPath, outPath string) error {
	// Read the source file
	source, err := os.ReadFile(inPath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Parse the source file with comments
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inPath, source, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Save AST archive with source preservation
	if err := archive.SaveASTWithSourcePreservation(file, fset, filepath.Base(inPath), outPath); err != nil {
		return fmt.Errorf("failed to create AST archive: %w", err)
	}

	return nil
}

