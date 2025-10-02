// Package main demonstrates declaration AST nodes.
// This file exercises all ast.Decl and ast.Spec node types.
package main

import (
	"fmt"
	"os"
)

// AST Nodes Covered:
// - ast.GenDecl (import, const, var, type)
// - ast.FuncDecl
// - ast.ImportSpec
// - ast.ValueSpec
// - ast.TypeSpec

// Import declarations (ast.GenDecl with ast.ImportSpec)
// Single import covered above
// Grouped imports
// Aliased import (os as operating system - but we use standard)

// Constant declarations (ast.GenDecl with ast.ValueSpec)
const SingleConst = 42

const (
	First  = 1
	Second = 2
	Third  = 3
)

// Const with iota
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Const with type
const TypedConst int = 100

// Multiple consts in one line
const X, Y, Z = 1, 2, 3

// Variable declarations (ast.GenDecl with ast.ValueSpec)
var SingleVar = "single"

var (
	GroupedVar1 = 10
	GroupedVar2 = 20
	GroupedVar3 = 30
)

// Var with explicit type
var TypedVar int = 42

// Multiple vars
var A, B, C int = 1, 2, 3

// Var without initialization
var UninitVar int

// Type declarations (ast.GenDecl with ast.TypeSpec)
type SingleType int

type (
	GroupedType1 int
	GroupedType2 string
	GroupedType3 bool
)

// Type alias (Go 1.9+)
type TypeAlias = int

// Struct type
type Person struct {
	Name string
	Age  int
}

// Interface type
type Speaker interface {
	Speak() string
}

// Function type
type BinaryOp func(int, int) int

// Channel type
type IntChannel chan int

// Map type
type StringIntMap map[string]int

// Slice type
type IntSlice []int

// Array type
type IntArray [5]int

// Pointer type
type IntPointer *int

// Function declarations (ast.FuncDecl)

// Simple function
func simpleFunc() {
	fmt.Println("  ✓ ast.FuncDecl (simple): simpleFunc called")
}

// Function with parameters
func withParams(a int, b string) {
	fmt.Printf("  ✓ ast.FuncDecl (with params): a=%d, b=%q\n", a, b)
}

// Function with return value
func withReturn() int {
	fmt.Println("  ✓ ast.FuncDecl (with return): withReturn called")
	return 42
}

// Function with multiple returns
func multipleReturns() (int, string, bool) {
	fmt.Println("  ✓ ast.FuncDecl (multiple returns): called")
	return 1, "hello", true
}

// Function with named returns
func namedReturns() (result int, err error) {
	fmt.Println("  ✓ ast.FuncDecl (named returns): called")
	result = 100
	err = nil
	return
}

// Function with variadic parameters
func variadicFunc(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Printf("  ✓ ast.FuncDecl (variadic): sum=%d\n", sum)
	return sum
}

// Method declaration (function with receiver)
func (p Person) Speak() string {
	return fmt.Sprintf("Hi, I'm %s", p.Name)
}

// Method with pointer receiver
func (p *Person) SetAge(age int) {
	p.Age = age
}

// Method declarations demonstrated
type Calculator struct {
	value int
}

func (c Calculator) Add(n int) int {
	fmt.Printf("  ✓ ast.FuncDecl (method value receiver): Add called\n")
	return c.value + n
}

func (c *Calculator) Multiply(n int) {
	fmt.Printf("  ✓ ast.FuncDecl (method pointer receiver): Multiply called\n")
	c.value *= n
}

func main() {
	fmt.Println("=== declarations.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Demonstrate imports
	fmt.Printf("  ✓ ast.GenDecl (import) and ast.ImportSpec:\n")
	fmt.Printf("    fmt package: %T\n", fmt.Println)
	fmt.Printf("    os package: %T\n", os.Stdout)

	// Demonstrate constants
	fmt.Printf("  ✓ ast.GenDecl (const) and ast.ValueSpec:\n")
	fmt.Printf("    SingleConst=%d\n", SingleConst)
	fmt.Printf("    First=%d, Second=%d, Third=%d\n", First, Second, Third)
	fmt.Printf("    Monday=%d, Friday=%d, Sunday=%d\n", Monday, Friday, Sunday)
	fmt.Printf("    TypedConst=%d (type: int)\n", TypedConst)
	fmt.Printf("    X=%d, Y=%d, Z=%d\n", X, Y, Z)

	// Demonstrate variables
	fmt.Printf("  ✓ ast.GenDecl (var) and ast.ValueSpec:\n")
	fmt.Printf("    SingleVar=%q\n", SingleVar)
	fmt.Printf("    GroupedVar1=%d, GroupedVar2=%d, GroupedVar3=%d\n", GroupedVar1, GroupedVar2, GroupedVar3)
	fmt.Printf("    TypedVar=%d\n", TypedVar)
	fmt.Printf("    A=%d, B=%d, C=%d\n", A, B, C)
	fmt.Printf("    UninitVar=%d (zero value)\n", UninitVar)

	// Demonstrate types
	fmt.Printf("  ✓ ast.GenDecl (type) and ast.TypeSpec:\n")
	var st SingleType = 42
	fmt.Printf("    SingleType: %d\n", st)
	var gt1 GroupedType1 = 10
	var gt2 GroupedType2 = "hello"
	fmt.Printf("    GroupedType1: %d, GroupedType2: %q\n", gt1, gt2)
	var ta TypeAlias = 99
	fmt.Printf("    TypeAlias: %d\n", ta)

	// Demonstrate struct
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("    Person struct: %+v\n", p)

	// Demonstrate interface implementation
	var s Speaker = p
	fmt.Printf("    Speaker interface: %s\n", s.Speak())

	// Demonstrate function type
	var op BinaryOp = func(a, b int) int { return a + b }
	fmt.Printf("    BinaryOp function type: 5+3=%d\n", op(5, 3))

	// Demonstrate functions
	fmt.Printf("  ✓ ast.FuncDecl:\n")
	simpleFunc()
	withParams(42, "test")
	ret := withReturn()
	fmt.Printf("    Returned: %d\n", ret)
	i, str, b := multipleReturns()
	fmt.Printf("    Multiple returns: %d, %q, %v\n", i, str, b)
	nr, _ := namedReturns()
	fmt.Printf("    Named returns: %d\n", nr)
	variadicFunc(1, 2, 3, 4, 5)

	// Demonstrate methods
	msg := p.Speak()
	fmt.Printf("  ✓ ast.FuncDecl (method): %s\n", msg)
	p.SetAge(31)
	fmt.Printf("  ✓ ast.FuncDecl (pointer method): Age set to %d\n", p.Age)

	calc := Calculator{value: 10}
	result := calc.Add(5)
	fmt.Printf("    Calculator.Add: %d\n", result)
	calc.Multiply(3)
	fmt.Printf("    Calculator.Multiply: value=%d\n", calc.value)

	fmt.Println("Summary: All declaration AST node types exercised")
	fmt.Println("Primary AST Nodes: ast.GenDecl, ast.FuncDecl, ast.ImportSpec, ast.ValueSpec, ast.TypeSpec")
	fmt.Println("========================================")
}
