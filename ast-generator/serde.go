package generator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"reflect"
)

// ASTBundle contains everything needed to perfectly reconstruct an AST
type ASTBundle struct {
	FileSet *token.FileSet `gob:"fileset"`
	File    *ast.File      `gob:"file"`
	// Optional: store the original source for verification
	OriginalSource string `gob:"source,omitempty"`
}

// RegisterAllASTTypes registers all AST types with gob for serialization
// This is crucial for preserving exact type information
func RegisterAllASTTypes() {
	// Core interfaces - these must be registered first
	gob.Register((*ast.Node)(nil))
	gob.Register((*ast.Expr)(nil))
	gob.Register((*ast.Stmt)(nil))
	gob.Register((*ast.Decl)(nil))
	gob.Register((*ast.Spec)(nil))

	// Expression types
	gob.Register(&ast.BadExpr{})
	gob.Register(&ast.Ident{})
	gob.Register(&ast.Ellipsis{})
	gob.Register(&ast.BasicLit{})
	gob.Register(&ast.FuncLit{})
	gob.Register(&ast.CompositeLit{})
	gob.Register(&ast.ParenExpr{})
	gob.Register(&ast.SelectorExpr{})
	gob.Register(&ast.IndexExpr{})
	gob.Register(&ast.IndexListExpr{}) // Go 1.18+
	gob.Register(&ast.SliceExpr{})
	gob.Register(&ast.TypeAssertExpr{})
	gob.Register(&ast.CallExpr{})
	gob.Register(&ast.StarExpr{})
	gob.Register(&ast.UnaryExpr{})
	gob.Register(&ast.BinaryExpr{})
	gob.Register(&ast.KeyValueExpr{})

	// Statement types
	gob.Register(&ast.BadStmt{})
	gob.Register(&ast.DeclStmt{})
	gob.Register(&ast.EmptyStmt{})
	gob.Register(&ast.LabeledStmt{})
	gob.Register(&ast.ExprStmt{})
	gob.Register(&ast.SendStmt{})
	gob.Register(&ast.IncDecStmt{})
	gob.Register(&ast.AssignStmt{})
	gob.Register(&ast.GoStmt{})
	gob.Register(&ast.DeferStmt{})
	gob.Register(&ast.ReturnStmt{})
	gob.Register(&ast.BranchStmt{})
	gob.Register(&ast.BlockStmt{})
	gob.Register(&ast.IfStmt{})
	gob.Register(&ast.CaseClause{})
	gob.Register(&ast.SwitchStmt{})
	gob.Register(&ast.TypeSwitchStmt{})
	gob.Register(&ast.CommClause{})
	gob.Register(&ast.SelectStmt{})
	gob.Register(&ast.ForStmt{})
	gob.Register(&ast.RangeStmt{})

	// Declaration types
	gob.Register(&ast.BadDecl{})
	gob.Register(&ast.GenDecl{})
	gob.Register(&ast.FuncDecl{})

	// Spec types
	gob.Register(&ast.ImportSpec{})
	gob.Register(&ast.ValueSpec{})
	gob.Register(&ast.TypeSpec{})

	// Type expression types
	gob.Register(&ast.ArrayType{})
	gob.Register(&ast.StructType{})
	gob.Register(&ast.FuncType{})
	gob.Register(&ast.InterfaceType{})
	gob.Register(&ast.MapType{})
	gob.Register(&ast.ChanType{})

	// Other important types
	gob.Register(&ast.Field{})
	gob.Register(&ast.FieldList{})
	gob.Register(&ast.File{})
	gob.Register(&ast.Package{})
	gob.Register(&ast.Comment{})
	gob.Register(&ast.CommentGroup{})
	gob.Register(&ast.Scope{})
	gob.Register(&ast.Object{})

	// Token types
	gob.Register(token.Token(0))
	gob.Register(token.Pos(0))

	// Slices of interfaces (these are crucial!)
	gob.Register([]ast.Expr{})
	gob.Register([]ast.Stmt{})
	gob.Register([]ast.Decl{})
	gob.Register([]ast.Spec{})
	gob.Register([]*ast.Ident{})
	gob.Register([]*ast.Field{})
	gob.Register([]*ast.Comment{})
	gob.Register([]*ast.CommentGroup{})
	gob.Register([]*ast.ImportSpec{})

	// Maps with interface values
	gob.Register(map[string]*ast.Object{})
	gob.Register(map[string]*ast.File{})

	// Object kinds
	gob.Register(ast.ObjKind(0))
}

