package archive

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
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

// TestConsumerAPI demonstrates the complete API for consuming libraries
func TestConsumerAPI(t *testing.T) {
	// Setup: Create test archives
	testDir := "test_archives"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create a test archive
	source := `package example

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

func main() {
	person := NewPerson("Alice")
	person.Greet()
}
`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example.go", source, parser.ParseComments)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	archivePath := filepath.Join(testDir, "example.asta")
	if err := SaveASTWithSourcePreservation(file, fset, "example.go", archivePath); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	// Test 1: Load
	t.Run("Load", func(t *testing.T) {
		archive, err := Load(archivePath)
		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}

		if archive.GetFilename() != "example.go" {
			t.Errorf("expected filename 'example.go', got '%s'", archive.GetFilename())
		}

		if archive.GetPackageName() != "example" {
			t.Errorf("expected package 'example', got '%s'", archive.GetPackageName())
		}

		if archive.GetSourceCode() != source {
			t.Errorf("source code mismatch")
		}
	})

	// Test 2: GetAST with full reconstruction
	t.Run("GetAST", func(t *testing.T) {
		archive, _ := Load(archivePath)
		file, fset, err := archive.GetAST()
		if err != nil {
			t.Fatalf("GetAST failed: %v", err)
		}

		if file.Scope == nil {
			t.Error("expected Scope to be reconstructed")
		}

		if fset == nil {
			t.Error("expected FileSet to be returned")
		}
	})

	// Test 3: GetCleanedAST (fast access)
	t.Run("GetCleanedAST", func(t *testing.T) {
		archive, _ := Load(archivePath)
		cleanAST := archive.GetCleanedAST()
		if cleanAST == nil {
			t.Error("expected cleaned AST")
		}
	})

	// Test 4: ExtractFunctions
	t.Run("ExtractFunctions", func(t *testing.T) {
		archive, _ := Load(archivePath)
		funcs, err := ExtractFunctions(archive)
		if err != nil {
			t.Fatalf("ExtractFunctions failed: %v", err)
		}

		// Should have: NewPerson, (*Person).Greet, main
		if len(funcs) != 3 {
			t.Errorf("expected 3 functions, got %d", len(funcs))
		}
	})

	// Test 5: ExtractTypes
	t.Run("ExtractTypes", func(t *testing.T) {
		archive, _ := Load(archivePath)
		types, err := ExtractTypes(archive)
		if err != nil {
			t.Fatalf("ExtractTypes failed: %v", err)
		}

		// Should have: Person
		if len(types) != 1 {
			t.Errorf("expected 1 type, got %d", len(types))
		}

		if len(types) > 0 && types[0].Name.Name != "Person" {
			t.Errorf("expected type 'Person', got '%s'", types[0].Name.Name)
		}
	})

	// Test 6: GetImports
	t.Run("GetImports", func(t *testing.T) {
		archive, _ := Load(archivePath)
		imports := GetImports(archive)

		if len(imports) != 1 {
			t.Errorf("expected 1 import, got %d", len(imports))
		}

		if len(imports) > 0 && imports[0] != "fmt" {
			t.Errorf("expected import 'fmt', got '%s'", imports[0])
		}
	})

	// Test 7: GetFunctionNames
	t.Run("GetFunctionNames", func(t *testing.T) {
		archive, _ := Load(archivePath)
		names, err := GetFunctionNames(archive)
		if err != nil {
			t.Fatalf("GetFunctionNames failed: %v", err)
		}

		if len(names) != 3 {
			t.Errorf("expected 3 function names, got %d", len(names))
		}
	})

	// Test 8: GetTypeNames
	t.Run("GetTypeNames", func(t *testing.T) {
		archive, _ := Load(archivePath)
		names, err := GetTypeNames(archive)
		if err != nil {
			t.Fatalf("GetTypeNames failed: %v", err)
		}

		if len(names) != 1 || names[0] != "Person" {
			t.Errorf("expected type name 'Person', got %v", names)
		}
	})

	// Test 9: FindNodesByType
	t.Run("FindNodesByType", func(t *testing.T) {
		archive, _ := Load(archivePath)
		nodes, err := FindNodesByType(archive, "*ast.FuncDecl")
		if err != nil {
			t.Fatalf("FindNodesByType failed: %v", err)
		}

		if len(nodes) != 3 {
			t.Errorf("expected 3 FuncDecl nodes, got %d", len(nodes))
		}
	})
}

// TestLoadAll tests loading all archives from nodes/ast directory
func TestLoadAll(t *testing.T) {
	// This test uses the actual generated .asta files
	archives, err := LoadAll("../nodes/ast")
	if err != nil {
		// It's okay if the directory doesn't exist in test env
		t.Skipf("Skipping test - nodes/ast not available: %v", err)
	}

	if len(archives) == 0 {
		t.Error("expected at least one archive")
	}

	// Verify we can access each archive
	for _, archive := range archives {
		if archive.GetFilename() == "" {
			t.Error("archive has empty filename")
		}

		if archive.GetPackageName() != "main" {
			t.Errorf("expected package 'main', got '%s'", archive.GetPackageName())
		}
	}
}

// TestWalk tests the iterator pattern
func TestWalk(t *testing.T) {
	testDir := "test_walk"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create two test archives
	for i, name := range []string{"file1.asta", "file2.asta"} {
		source := fmt.Sprintf("package test%d\n\nfunc Test%d() {}\n", i, i)
		fset := token.NewFileSet()
		file, _ := parser.ParseFile(fset, name, source, parser.ParseComments)
		SaveASTWithSourcePreservation(file, fset, name, filepath.Join(testDir, name))
	}

	// Walk archives
	count := 0
	err := Walk(testDir, func(archive *ASTArchive) error {
		count++
		if archive.GetFilename() == "" {
			return fmt.Errorf("empty filename")
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Walk failed: %v", err)
	}

	if count != 2 {
		t.Errorf("expected to walk 2 archives, got %d", count)
	}
}

// TestWalkError tests error handling in Walk
func TestWalkError(t *testing.T) {
	testDir := "test_walk_error"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create a test archive
	source := "package test\n"
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "test.go", source, parser.ParseComments)
	SaveASTWithSourcePreservation(file, fset, "test.go", filepath.Join(testDir, "test.asta"))

	// Walk should stop on error
	count := 0
	err := Walk(testDir, func(archive *ASTArchive) error {
		count++
		return fmt.Errorf("intentional error")
	})

	if err == nil {
		t.Error("expected error from Walk")
	}

	if count != 1 {
		t.Errorf("expected walk to stop after 1 archive, processed %d", count)
	}
}
