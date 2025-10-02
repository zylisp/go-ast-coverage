// Package main demonstrates generics/type parameters AST nodes (Go 1.18+).
// This file exercises ast.IndexListExpr and type parameter nodes.
package main

import (
	"fmt"
)

// AST Nodes Covered:
// - ast.IndexListExpr (type parameter instantiation)
// - ast.FieldList (type parameters)
// - Type constraints in interfaces

// Generic type with single type parameter
type Box[T any] struct {
	value T
}

// Generic type with multiple type parameters
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Generic type with constraint
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Calculator[T Number] struct {
	value T
}

// Generic function with single type parameter
func Identity[T any](val T) T {
	return val
}

// Generic function with multiple type parameters
func MakePair[K comparable, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: value}
}

// Generic function with constraint
func Add[T Number](a, b T) T {
	return a + b
}

// Generic function with interface constraint
func Print[T fmt.Stringer](val T) {
	fmt.Println(val.String())
}

// Comparable constraint
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Generic slice function
func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Generic max function
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Type set with approximation (~)
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Generic method
func (b *Box[T]) Set(val T) {
	b.value = val
}

func (b Box[T]) Get() T {
	return b.value
}

// Generic interface
type Container[T any] interface {
	Add(T)
	Get() T
}

// Implementation of generic interface
type SliceContainer[T any] struct {
	items []T
}

func (s *SliceContainer[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s *SliceContainer[T]) Get() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[0]
}

func funcMain() {
	fmt.Println("=== generics.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Generic type instantiation (ast.IndexListExpr or ast.IndexExpr)
	fmt.Println("  ✓ Generic type with single type parameter:")
	intBox := Box[int]{value: 42}
	strBox := Box[string]{value: "hello"}
	fmt.Printf("    Box[int]: %+v\n", intBox)
	fmt.Printf("    Box[string]: %+v\n", strBox)

	// Generic type with multiple type parameters
	fmt.Println("  ✓ Generic type with multiple type parameters:")
	pair1 := Pair[string, int]{Key: "age", Value: 30}
	pair2 := Pair[int, string]{Key: 1, Value: "first"}
	fmt.Printf("    Pair[string, int]: %+v\n", pair1)
	fmt.Printf("    Pair[int, string]: %+v\n", pair2)

	// Generic function calls
	fmt.Println("  ✓ Generic function instantiation:")
	v1 := Identity[int](42)
	v2 := Identity[string]("test")
	fmt.Printf("    Identity[int](42): %d\n", v1)
	fmt.Printf("    Identity[string](\"test\"): %s\n", v2)

	// Type inference (no explicit type parameters)
	fmt.Println("  ✓ Generic function with type inference:")
	v3 := Identity(100)     // infers int
	v4 := Identity("infer") // infers string
	fmt.Printf("    Identity(100): %d (inferred int)\n", v3)
	fmt.Printf("    Identity(\"infer\"): %s (inferred string)\n", v4)

	// Generic function with multiple type parameters
	fmt.Println("  ✓ Generic function with multiple type parameters:")
	p1 := MakePair[string, int]("count", 5)
	p2 := MakePair("name", "Alice") // type inference
	fmt.Printf("    MakePair[string, int]: %+v\n", p1)
	fmt.Printf("    MakePair (inferred): %+v\n", p2)

	// Generic function with constraint
	fmt.Println("  ✓ Generic function with Number constraint:")
	sum1 := Add[int](5, 3)
	sum2 := Add[float64](3.14, 2.86)
	fmt.Printf("    Add[int](5, 3): %d\n", sum1)
	fmt.Printf("    Add[float64](3.14, 2.86): %f\n", sum2)

	// Comparable constraint
	fmt.Println("  ✓ Generic function with comparable constraint:")
	eq1 := Equal[int](5, 5)
	eq2 := Equal[string]("hi", "bye")
	fmt.Printf("    Equal[int](5, 5): %v\n", eq1)
	fmt.Printf("    Equal[string](\"hi\", \"bye\"): %v\n", eq2)

	// Generic higher-order function
	fmt.Println("  ✓ Generic higher-order function (Map):")
	ints := []int{1, 2, 3, 4, 5}
	doubled := Map(ints, func(x int) int { return x * 2 })
	fmt.Printf("    Map(ints, double): %v\n", doubled)

	strs := []string{"a", "b", "c"}
	lengths := Map(strs, func(s string) int { return len(s) })
	fmt.Printf("    Map(strings, len): %v\n", lengths)

	// Generic max
	fmt.Println("  ✓ Generic Max with Number constraint:")
	maxInt := Max(10, 20)
	maxFloat := Max(3.14, 2.71)
	fmt.Printf("    Max(10, 20): %d\n", maxInt)
	fmt.Printf("    Max(3.14, 2.71): %f\n", maxFloat)

	// Generic methods
	fmt.Println("  ✓ Generic type methods:")
	box := Box[int]{value: 10}
	fmt.Printf("    box.Get(): %d\n", box.Get())
	box.Set(20)
	fmt.Printf("    After box.Set(20): %d\n", box.Get())

	// Generic interface
	fmt.Println("  ✓ Generic interface:")
	var container Container[string] = &SliceContainer[string]{}
	container.Add("first")
	container.Add("second")
	fmt.Printf("    Container[string].Get(): %s\n", container.Get())

	// Type constraints with union
	fmt.Println("  ✓ Type constraints with union (|):")
	fmt.Println("    Number interface defined with union of numeric types")

	// Type constraints with approximation (~)
	fmt.Println("  ✓ Type constraints with approximation (~):")
	fmt.Println("    Integer interface uses ~int | ~int8 | ...")

	// Nested generic types
	fmt.Println("  ✓ Nested generic types:")
	nestedBox := Box[Box[int]]{value: Box[int]{value: 42}}
	fmt.Printf("    Box[Box[int]]: inner value = %d\n", nestedBox.value.value)

	// Generic slice types
	fmt.Println("  ✓ Generic with slice types:")
	sliceBox := Box[[]int]{value: []int{1, 2, 3}}
	fmt.Printf("    Box[[]int]: %v\n", sliceBox.value)

	// Generic map types
	mapBox := Box[map[string]int]{value: map[string]int{"a": 1}}
	fmt.Printf("    Box[map[string]int]: %v\n", mapBox.value)

	// any constraint (alias for interface{})
	fmt.Println("  ✓ any constraint (builtin):")
	fmt.Println("    Used in Box[T any] and Identity[T any]")

	fmt.Println("Summary: Comprehensive generics AST node coverage (Go 1.18+)")
	fmt.Println("Primary AST Nodes: ast.IndexListExpr, ast.FieldList (type parameters)")
	fmt.Println("Features: type parameters, constraints, inference, generic types/functions/methods")
	fmt.Println("========================================")
}

func main() {
	funcMain()
}