// SaveASTWithPerfectFidelity saves an AST with complete fidelity
func SaveASTWithPerfectFidelity(file *ast.File, fset *token.FileSet, originalSource string, filename string) error {
	// Ensure all types are registered
	RegisterAllASTTypes()

	bundle := ASTBundle{
		FileSet:        fset,
		File:           file,
		OriginalSource: originalSource,
	}

	// Use a buffer for better error handling
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	if err := encoder.Encode(&bundle); err != nil {
		return fmt.Errorf("failed to encode AST bundle: %w", err)
	}

	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// LoadASTWithPerfectFidelity loads an AST with complete fidelity
func LoadASTWithPerfectFidelity(filename string) (*ast.File, *token.FileSet, string, error) {
	// Ensure all types are registered
	RegisterAllASTTypes()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to read file: %w", err)
	}

	var bundle ASTBundle
	decoder := gob.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&bundle); err != nil {
		return nil, nil, "", fmt.Errorf("failed to decode AST bundle: %w", err)
	}

	return bundle.File, bundle.FileSet, bundle.OriginalSource, nil
}

// VerifyASTFidelity compares two AST structures for exact equality
func VerifyASTFidelity(original, restored *ast.File, originalFset, restoredFset *token.FileSet) error {
	// Check if we can format both ASTs back to source code
	var originalBuf, restoredBuf bytes.Buffer

	if err := format.Node(&originalBuf, originalFset, original); err != nil {
		return fmt.Errorf("failed to format original AST: %w", err)
	}

	if err := format.Node(&restoredBuf, restoredFset, restored); err != nil {
		return fmt.Errorf("failed to format restored AST: %w", err)
	}

	originalCode := originalBuf.String()
	restoredCode := restoredBuf.String()

	if originalCode != restoredCode {
		return fmt.Errorf("AST fidelity check failed:\nOriginal:\n%s\nRestored:\n%s",
			originalCode, restoredCode)
	}

	return nil
}

