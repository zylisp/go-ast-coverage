// Package generator provides utilities for generating AST representation files.
// It parses Go source files and writes human-readable AST representations.
package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// WriteASTFiles generates AST files for all Go files in the input directory.
// It reads .go files from inDir and writes .ast files to outDir.
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
		outPath := filepath.Join(outDir, strings.TrimSuffix(entry.Name(), ".go")+".ast")

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

// generateASTFile parses a single Go file and writes its AST representation.
func generateASTFile(inPath, outPath string) error {
	// Parse the source file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, inPath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Generate AST representation
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("=== AST for %s ===\n\n", filepath.Base(inPath)))

	// Write file-level information
	writeFileInfo(&builder, fset, file)

	// Traverse and write AST nodes
	builder.WriteString("\n=== AST Node Tree ===\n\n")
	writeASTNode(&builder, fset, file, 0)

	// Write to output file
	if err := os.WriteFile(outPath, []byte(builder.String()), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// writeFileInfo writes high-level file information.
func writeFileInfo(builder *strings.Builder, fset *token.FileSet, file *ast.File) {
	builder.WriteString("File Information:\n")
	builder.WriteString(fmt.Sprintf("  Package: %s\n", file.Name.Name))
	builder.WriteString(fmt.Sprintf("  Position: %s\n", fset.Position(file.Pos())))

	// Imports
	if len(file.Imports) > 0 {
		builder.WriteString("  Imports:\n")
		for _, imp := range file.Imports {
			path := imp.Path.Value
			if imp.Name != nil {
				builder.WriteString(fmt.Sprintf("    %s %s\n", imp.Name.Name, path))
			} else {
				builder.WriteString(fmt.Sprintf("    %s\n", path))
			}
		}
	}

	// Comments
	if len(file.Comments) > 0 {
		builder.WriteString(fmt.Sprintf("  Comment Groups: %d\n", len(file.Comments)))
	}

	// Declarations
	builder.WriteString(fmt.Sprintf("  Declarations: %d\n", len(file.Decls)))
}

// writeASTNode recursively writes AST nodes with indentation.
func writeASTNode(builder *strings.Builder, fset *token.FileSet, node ast.Node, depth int) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", depth)
	pos := fset.Position(node.Pos())

	switch n := node.(type) {
	case *ast.File:
		builder.WriteString(fmt.Sprintf("%s*ast.File (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  Package: %s\n", indent, n.Name.Name))
		for _, decl := range n.Decls {
			writeASTNode(builder, fset, decl, depth+1)
		}

	case *ast.FuncDecl:
		funcName := n.Name.Name
		builder.WriteString(fmt.Sprintf("%s*ast.FuncDecl (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  Name: %s\n", indent, funcName))
		if n.Recv != nil {
			builder.WriteString(fmt.Sprintf("%s  Receiver:\n", indent))
			writeASTNode(builder, fset, n.Recv, depth+2)
		}
		if n.Type != nil {
			builder.WriteString(fmt.Sprintf("%s  Type:\n", indent))
			writeASTNode(builder, fset, n.Type, depth+2)
		}
		if n.Body != nil {
			builder.WriteString(fmt.Sprintf("%s  Body:\n", indent))
			writeASTNode(builder, fset, n.Body, depth+2)
		}

	case *ast.GenDecl:
		builder.WriteString(fmt.Sprintf("%s*ast.GenDecl (%s) Token: %s\n", indent, pos, n.Tok))
		for _, spec := range n.Specs {
			writeASTNode(builder, fset, spec, depth+1)
		}

	case *ast.ImportSpec:
		path := n.Path.Value
		builder.WriteString(fmt.Sprintf("%s*ast.ImportSpec (%s)\n", indent, pos))
		if n.Name != nil {
			builder.WriteString(fmt.Sprintf("%s  Name: %s\n", indent, n.Name.Name))
		}
		builder.WriteString(fmt.Sprintf("%s  Path: %s\n", indent, path))

	case *ast.TypeSpec:
		builder.WriteString(fmt.Sprintf("%s*ast.TypeSpec (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  Name: %s\n", indent, n.Name.Name))
		builder.WriteString(fmt.Sprintf("%s  Type:\n", indent))
		writeASTNode(builder, fset, n.Type, depth+2)

	case *ast.ValueSpec:
		builder.WriteString(fmt.Sprintf("%s*ast.ValueSpec (%s)\n", indent, pos))
		for _, name := range n.Names {
			builder.WriteString(fmt.Sprintf("%s  Name: %s\n", indent, name.Name))
		}
		if n.Type != nil {
			builder.WriteString(fmt.Sprintf("%s  Type:\n", indent))
			writeASTNode(builder, fset, n.Type, depth+2)
		}
		for i, val := range n.Values {
			builder.WriteString(fmt.Sprintf("%s  Value[%d]:\n", indent, i))
			writeASTNode(builder, fset, val, depth+2)
		}

	case *ast.BlockStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.BlockStmt (%s) Stmts: %d\n", indent, pos, len(n.List)))
		for _, stmt := range n.List {
			writeASTNode(builder, fset, stmt, depth+1)
		}

	case *ast.ExprStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.ExprStmt (%s)\n", indent, pos))
		writeASTNode(builder, fset, n.X, depth+1)

	case *ast.AssignStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.AssignStmt (%s) Token: %s\n", indent, pos, n.Tok))
		builder.WriteString(fmt.Sprintf("%s  LHS:\n", indent))
		for _, lhs := range n.Lhs {
			writeASTNode(builder, fset, lhs, depth+2)
		}
		builder.WriteString(fmt.Sprintf("%s  RHS:\n", indent))
		for _, rhs := range n.Rhs {
			writeASTNode(builder, fset, rhs, depth+2)
		}

	case *ast.DeclStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.DeclStmt (%s)\n", indent, pos))
		writeASTNode(builder, fset, n.Decl, depth+1)

	case *ast.ReturnStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.ReturnStmt (%s)\n", indent, pos))
		for i, result := range n.Results {
			builder.WriteString(fmt.Sprintf("%s  Result[%d]:\n", indent, i))
			writeASTNode(builder, fset, result, depth+2)
		}

	case *ast.IfStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.IfStmt (%s)\n", indent, pos))
		if n.Init != nil {
			builder.WriteString(fmt.Sprintf("%s  Init:\n", indent))
			writeASTNode(builder, fset, n.Init, depth+2)
		}
		builder.WriteString(fmt.Sprintf("%s  Cond:\n", indent))
		writeASTNode(builder, fset, n.Cond, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Body:\n", indent))
		writeASTNode(builder, fset, n.Body, depth+2)
		if n.Else != nil {
			builder.WriteString(fmt.Sprintf("%s  Else:\n", indent))
			writeASTNode(builder, fset, n.Else, depth+2)
		}

	case *ast.ForStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.ForStmt (%s)\n", indent, pos))
		if n.Init != nil {
			builder.WriteString(fmt.Sprintf("%s  Init:\n", indent))
			writeASTNode(builder, fset, n.Init, depth+2)
		}
		if n.Cond != nil {
			builder.WriteString(fmt.Sprintf("%s  Cond:\n", indent))
			writeASTNode(builder, fset, n.Cond, depth+2)
		}
		if n.Post != nil {
			builder.WriteString(fmt.Sprintf("%s  Post:\n", indent))
			writeASTNode(builder, fset, n.Post, depth+2)
		}
		builder.WriteString(fmt.Sprintf("%s  Body:\n", indent))
		writeASTNode(builder, fset, n.Body, depth+2)

	case *ast.RangeStmt:
		builder.WriteString(fmt.Sprintf("%s*ast.RangeStmt (%s) Token: %s\n", indent, pos, n.Tok))
		if n.Key != nil {
			builder.WriteString(fmt.Sprintf("%s  Key:\n", indent))
			writeASTNode(builder, fset, n.Key, depth+2)
		}
		if n.Value != nil {
			builder.WriteString(fmt.Sprintf("%s  Value:\n", indent))
			writeASTNode(builder, fset, n.Value, depth+2)
		}
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Body:\n", indent))
		writeASTNode(builder, fset, n.Body, depth+2)

	case *ast.CallExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.CallExpr (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  Fun:\n", indent))
		writeASTNode(builder, fset, n.Fun, depth+2)
		if len(n.Args) > 0 {
			builder.WriteString(fmt.Sprintf("%s  Args:\n", indent))
			for i, arg := range n.Args {
				builder.WriteString(fmt.Sprintf("%s    [%d]:\n", indent, i))
				writeASTNode(builder, fset, arg, depth+3)
			}
		}

	case *ast.BinaryExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.BinaryExpr (%s) Op: %s\n", indent, pos, n.Op))
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Y:\n", indent))
		writeASTNode(builder, fset, n.Y, depth+2)

	case *ast.UnaryExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.UnaryExpr (%s) Op: %s\n", indent, pos, n.Op))
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)

	case *ast.SelectorExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.SelectorExpr (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Sel: %s\n", indent, n.Sel.Name))

	case *ast.IndexExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.IndexExpr (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Index:\n", indent))
		writeASTNode(builder, fset, n.Index, depth+2)

	case *ast.SliceExpr:
		builder.WriteString(fmt.Sprintf("%s*ast.SliceExpr (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  X:\n", indent))
		writeASTNode(builder, fset, n.X, depth+2)
		if n.Low != nil {
			builder.WriteString(fmt.Sprintf("%s  Low:\n", indent))
			writeASTNode(builder, fset, n.Low, depth+2)
		}
		if n.High != nil {
			builder.WriteString(fmt.Sprintf("%s  High:\n", indent))
			writeASTNode(builder, fset, n.High, depth+2)
		}
		if n.Max != nil {
			builder.WriteString(fmt.Sprintf("%s  Max:\n", indent))
			writeASTNode(builder, fset, n.Max, depth+2)
		}

	case *ast.BasicLit:
		builder.WriteString(fmt.Sprintf("%s*ast.BasicLit (%s) Kind: %s Value: %s\n", indent, pos, n.Kind, n.Value))

	case *ast.Ident:
		builder.WriteString(fmt.Sprintf("%s*ast.Ident (%s) Name: %s\n", indent, pos, n.Name))

	case *ast.CompositeLit:
		builder.WriteString(fmt.Sprintf("%s*ast.CompositeLit (%s)\n", indent, pos))
		if n.Type != nil {
			builder.WriteString(fmt.Sprintf("%s  Type:\n", indent))
			writeASTNode(builder, fset, n.Type, depth+2)
		}
		if len(n.Elts) > 0 {
			builder.WriteString(fmt.Sprintf("%s  Elts:\n", indent))
			for i, elt := range n.Elts {
				builder.WriteString(fmt.Sprintf("%s    [%d]:\n", indent, i))
				writeASTNode(builder, fset, elt, depth+3)
			}
		}

	case *ast.FuncType:
		builder.WriteString(fmt.Sprintf("%s*ast.FuncType (%s)\n", indent, pos))
		if n.Params != nil {
			builder.WriteString(fmt.Sprintf("%s  Params:\n", indent))
			writeASTNode(builder, fset, n.Params, depth+2)
		}
		if n.Results != nil {
			builder.WriteString(fmt.Sprintf("%s  Results:\n", indent))
			writeASTNode(builder, fset, n.Results, depth+2)
		}

	case *ast.FieldList:
		builder.WriteString(fmt.Sprintf("%s*ast.FieldList (%s) NumFields: %d\n", indent, pos, n.NumFields()))
		for i, field := range n.List {
			builder.WriteString(fmt.Sprintf("%s  Field[%d]:\n", indent, i))
			writeASTNode(builder, fset, field, depth+2)
		}

	case *ast.Field:
		builder.WriteString(fmt.Sprintf("%s*ast.Field (%s)\n", indent, pos))
		for _, name := range n.Names {
			builder.WriteString(fmt.Sprintf("%s  Name: %s\n", indent, name.Name))
		}
		if n.Type != nil {
			builder.WriteString(fmt.Sprintf("%s  Type:\n", indent))
			writeASTNode(builder, fset, n.Type, depth+2)
		}

	case *ast.StructType:
		builder.WriteString(fmt.Sprintf("%s*ast.StructType (%s)\n", indent, pos))
		if n.Fields != nil {
			writeASTNode(builder, fset, n.Fields, depth+1)
		}

	case *ast.InterfaceType:
		builder.WriteString(fmt.Sprintf("%s*ast.InterfaceType (%s)\n", indent, pos))
		if n.Methods != nil {
			writeASTNode(builder, fset, n.Methods, depth+1)
		}

	case *ast.ArrayType:
		builder.WriteString(fmt.Sprintf("%s*ast.ArrayType (%s)\n", indent, pos))
		if n.Len != nil {
			builder.WriteString(fmt.Sprintf("%s  Len:\n", indent))
			writeASTNode(builder, fset, n.Len, depth+2)
		}
		builder.WriteString(fmt.Sprintf("%s  Elt:\n", indent))
		writeASTNode(builder, fset, n.Elt, depth+2)

	case *ast.MapType:
		builder.WriteString(fmt.Sprintf("%s*ast.MapType (%s)\n", indent, pos))
		builder.WriteString(fmt.Sprintf("%s  Key:\n", indent))
		writeASTNode(builder, fset, n.Key, depth+2)
		builder.WriteString(fmt.Sprintf("%s  Value:\n", indent))
		writeASTNode(builder, fset, n.Value, depth+2)

	case *ast.ChanType:
		builder.WriteString(fmt.Sprintf("%s*ast.ChanType (%s) Dir: %s\n", indent, pos, n.Dir))
		builder.WriteString(fmt.Sprintf("%s  Value:\n", indent))
		writeASTNode(builder, fset, n.Value, depth+2)

	default:
		// Generic fallback for any node type not explicitly handled
		builder.WriteString(fmt.Sprintf("%s%T (%s)\n", indent, node, pos))
	}
}
