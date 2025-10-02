// Package main demonstrates array and slice AST nodes.
// This file exercises ast.ArrayType nodes in detail.
package main

import "fmt"

// AST Nodes Covered:
// - ast.ArrayType (arrays and slices)
// - ast.Ellipsis (in array literals)
// - ast.SliceExpr
// - ast.IndexExpr

func funcMain() {
	fmt.Println("=== array_slice_types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Array declarations and initialization
	fmt.Println("  ✓ ast.ArrayType (fixed-size arrays):")
	var arr1 [5]int
	fmt.Printf("    Zero-value [5]int: %v\n", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("    Initialized [5]int: %v\n", arr2)

	arr3 := [3]string{"a", "b", "c"}
	fmt.Printf("    [3]string: %v\n", arr3)

	// Array with ellipsis (compiler counts elements)
	arr4 := [...]int{10, 20, 30}
	fmt.Printf("    [...] int: %v (length: %d)\n", arr4, len(arr4))

	// Partial initialization
	arr5 := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Printf("    Partial init [5]int: %v\n", arr5)

	// Multi-dimensional arrays
	arr6 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("    [2][3]int: %v\n", arr6)

	arr7 := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Printf("    [2][2][2]int: %v\n", arr7)

	// Array indexing (ast.IndexExpr)
	fmt.Println("  ✓ ast.IndexExpr (array indexing):")
	fmt.Printf("    arr2[0]=%d, arr2[2]=%d, arr2[4]=%d\n", arr2[0], arr2[2], arr2[4])
	fmt.Printf("    arr6[0][1]=%d, arr6[1][2]=%d\n", arr6[0][1], arr6[1][2])

	// Array assignment
	arr2[2] = 99
	fmt.Printf("    After arr2[2]=99: %v\n", arr2)

	// Array iteration
	fmt.Println("  ✓ Array iteration:")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("    arr3[%d]=%s\n", i, arr3[i])
	}

	// Array comparison
	comp1 := [3]int{1, 2, 3}
	comp2 := [3]int{1, 2, 3}
	comp3 := [3]int{1, 2, 4}
	fmt.Printf("    Array comparison: comp1==comp2: %v, comp1==comp3: %v\n", comp1 == comp2, comp1 == comp3)

	// Slice declarations
	fmt.Println("  ✓ ast.ArrayType (slices - nil Len):")
	var slice1 []int
	fmt.Printf("    Zero-value []int: %v (nil: %v, len: %d, cap: %d)\n", slice1, slice1 == nil, len(slice1), cap(slice1))

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("    Initialized []int: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))

	slice3 := []string{"x", "y", "z"}
	fmt.Printf("    []string: %v\n", slice3)

	// Slice with make
	slice4 := make([]int, 5)
	fmt.Printf("    make([]int, 5): %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))

	slice5 := make([]int, 3, 10)
	fmt.Printf("    make([]int, 3, 10): %v (len: %d, cap: %d)\n", slice5, len(slice5), cap(slice5))

	// Slice expressions (ast.SliceExpr)
	fmt.Println("  ✓ ast.SliceExpr (slicing operations):")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Simple slicing
	fmt.Printf("    s[2:5]: %v\n", s[2:5])
	fmt.Printf("    s[:4]: %v\n", s[:4])
	fmt.Printf("    s[6:]: %v\n", s[6:])
	fmt.Printf("    s[:]: %v\n", s[:])

	// Full slice expression (3-index slice)
	fullSlice := s[2:5:7]
	fmt.Printf("    s[2:5:7]: %v (len: %d, cap: %d)\n", fullSlice, len(fullSlice), cap(fullSlice))

	// Slice of array
	arrayForSlice := [5]int{10, 20, 30, 40, 50}
	sliceFromArray := arrayForSlice[1:4]
	fmt.Printf("    array[1:4]: %v\n", sliceFromArray)

	// Slice indexing and assignment
	fmt.Println("  ✓ Slice indexing (ast.IndexExpr):")
	fmt.Printf("    slice2[0]=%d, slice2[2]=%d\n", slice2[0], slice2[2])
	slice2[1] = 99
	fmt.Printf("    After slice2[1]=99: %v\n", slice2)

	// Slice append
	fmt.Println("  ✓ Slice operations:")
	slice6 := []int{1, 2, 3}
	slice6 = append(slice6, 4)
	fmt.Printf("    append(slice6, 4): %v\n", slice6)
	slice6 = append(slice6, 5, 6, 7)
	fmt.Printf("    append(slice6, 5,6,7): %v\n", slice6)

	// Append with ellipsis (ast.Ellipsis)
	slice7 := []int{10, 20}
	slice8 := []int{30, 40, 50}
	slice7 = append(slice7, slice8...)
	fmt.Printf("    append(slice7, slice8...): %v\n", slice7)

	// Slice copy
	slice9 := make([]int, len(slice2))
	n := copy(slice9, slice2)
	fmt.Printf("    copy(slice9, slice2): %v (copied %d elements)\n", slice9, n)

	// Multi-dimensional slices
	fmt.Println("  ✓ Multi-dimensional slices:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("    [][]int: %v\n", matrix)
	fmt.Printf("    matrix[1][2]: %d\n", matrix[1][2])

	// Jagged slices
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	fmt.Printf("    Jagged [][]int: %v\n", jagged)

	// Slice iteration
	fmt.Println("  ✓ Slice iteration:")
	for i, v := range []int{100, 200, 300} {
		fmt.Printf("    index=%d, value=%d\n", i, v)
	}

	// Slice of bytes (special case - string conversion)
	byteSlice := []byte("Hello")
	fmt.Printf("    []byte: %v, as string: %s\n", byteSlice, string(byteSlice))

	// Slice of runes
	runeSlice := []rune("世界")
	fmt.Printf("    []rune: %v, as string: %s\n", runeSlice, string(runeSlice))

	// Empty slice vs nil slice
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("    nil slice: %v (nil: %v)\n", nilSlice, nilSlice == nil)
	fmt.Printf("    empty slice: %v (nil: %v)\n", emptySlice, emptySlice == nil)

	// Slice with capacity growth
	growing := make([]int, 0, 2)
	fmt.Printf("    Initial: len=%d, cap=%d\n", len(growing), cap(growing))
	for i := 0; i < 5; i++ {
		growing = append(growing, i)
		fmt.Printf("    After append %d: len=%d, cap=%d\n", i, len(growing), cap(growing))
	}

	fmt.Println("Summary: Comprehensive array and slice AST node coverage")
	fmt.Println("Primary AST Nodes: ast.ArrayType, ast.IndexExpr, ast.SliceExpr, ast.Ellipsis")
	fmt.Println("Features: fixed arrays, slices, indexing, slicing, multi-dimensional, append, copy")
	fmt.Println("========================================")
}

func main() {
	funcMain()
}
