package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"testing"
)

// TestSimpleASTArchive tests basic AST archiving and restoration
func TestSimpleASTArchive(t *testing.T) {
	source := `package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
}
`

	// Cleanup
	defer os.Remove("test_simple.asta")

	// Parse original
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "simple.go", source, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse source: %v", err)
	}

	// Save
	if err := SaveASTWithSourcePreservation(file, fset, "simple.go", "test_simple.asta"); err != nil {
		t.Fatalf("failed to save AST: %v", err)
	}

	// Load
	restoredFile, restoredFset, restoredSource, err := LoadASTWithSourceReconstruction("test_simple.asta")
	if err != nil {
		t.Fatalf("failed to load AST: %v", err)
	}

	// Verify
	if err := VerifyPerfectFidelity(file, restoredFile, fset, restoredFset); err != nil {
		t.Errorf("fidelity check failed: %v", err)
	}

	// Verify source preserved
	if restoredSource != source {
		t.Errorf("source code mismatch")
	}
}

// TestComplexASTWithScopes tests AST archiving with complex scoping
func TestComplexASTWithScopes(t *testing.T) {
	source := `package main

import (
	"fmt"
	"os"
)

var globalVar = "I'm global"

type MyStruct struct {
	Field1 string
	Field2 int
}

func (m *MyStruct) Method() {
	localVar := "I'm local"
	fmt.Println(localVar)
}

func main() {
	mainVar := 42

	if true {
		ifVar := "inside if"
		fmt.Println(ifVar, mainVar, globalVar)

		for i := 0; i < 3; i++ {
			loopVar := i * 2
			fmt.Printf("Loop %d: %d\n", i, loopVar)
		}
	}

	// Function closure
	closure := func(param string) {
		closureVar := "I'm in closure"
		fmt.Println(param, closureVar, mainVar)
	}

	closure("test")

	// Struct usage
	s := &MyStruct{Field1: "hello", Field2: mainVar}
	s.Method()
}`

	// Cleanup
	defer os.Remove("test_complex.asta")

	// Parse original
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "complex.go", source, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse source: %v", err)
	}

	// Verify original has scope information
	if file.Scope == nil {
		t.Fatal("original file should have scope")
	}

	originalScopeCount := len(file.Scope.Objects)
	if originalScopeCount == 0 {
		t.Fatal("original scope should have objects")
	}

	// Count identifiers with objects in original
	originalIdentCount := 0
	ast.Inspect(file, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Obj != nil {
			originalIdentCount++
		}
		return true
	})

	if originalIdentCount == 0 {
		t.Fatal("original should have identifiers with objects")
	}

	t.Logf("Original: %d scope objects, %d identifiers with objects",
		originalScopeCount, originalIdentCount)

	// Save AST
	if err := SaveASTWithSourcePreservation(file, fset, "complex.go", "test_complex.asta"); err != nil {
		t.Fatalf("failed to save AST: %v", err)
	}

	// Load AST back
	restoredFile, restoredFset, _, err := LoadASTWithSourceReconstruction("test_complex.asta")
	if err != nil {
		t.Fatalf("failed to load AST: %v", err)
	}

	// Verify restored has scope information (this is the key test!)
	if restoredFile.Scope == nil {
		t.Fatal("restored file should have scope")
	}

	restoredScopeCount := len(restoredFile.Scope.Objects)
	if restoredScopeCount != originalScopeCount {
		t.Errorf("scope object count mismatch: original=%d, restored=%d",
			originalScopeCount, restoredScopeCount)
	}

	// Count restored identifiers with objects
	restoredIdentCount := 0
	ast.Inspect(restoredFile, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Obj != nil {
			restoredIdentCount++
		}
		return true
	})

	if restoredIdentCount != originalIdentCount {
		t.Errorf("identifier object count mismatch: original=%d, restored=%d",
			originalIdentCount, restoredIdentCount)
	}

	t.Logf("Restored: %d scope objects, %d identifiers with objects",
		restoredScopeCount, restoredIdentCount)

	// Verify perfect fidelity
	if err := VerifyPerfectFidelity(file, restoredFile, fset, restoredFset); err != nil {
		t.Errorf("fidelity check failed: %v", err)
	}
}

// TestGenericsAndComments tests AST archiving with generics and comments
func TestGenericsAndComments(t *testing.T) {
	source := `package main

// Generic function with type constraints
func Generic[T comparable](x, y T) bool {
	return x == y
}

// Generic type
type Container[T any] struct {
	Value T
}

func main() {
	// Test with int
	result := Generic(1, 2)
	_ = result

	// Test with container
	c := Container[string]{Value: "hello"}
	_ = c
}
`

	// Cleanup
	defer os.Remove("test_generics.asta")

	// Parse original
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "generics.go", source, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse source: %v", err)
	}

	// Verify comments are present
	if len(file.Comments) == 0 {
		t.Fatal("original should have comments")
	}

	// Save
	if err := SaveASTWithSourcePreservation(file, fset, "generics.go", "test_generics.asta"); err != nil {
		t.Fatalf("failed to save AST: %v", err)
	}

	// Load
	restoredFile, restoredFset, restoredSource, err := LoadASTWithSourceReconstruction("test_generics.asta")
	if err != nil {
		t.Fatalf("failed to load AST: %v", err)
	}

	// Verify comments are restored
	if len(restoredFile.Comments) != len(file.Comments) {
		t.Errorf("comment count mismatch: original=%d, restored=%d",
			len(file.Comments), len(restoredFile.Comments))
	}

	// Verify fidelity
	if err := VerifyPerfectFidelity(file, restoredFile, fset, restoredFset); err != nil {
		t.Errorf("fidelity check failed: %v", err)
	}

	// Verify source
	if restoredSource != source {
		t.Errorf("source code mismatch")
	}
}
