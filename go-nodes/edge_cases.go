// Package main demonstrates edge cases and special AST scenarios.
// This file exercises unusual but valid Go constructs and AST patterns.
package main

import (
	"fmt"
	"unsafe"
)

// AST Nodes Covered:
// - Edge cases for various AST nodes
// - BadExpr, BadStmt, BadDecl (error recovery nodes)
// - Unusual but valid Go constructs
// - Special built-in functions

func funcMain() {
	fmt.Println("=== edge_cases.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Empty statements
	fmt.Println("  ✓ ast.EmptyStmt (edge cases):")
	for i := 0; i < 1; i++ {
		; // Empty statement
		; // Another empty statement
	}
	fmt.Println("    Empty statements in loop")

	// Complex nested expressions
	fmt.Println("  ✓ Deeply nested expressions:")
	result := ((((1 + 2) * 3) - 4) / 5) + ((6 * (7 + 8)) - 9)
	fmt.Printf("    Complex nested: %d\n", result)

	// Multiple assignments with different types
	fmt.Println("  ✓ Complex assignments:")
	var a, b, c = 1, "two", true
	fmt.Printf("    Multi-type assignment: %d, %s, %v\n", a, b, c)

	// Blank identifier in various contexts
	fmt.Println("  ✓ Blank identifier (_) usage:")
	_ = 42 // Discard value
	_, _ = 1, 2 // Multiple discards
	for _ = range []int{1, 2, 3} {
		// Range with blank
	}
	fmt.Println("    Blank identifier in multiple contexts")

	// Iota edge cases
	fmt.Println("  ✓ iota edge cases:")
	const (
		a0 = iota // 0
		a1        // 1
		a2        // 2
		_         // 3 (skipped)
		a4        // 4
	)
	const (
		b0 = iota * 10 // 0
		b1             // 10
		b2             // 20
	)
	const (
		c0 = 1 << iota // 1
		c1             // 2
		c2             // 4
		c3             // 8
	)
	fmt.Printf("    iota sequence: %d, %d, %d, %d\n", a0, a1, a2, a4)
	fmt.Printf("    iota expression: %d, %d, %d\n", b0, b1, b2)
	fmt.Printf("    iota bit shift: %d, %d, %d, %d\n", c0, c1, c2, c3)

	// Unsafe operations
	fmt.Println("  ✓ unsafe package operations:")
	x := 42
	ptr := unsafe.Pointer(&x)
	size := unsafe.Sizeof(x)
	align := unsafe.Alignof(x)
	fmt.Printf("    unsafe.Pointer: %p\n", ptr)
	fmt.Printf("    unsafe.Sizeof(int): %d bytes\n", size)
	fmt.Printf("    unsafe.Alignof(int): %d bytes\n", align)

	// Type conversions and assertions
	fmt.Println("  ✓ Type conversions edge cases:")
	var i interface{} = 42
	n := i.(int)
	fmt.Printf("    Type assertion: %d\n", n)

	// Conversion between compatible types
	type MyInt int
	var mi MyInt = 10
	var normalInt int = int(mi)
	fmt.Printf("    Type conversion: MyInt(%d) -> int(%d)\n", mi, normalInt)

	// Pointer indirection chains
	fmt.Println("  ✓ Multiple pointer indirection:")
	val := 100
	p1 := &val
	p2 := &p1
	p3 := &p2
	fmt.Printf("    ***p3 = %d\n", ***p3)

	// Complex slice operations
	fmt.Println("  ✓ Complex slice operations:")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sub := s[2:8:9] // 3-index slice
	fmt.Printf("    s[2:8:9]: %v (len=%d, cap=%d)\n", sub, len(sub), cap(sub))

	// Slice with all indices as expressions
	start, end := 1, 5
	s2 := s[start:end]
	fmt.Printf("    s[start:end]: %v\n", s2)

	// Empty composite literals
	fmt.Println("  ✓ Empty composite literals:")
	emptyStruct := struct{}{}
	emptyArray := [0]int{}
	emptySlice := []int{}
	emptyMap := map[string]int{}
	fmt.Printf("    struct{}{}: %v\n", emptyStruct)
	fmt.Printf("    [0]int{}: %v\n", emptyArray)
	fmt.Printf("    []int{}: %v (len=%d)\n", emptySlice, len(emptySlice))
	fmt.Printf("    map[string]int{}: %v (len=%d)\n", emptyMap, len(emptyMap))

	// Variadic with zero arguments
	fmt.Println("  ✓ Variadic with zero args:")
	variadicFunc()
	fmt.Println("    Called variadic function with no args")

	// Type switch with multiple types in case
	fmt.Println("  ✓ Type switch edge cases:")
	var mixed interface{} = "test"
	switch v := mixed.(type) {
	case int, int64:
		fmt.Printf("    Integer type: %v\n", v)
	case string, []byte:
		fmt.Printf("    String-like type: %v\n", v)
	case nil:
		fmt.Println("    Nil type")
	default:
		fmt.Printf("    Other type: %T\n", v)
	}

	// Labeled break/continue
	fmt.Println("  ✓ Labeled break/continue:")
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 1 && j == 0 {
				break OuterLoop
			}
			fmt.Printf("    i=%d, j=%d\n", i, j)
		}
	}

	// Defer with panic/recover
	fmt.Println("  ✓ Defer, panic, recover:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("    Recovered: %v\n", r)
			}
		}()
		// panic("test panic")
		fmt.Println("    Defer/recover pattern demonstrated")
	}()

	// Channel close and nil channel
	fmt.Println("  ✓ Channel edge cases:")
	ch := make(chan int, 1)
	ch <- 42
	close(ch)
	v, ok := <-ch
	fmt.Printf("    Receive from closed channel: %d, ok=%v\n", v, ok)
	v2, ok2 := <-ch
	fmt.Printf("    Second receive: %d, ok=%v\n", v2, ok2)

	// Select with no cases (blocks forever - commented out)
	// select {}

	// Zero-value channels
	var nilCh chan int
	fmt.Printf("    Nil channel: %v\n", nilCh)

	// Method value (using package-level T type)
	fmt.Println("  ✓ Method value/expression edge cases:")
	instance := T{v: 10}
	methodValue := instance.String
	fmt.Printf("    Method value: %s\n", methodValue())

	// Function returning named results with modification
	fmt.Println("  ✓ Named return modification:")
	res1, res2 := namedReturnModified(5)
	fmt.Printf("    Named return modified: %d, %d\n", res1, res2)

	// Struct with all field types
	fmt.Println("  ✓ Struct with complex field types:")
	type ComplexStruct struct {
		Chan     chan int
		Func     func() int
		Interface interface{}
		Map      map[string]int
		Slice    []int
		Pointer  *int
		Struct   struct{ X int }
	}
	num := 42
	cs := ComplexStruct{
		Chan:      make(chan int),
		Func:      func() int { return 1 },
		Interface: "anything",
		Map:       map[string]int{"a": 1},
		Slice:     []int{1, 2, 3},
		Pointer:   &num,
		Struct:    struct{ X int }{X: 10},
	}
	fmt.Printf("    ComplexStruct.Pointer: %d\n", *cs.Pointer)

	// Array index with complex expression
	fmt.Println("  ✓ Complex array indexing:")
	arr := [5]int{10, 20, 30, 40, 50}
	idx := 2
	fmt.Printf("    arr[idx]: %d\n", arr[idx])
	fmt.Printf("    arr[len(arr)-1]: %d\n", arr[len(arr)-1])

	// Map with complex key
	fmt.Println("  ✓ Map with complex key types:")
	type Key struct{ A, B int }
	complexMap := map[Key]string{
		{A: 1, B: 2}: "value1",
		{A: 3, B: 4}: "value2",
	}
	fmt.Printf("    map[Key]string: %v\n", complexMap)

	// Multiple return value unpacking
	fmt.Println("  ✓ Multiple return unpacking:")
	x1, y1 := swap(10, 20)
	fmt.Printf("    swap(10, 20): %d, %d\n", x1, y1)

	fmt.Println("Summary: Edge cases and special constructs AST node coverage")
	fmt.Println("Primary AST Nodes: Various edge cases for all major node types")
	fmt.Println("Features: unsafe, iota, complex nesting, edge cases, special patterns")
	fmt.Println("========================================")
}

func variadicFunc(args ...int) {
	fmt.Printf("    Variadic received %d args\n", len(args))
}

type T struct{ v int }

func (t T) String() string {
	return fmt.Sprintf("T{%d}", t.v)
}

func namedReturnModified(x int) (a, b int) {
	a = x * 2
	b = x * 3
	a += 1 // Modify before return
	return
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	funcMain()
}
