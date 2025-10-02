// Package main demonstrates expression AST nodes.
// This file exercises all ast.Expr node types.
package main

import "fmt"

// AST Nodes Covered:
// - ast.ParenExpr
// - ast.SelectorExpr
// - ast.IndexExpr
// - ast.SliceExpr
// - ast.TypeAssertExpr
// - ast.CallExpr
// - ast.StarExpr
// - ast.UnaryExpr
// - ast.BinaryExpr
// - ast.KeyValueExpr
// - ast.CompositeLit
// - ast.FuncLit
// - ast.Ellipsis

func main() {
	fmt.Println("=== expressions.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// ParenExpr - parenthesized expressions
	result := (2 + 3) * 4
	fmt.Printf("  ✓ ast.ParenExpr: (2 + 3) * 4 = %d\n", result)

	// BinaryExpr - all binary operators
	fmt.Println("  ✓ ast.BinaryExpr (arithmetic):")
	fmt.Printf("    ADD: 5 + 3 = %d\n", 5+3)
	fmt.Printf("    SUB: 5 - 3 = %d\n", 5-3)
	fmt.Printf("    MUL: 5 * 3 = %d\n", 5*3)
	fmt.Printf("    QUO: 15 / 3 = %d\n", 15/3)
	fmt.Printf("    REM: 17 %% 5 = %d\n", 17%5)

	fmt.Println("  ✓ ast.BinaryExpr (bitwise):")
	fmt.Printf("    AND: 12 & 10 = %d\n", 12&10)
	fmt.Printf("    OR: 12 | 10 = %d\n", 12|10)
	fmt.Printf("    XOR: 12 ^ 10 = %d\n", 12^10)
	fmt.Printf("    AND_NOT: 12 &^ 10 = %d\n", 12&^10)
	fmt.Printf("    SHL: 3 << 2 = %d\n", 3<<2)
	fmt.Printf("    SHR: 12 >> 2 = %d\n", 12>>2)

	fmt.Println("  ✓ ast.BinaryExpr (comparison):")
	fmt.Printf("    EQL: 5 == 5 = %v\n", 5 == 5)
	fmt.Printf("    NEQ: 5 != 3 = %v\n", 5 != 3)
	fmt.Printf("    LSS: 3 < 5 = %v\n", 3 < 5)
	fmt.Printf("    LEQ: 5 <= 5 = %v\n", 5 <= 5)
	fmt.Printf("    GTR: 5 > 3 = %v\n", 5 > 3)
	fmt.Printf("    GEQ: 5 >= 5 = %v\n", 5 >= 5)

	fmt.Println("  ✓ ast.BinaryExpr (logical):")
	fmt.Printf("    LAND: true && false = %v\n", true && false)
	fmt.Printf("    LOR: true || false = %v\n", true || false)

	// UnaryExpr - all unary operators
	fmt.Println("  ✓ ast.UnaryExpr:")
	var x int = 42
	var ptr *int = &x
	fmt.Printf("    ADD (unary plus): +42 = %d\n", +42)
	fmt.Printf("    SUB (unary minus): -42 = %d\n", -42)
	fmt.Printf("    NOT: !true = %v\n", !true)
	fmt.Printf("    XOR (bitwise not): ^42 = %d\n", ^42)
	fmt.Printf("    AND (address-of): &x = %p\n", ptr)
	fmt.Printf("    MUL (dereference): *ptr = %d\n", *ptr)
	fmt.Printf("    ARROW (receive): will be in channels\n")

	// StarExpr - pointer type and dereference
	var value int = 100
	var pointer *int = &value
	var deref int = *pointer
	fmt.Printf("  ✓ ast.StarExpr (dereference): *pointer = %d\n", deref)

	// SelectorExpr - field and method selection
	type Point struct {
		X, Y int
	}
	p := Point{X: 10, Y: 20}
	fmt.Printf("  ✓ ast.SelectorExpr (field): p.X = %d, p.Y = %d\n", p.X, p.Y)
	fmt.Printf("  ✓ ast.SelectorExpr (package.Name): fmt.Println\n")

	// IndexExpr - array/slice/map indexing
	arr := [5]int{10, 20, 30, 40, 50}
	slice := []string{"a", "b", "c"}
	m := map[string]int{"one": 1, "two": 2}
	fmt.Printf("  ✓ ast.IndexExpr (array): arr[2] = %d\n", arr[2])
	fmt.Printf("  ✓ ast.IndexExpr (slice): slice[1] = %q\n", slice[1])
	fmt.Printf("  ✓ ast.IndexExpr (map): m[\"two\"] = %d\n", m["two"])

	// SliceExpr - slice operations
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("  ✓ ast.SliceExpr (simple): s[1:3] = %v\n", s[1:3])
	fmt.Printf("  ✓ ast.SliceExpr (from start): s[:3] = %v\n", s[:3])
	fmt.Printf("  ✓ ast.SliceExpr (to end): s[2:] = %v\n", s[2:])
	fmt.Printf("  ✓ ast.SliceExpr (full): s[1:4:5] = %v\n", s[1:4:5])

	// CallExpr - function calls
	fmt.Printf("  ✓ ast.CallExpr (function): add(2, 3) = %d\n", add(2, 3))
	fmt.Printf("  ✓ ast.CallExpr (builtin): len(slice) = %d\n", len(slice))
	fmt.Printf("  ✓ ast.CallExpr (method): will be shown in methods\n")
	floatVal := 3.14
	intVal := int(floatVal)
	fmt.Printf("  ✓ ast.CallExpr (type conversion): int(3.14) = %d\n", intVal)

	// Ellipsis - variadic parameters
	fmt.Printf("  ✓ ast.Ellipsis (variadic): sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
	nums := []int{10, 20, 30}
	fmt.Printf("  ✓ ast.Ellipsis (spread): sum(nums...) = %d\n", sum(nums...))

	// TypeAssertExpr - type assertions
	var i interface{} = "hello"
	str, ok := i.(string)
	fmt.Printf("  ✓ ast.TypeAssertExpr: i.(string) = %q, ok = %v\n", str, ok)
	num, ok := i.(int)
	fmt.Printf("  ✓ ast.TypeAssertExpr (failed): i.(int) = %d, ok = %v\n", num, ok)

	// CompositeLit - composite literals
	fmt.Println("  ✓ ast.CompositeLit:")
	arrLit := [3]int{1, 2, 3}
	fmt.Printf("    Array: %v\n", arrLit)
	sliceLit := []string{"a", "b", "c"}
	fmt.Printf("    Slice: %v\n", sliceLit)
	structLit := Point{X: 5, Y: 10}
	fmt.Printf("    Struct: %v\n", structLit)
	mapLit := map[string]int{"x": 1, "y": 2}
	fmt.Printf("    Map: %v\n", mapLit)

	// KeyValueExpr - key-value pairs in composite literals
	person := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  25,
	}
	fmt.Printf("  ✓ ast.KeyValueExpr: {Name: %q, Age: %d}\n", person.Name, person.Age)

	// FuncLit - function literals (closures)
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("  ✓ ast.FuncLit (closure): multiply(3, 4) = %d\n", multiply(3, 4))

	// Nested expressions
	complex := ((5 + 3) * 2) - (10 / 2)
	fmt.Printf("  ✓ Complex nested expression: ((5+3)*2)-(10/2) = %d\n", complex)

	fmt.Println("Summary: All major expression AST nodes exercised")
	fmt.Println("Primary AST Nodes: ast.BinaryExpr, ast.UnaryExpr, ast.CallExpr, ast.SelectorExpr,")
	fmt.Println("                   ast.IndexExpr, ast.SliceExpr, ast.TypeAssertExpr, ast.CompositeLit,")
	fmt.Println("                   ast.FuncLit, ast.ParenExpr, ast.KeyValueExpr, ast.Ellipsis, ast.StarExpr")
	fmt.Println("========================================")
}

func add(a, b int) int {
	return a + b
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