// DeepCompareAST performs a deep structural comparison of AST nodes
func DeepCompareAST(a, b ast.Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Compare types
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	// Use reflection to compare all fields
	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
	if va.Kind() == reflect.Ptr {
		va = va.Elem()
	}
	if vb.Kind() == reflect.Ptr {
		vb = vb.Elem()
	}

	for i := 0; i < va.NumField(); i++ {
		fieldA, fieldB := va.Field(i), vb.Field(i)

		// Skip unexported fields and function fields
		if !va.Type().Field(i).IsExported() {
			continue
		}

		if !compareField(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	return true
}

func compareField(a, b interface{}) bool {
	// Handle different types of fields
	switch va := a.(type) {
	case ast.Node:
		if vb, ok := b.(ast.Node); ok {
			return DeepCompareAST(va, vb)
		}
		return false
	case []ast.Expr:
		if vb, ok := b.([]ast.Expr); ok {
			if len(va) != len(vb) {
				return false
			}
			for i := range va {
				if !DeepCompareAST(va[i], vb[i]) {
					return false
				}
			}
			return true
		}
		return false
	case []ast.Stmt:
		if vb, ok := b.([]ast.Stmt); ok {
			if len(va) != len(vb) {
				return false
			}
			for i := range va {
				if !DeepCompareAST(va[i], vb[i]) {
					return false
				}
			}
			return true
		}
		return false
	case []ast.Decl:
		if vb, ok := b.([]ast.Decl); ok {
			if len(va) != len(vb) {
				return false
			}
			for i := range va {
				if !DeepCompareAST(va[i], vb[i]) {
					return false
				}
			}
			return true
		}
		return false
	default:
		// For primitive types, use reflect.DeepEqual
		return reflect.DeepEqual(a, b)
	}
}

// ComprehensiveASTTest performs a thorough test of AST serialization fidelity
func ComprehensiveASTTest() error {
	// Test source with complex structures
	source := `package main

import (
	"fmt"
	"context"
	"go/ast"
)

// TestStruct demonstrates various Go constructs
type TestStruct[T any] struct {
	Name string ` + "`json:\"name\"`" + `
	Value T
	*EmbeddedStruct
}

type EmbeddedStruct struct {
	ID int
}

// TestInterface with type constraints
type TestInterface[T comparable] interface {
	Method(ctx context.Context, param T) (T, error)
	~string | ~int
}

func main() {
	// Various expression types
	x := 42
	y := &x
	z := *y + 10

	// Composite literals
	slice := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2}

	// Control flow
	for i, v := range slice {
		switch {
		case v > 2:
			fmt.Printf("Index %d: %d is big\n", i, v)
		default:
			if v < 0 {
				continue
			}
			fmt.Printf("Index %d: %d is small\n", i, v)
		}
	}

	// Channels and goroutines
	ch := make(chan int, 1)
	go func() {
		defer close(ch)
		ch <- z
	}()

	select {
	case result := <-ch:
		fmt.Println("Result:", result)
	default:
		fmt.Println("No result")
	}

	// Function literal
	fn := func(a, b int) int { return a + b }
	_ = fn(1, 2)

	// Type assertion
	var i interface{} = "hello"
	if s, ok := i.(string); ok {
		fmt.Println("String:", s)
	}
}`

	fmt.Println("Testing comprehensive AST serialization...")

	// Parse the source
	fset := token.NewFileSet()
	originalFile, err := parser.ParseFile(fset, "test.go", source, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse source: %w", err)
	}

	fmt.Printf("Original AST parsed successfully. Package: %s\n", originalFile.Name.Name)

	// Save with perfect fidelity
	filename := "comprehensive_ast.gob"
	if err := SaveASTWithPerfectFidelity(originalFile, fset, source, filename); err != nil {
		return fmt.Errorf("failed to save AST: %w", err)
	}
	fmt.Printf("AST saved to %s\n", filename)

	// Load back
	restoredFile, restoredFset, restoredSource, err := LoadASTWithPerfectFidelity(filename)
	if err != nil {
		return fmt.Errorf("failed to load AST: %w", err)
	}
	fmt.Printf("AST loaded successfully. Package: %s\n", restoredFile.Name.Name)

	// Verify source code fidelity
	if restoredSource != source {
		return fmt.Errorf("source code fidelity check failed")
	}
	fmt.Println("✓ Source code fidelity verified")

	// Verify AST structural fidelity
	if err := VerifyASTFidelity(originalFile, restoredFile, fset, restoredFset); err != nil {
		return fmt.Errorf("AST fidelity verification failed: %w", err)
	}
	fmt.Println("✓ AST formatting fidelity verified")

	// Deep structural comparison
	if !DeepCompareAST(originalFile, restoredFile) {
		return fmt.Errorf("deep AST comparison failed")
	}
	fmt.Println("✓ Deep AST structural comparison passed")

	// Check specific node types
	nodeCount := 0
	ast.Inspect(restoredFile, func(n ast.Node) bool {
		if n != nil {
			nodeCount++
		}
		return true
	})
	fmt.Printf("✓ Restored AST contains %d nodes\n", nodeCount)

	fmt.Println("All fidelity tests passed!")
	return nil
}
