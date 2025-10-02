// Package main demonstrates function-related AST nodes.
// This file exercises ast.FuncType and ast.FuncLit nodes.
package main

import "fmt"

// AST Nodes Covered:
// - ast.FuncType
// - ast.FuncLit (function literals/closures)
// - ast.FuncDecl
// - ast.FieldList (parameters and results)

// Function type declarations
type UnaryFunc func(int) int
type BinaryFunc func(int, int) int
type NoParamFunc func() string
type NoReturnFunc func(int)
type MultiReturnFunc func(int, string) (int, string, bool)
type VariadicFunc func(string, ...interface{})
type ErrorFunc func() error
type NamedReturnFunc func(int) (result int, err error)

func funcMain() {
	fmt.Println("=== function_types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Simple function type
	fmt.Println("  ✓ ast.FuncType (simple):")
	var f1 UnaryFunc = func(x int) int {
		return x * 2
	}
	fmt.Printf("    UnaryFunc: f1(5) = %d\n", f1(5))

	// Binary function
	var f2 BinaryFunc = func(a, b int) int {
		return a + b
	}
	fmt.Printf("    BinaryFunc: f2(3, 4) = %d\n", f2(3, 4))

	// No parameters
	var f3 NoParamFunc = func() string {
		return "no params"
	}
	fmt.Printf("    NoParamFunc: f3() = %s\n", f3())

	// No return value
	var f4 NoReturnFunc = func(x int) {
		fmt.Printf("    NoReturnFunc: called with %d\n", x)
	}
	f4(42)

	// Multiple returns
	var f5 MultiReturnFunc = func(i int, s string) (int, string, bool) {
		return i * 2, s + s, true
	}
	i, s, b := f5(5, "hi")
	fmt.Printf("    MultiReturnFunc: returns (%d, %s, %v)\n", i, s, b)

	// Variadic function
	var f6 VariadicFunc = func(format string, args ...interface{}) {
		fmt.Printf("    VariadicFunc: "+format+"\n", args...)
	}
	f6("values: %d, %s, %v", 42, "test", true)

	// Function literal (ast.FuncLit) - closures
	fmt.Println("  ✓ ast.FuncLit (closures):")
	x := 10
	closure := func() int {
		x += 5 // captures x
		return x
	}
	fmt.Printf("    Closure (1st call): x = %d\n", closure())
	fmt.Printf("    Closure (2nd call): x = %d\n", closure())

	// Closure with parameters
	multiplier := 3
	closureWithParam := func(n int) int {
		return n * multiplier
	}
	fmt.Printf("    Closure with param: f(10) = %d\n", closureWithParam(10))

	// Immediately invoked function
	result := func(a, b int) int {
		return a * b
	}(6, 7)
	fmt.Printf("    IIFE (Immediately Invoked): 6*7 = %d\n", result)

	// Function returning function
	fmt.Println("  ✓ Higher-order functions:")
	makeAdder := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}
	add5 := makeAdder(5)
	add10 := makeAdder(10)
	fmt.Printf("    add5(3) = %d\n", add5(3))
	fmt.Printf("    add10(3) = %d\n", add10(3))

	// Function taking function as parameter
	applyFunc := func(f func(int) int, val int) int {
		return f(val)
	}
	double := func(x int) int { return x * 2 }
	triple := func(x int) int { return x * 3 }
	fmt.Printf("    applyFunc(double, 7) = %d\n", applyFunc(double, 7))
	fmt.Printf("    applyFunc(triple, 7) = %d\n", applyFunc(triple, 7))

	// Named return parameters (ast.FieldList)
	fmt.Println("  ✓ ast.FuncType with named returns:")
	var f7 NamedReturnFunc = func(x int) (result int, err error) {
		result = x * 2
		err = nil
		return // naked return
	}
	res, _ := f7(21)
	fmt.Printf("    NamedReturnFunc: result = %d\n", res)

	// Recursive function
	fmt.Println("  ✓ Recursive function:")
	var factorial func(int) int
	factorial = func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * factorial(n-1)
	}
	fmt.Printf("    factorial(5) = %d\n", factorial(5))

	// Method value
	fmt.Println("  ✓ Method value (function extracted from method):")
	calc := Calculator{value: 10}
	method := calc.Add // method value
	fmt.Printf("    Method value: method(5) = %d\n", method(5))

	// Method expression
	methodExpr := Calculator.Add // method expression
	fmt.Printf("    Method expression: methodExpr(calc, 7) = %d\n", methodExpr(calc, 7))

	// Defer with function literal
	fmt.Println("  ✓ Defer with function literal:")
	func() {
		defer func() {
			fmt.Println("    Deferred closure executed")
		}()
		fmt.Println("    Before defer")
	}()

	// Function literal in goroutine
	fmt.Println("  ✓ Goroutine with function literal:")
	done := make(chan bool)
	go func() {
		fmt.Println("    Goroutine function literal executed")
		done <- true
	}()
	<-done

	// Function with blank identifier parameters
	fmt.Println("  ✓ Function with blank identifiers:")
	ignoreParams := func(_ int, _ string) string {
		return "ignored params"
	}
	fmt.Printf("    ignoreParams(1, \"x\") = %s\n", ignoreParams(1, "x"))

	// Variadic function with different call styles
	fmt.Println("  ✓ Variadic function calls:")
	sum := func(nums ...int) int {
		total := 0
		for _, n := range nums {
			total += n
		}
		return total
	}
	fmt.Printf("    sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("    sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
	values := []int{10, 20, 30}
	fmt.Printf("    sum(values...) = %d\n", sum(values...))

	// Function type as struct field
	fmt.Println("  ✓ Function type in struct:")
	type Handler struct {
		OnClick func(int)
		OnHover func()
	}
	h := Handler{
		OnClick: func(x int) {
			fmt.Printf("    Clicked at position %d\n", x)
		},
		OnHover: func() {
			fmt.Println("    Hovered")
		},
	}
	h.OnClick(100)
	h.OnHover()

	fmt.Println("Summary: Comprehensive function type AST node coverage")
	fmt.Println("Primary AST Nodes: ast.FuncType, ast.FuncLit, ast.FuncDecl, ast.FieldList")
	fmt.Println("Features: function types, closures, higher-order functions, variadic, named returns")
	fmt.Println("========================================")
}

type Calculator struct {
	value int
}

func (c Calculator) Add(n int) int {
	return c.value + n
}

func main() {
	funcMain()
}
