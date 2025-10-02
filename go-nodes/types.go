// Package main demonstrates type expression AST nodes.
// This file exercises various type-related AST nodes.
package main

import "fmt"

// AST Nodes Covered:
// - ast.Ident (for built-in types)
// - ast.ArrayType
// - ast.StructType
// - ast.FuncType
// - ast.InterfaceType
// - ast.MapType
// - ast.ChanType
// - ast.StarExpr (pointer types)

func main() {
	fmt.Println("=== types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Built-in types (ast.Ident)
	fmt.Println("  ✓ ast.Ident (built-in types):")
	var vBool bool = true
	var vInt int = 42
	var vInt8 int8 = 127
	var vInt16 int16 = 32767
	var vInt32 int32 = 2147483647
	var vInt64 int64 = 9223372036854775807
	var vUint uint = 42
	var vUint8 uint8 = 255
	var vUint16 uint16 = 65535
	var vUint32 uint32 = 4294967295
	var vUint64 uint64 = 18446744073709551615
	var vUintptr uintptr = 0x12345
	var vFloat32 float32 = 3.14
	var vFloat64 float64 = 2.718281828
	var vComplex64 complex64 = 1 + 2i
	var vComplex128 complex128 = 3 + 4i
	var vString string = "hello"
	var vByte byte = 'A'      // alias for uint8
	var vRune rune = '世'       // alias for int32
	var vError error = nil     // built-in interface type
	var vAny any = 42          // alias for interface{} (Go 1.18+)

	fmt.Printf("    bool: %v\n", vBool)
	fmt.Printf("    int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", vInt, vInt8, vInt16, vInt32, vInt64)
	fmt.Printf("    uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d\n", vUint, vUint8, vUint16, vUint32, vUint64)
	fmt.Printf("    uintptr: 0x%x\n", vUintptr)
	fmt.Printf("    float32: %f, float64: %f\n", vFloat32, vFloat64)
	fmt.Printf("    complex64: %v, complex128: %v\n", vComplex64, vComplex128)
	fmt.Printf("    string: %q, byte: %c, rune: %c\n", vString, vByte, vRune)
	fmt.Printf("    error: %v, any: %v\n", vError, vAny)

	// Array types (ast.ArrayType)
	fmt.Println("  ✓ ast.ArrayType:")
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 [3]string = [3]string{"a", "b", "c"}
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}} // nested
	fmt.Printf("    [5]int: %v\n", arr1)
	fmt.Printf("    [3]string: %v\n", arr2)
	fmt.Printf("    [2][3]int: %v\n", arr3)

	// Slice types (ast.ArrayType with nil len)
	fmt.Println("  ✓ ast.ArrayType (slice):")
	var slice1 []int = []int{1, 2, 3}
	var slice2 []string = []string{"x", "y", "z"}
	var slice3 [][]int = [][]int{{1, 2}, {3, 4}}
	fmt.Printf("    []int: %v\n", slice1)
	fmt.Printf("    []string: %v\n", slice2)
	fmt.Printf("    [][]int: %v\n", slice3)

	// Pointer types (ast.StarExpr)
	fmt.Println("  ✓ ast.StarExpr (pointer types):")
	var ptr1 *int
	x := 42
	ptr1 = &x
	var ptr2 **int = &ptr1 // pointer to pointer
	fmt.Printf("    *int: %v (value: %d)\n", ptr1, *ptr1)
	fmt.Printf("    **int: %v (value: %d)\n", ptr2, **ptr2)

	// Struct types (ast.StructType)
	fmt.Println("  ✓ ast.StructType:")
	type Simple struct {
		Field1 int
		Field2 string
	}
	var s1 Simple = Simple{Field1: 10, Field2: "test"}
	fmt.Printf("    Simple struct: %+v\n", s1)

	// Empty struct
	type Empty struct{}
	var s2 Empty = Empty{}
	fmt.Printf("    Empty struct: %+v\n", s2)

	// Anonymous struct
	var s3 = struct {
		X int
		Y int
	}{X: 1, Y: 2}
	fmt.Printf("    Anonymous struct: %+v\n", s3)

	// Function types (ast.FuncType)
	fmt.Println("  ✓ ast.FuncType:")
	type SimpleFunc func(int) int
	type VoidFunc func()
	type MultiFunc func(int, string) (bool, error)
	type VariadicFunc func(...int) int

	var f1 SimpleFunc = func(x int) int { return x * 2 }
	var f2 VoidFunc = func() { fmt.Println("    VoidFunc called") }
	var f3 MultiFunc = func(i int, s string) (bool, error) { return true, nil }
	var f4 VariadicFunc = func(nums ...int) int { return len(nums) }

	fmt.Printf("    SimpleFunc: %d\n", f1(5))
	f2()
	b, _ := f3(1, "test")
	fmt.Printf("    MultiFunc: %v\n", b)
	fmt.Printf("    VariadicFunc: %d args\n", f4(1, 2, 3))

	// Interface types (ast.InterfaceType)
	fmt.Println("  ✓ ast.InterfaceType:")
	type Reader interface {
		Read() string
	}
	type Writer interface {
		Write(string)
	}
	type ReadWriter interface {
		Reader
		Writer
	}
	// Empty interface
	var iface interface{} = "anything"
	fmt.Printf("    interface{}: %v\n", iface)

	// Map types (ast.MapType)
	fmt.Println("  ✓ ast.MapType:")
	var map1 map[string]int = map[string]int{"a": 1, "b": 2}
	var map2 map[int]string = map[int]string{1: "one", 2: "two"}
	var map3 map[string]map[string]int = map[string]map[string]int{
		"outer": {"inner": 42},
	}
	fmt.Printf("    map[string]int: %v\n", map1)
	fmt.Printf("    map[int]string: %v\n", map2)
	fmt.Printf("    map[string]map[string]int: %v\n", map3)

	// Channel types (ast.ChanType)
	fmt.Println("  ✓ ast.ChanType:")
	var ch1 chan int = make(chan int, 1)
	var ch2 chan<- int = ch1       // send-only
	var ch3 <-chan int = ch1       // receive-only
	ch2 <- 42
	val := <-ch3
	fmt.Printf("    chan int: %T\n", ch1)
	fmt.Printf("    chan<- int (send-only): %T\n", ch2)
	fmt.Printf("    <-chan int (receive-only): %T, received: %d\n", ch3, val)

	// Combined types
	fmt.Println("  ✓ Combined complex types:")
	type ComplexType struct {
		PtrField  *int
		SliceField []string
		MapField  map[string]int
		ChanField chan bool
		FuncField func(int) int
	}
	num := 100
	ct := ComplexType{
		PtrField:  &num,
		SliceField: []string{"a", "b"},
		MapField:  map[string]int{"x": 1},
		ChanField: make(chan bool, 1),
		FuncField: func(x int) int { return x },
	}
	fmt.Printf("    ComplexType.PtrField: %d\n", *ct.PtrField)
	fmt.Printf("    ComplexType.SliceField: %v\n", ct.SliceField)
	fmt.Printf("    ComplexType.MapField: %v\n", ct.MapField)
	fmt.Printf("    ComplexType: contains all major types\n")

	fmt.Println("Summary: All type AST nodes exercised")
	fmt.Println("Primary AST Nodes: ast.Ident, ast.ArrayType, ast.StructType, ast.FuncType,")
	fmt.Println("                   ast.InterfaceType, ast.MapType, ast.ChanType, ast.StarExpr")
	fmt.Println("========================================")
}
