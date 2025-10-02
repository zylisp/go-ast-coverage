// Package main demonstrates identifier AST nodes.
// This file exercises ast.Ident in various contexts.
package main

import "fmt"

// AST Nodes Covered:
// - ast.Ident (variables, constants, types, functions, labels)
// - ast.Object
// - ast.Scope

// Package-level identifiers
const packageConst = "package constant"

var packageVar = "package variable"

type CustomType int

func main() {
	fmt.Println("=== identifiers.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Variable identifiers
	var localVar int = 42
	shortVar := "short declaration"
	fmt.Printf("  ✓ ast.Ident (var declaration): localVar = %d\n", localVar)
	fmt.Printf("  ✓ ast.Ident (short declaration): shortVar = %q\n", shortVar)

	// Multiple variable declarations
	var x, y, z int = 1, 2, 3
	fmt.Printf("  ✓ ast.Ident (multiple vars): x=%d, y=%d, z=%d\n", x, y, z)

	// Constant identifiers
	const localConst = 100
	const (
		First  = 1
		Second = 2
		Third  = 3
	)
	fmt.Printf("  ✓ ast.Ident (const): localConst = %d\n", localConst)
	fmt.Printf("  ✓ ast.Ident (const block): First=%d, Second=%d, Third=%d\n", First, Second, Third)

	// Type identifiers
	type MyInt int
	type MyString string
	var mi MyInt = 42
	var ms MyString = "hello"
	fmt.Printf("  ✓ ast.Ident (type declaration): MyInt = %d\n", mi)
	fmt.Printf("  ✓ ast.Ident (type declaration): MyString = %q\n", ms)

	// Function identifiers
	result := helperFunction(10, 20)
	fmt.Printf("  ✓ ast.Ident (function name): helperFunction result = %d\n", result)

	// Method identifiers
	obj := MyObject{value: 42}
	methodResult := obj.Method()
	fmt.Printf("  ✓ ast.Ident (method name): Method result = %d\n", methodResult)

	// Label identifiers
	counter := 0
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			counter++
			if counter == 5 {
				fmt.Printf("  ✓ ast.Ident (label): Breaking from OuterLoop at counter=%d\n", counter)
				break OuterLoop
			}
		}
	}

	// Blank identifier
	_, ignored := multiReturn()
	fmt.Printf("  ✓ ast.Ident (blank identifier): ignored second return value, got %d\n", ignored)

	// Package identifier in qualified names
	fmt.Printf("  ✓ ast.Ident (package qualifier): fmt.Println\n")

	// Predeclared identifiers
	var length = len("hello")
	var capacity = cap(make([]int, 5, 10))
	fmt.Printf("  ✓ ast.Ident (builtin len): %d\n", length)
	fmt.Printf("  ✓ ast.Ident (builtin cap): %d\n", capacity)

	// Field identifiers in struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("  ✓ ast.Ident (struct field): Name=%q, Age=%d\n", p.Name, p.Age)

	// Iota identifier
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)
	fmt.Printf("  ✓ ast.Ident (iota): A=%d, B=%d, C=%d\n", A, B, C)

	// Underscore in identifier names
	var my_variable = "underscore_style"
	var myVariable = "camelCase"
	var MyExportedVar = "PascalCase"
	fmt.Printf("  ✓ ast.Ident (naming styles): %q, %q, %q\n", my_variable, myVariable, MyExportedVar)

	fmt.Println("Summary: 20+ unique identifier contexts exercised")
	fmt.Println("Primary AST Nodes: ast.Ident, ast.Object, ast.Scope")
	fmt.Println("========================================")
}

// Helper function to demonstrate function identifiers
func helperFunction(a, b int) int {
	return a + b
}

// Type with method to demonstrate method identifiers
type MyObject struct {
	value int
}

func (m MyObject) Method() int {
	return m.value * 2
}

// Function with multiple returns to demonstrate blank identifier
func multiReturn() (int, int) {
	return 1, 2
}
