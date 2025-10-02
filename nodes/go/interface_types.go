// Package main demonstrates interface type AST nodes in detail.
// This file exercises ast.InterfaceType nodes.
package main

import "fmt"

// AST Nodes Covered:
// - ast.InterfaceType
// - ast.Field (interface methods)
// - ast.FieldList (method set)

// Empty interface
type Any interface{}

// Simple interface with one method
type Reader interface{
	Read() string
}

// Interface with multiple methods
type ReadWriter interface {
	Read() string
	Write(string) error
}

// Interface with parameter names
type Calculator interface {
	Add(a, b int) int
	Subtract(x, y int) int
	Multiply(int, int) int // no parameter names
}

// Embedded interfaces
type Closer interface {
	Close() error
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type Writer interface {
	Write(string) error
}

// Interface with variadic method
type Logger interface {
	Log(format string, args ...interface{})
}

// Interface with generic constraints (Go 1.18+)
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

// Interface with type constraint and method
type Stringable interface {
	~string | ~int | ~bool
	String() string
}

// Comparable interface (built-in in Go 1.18+)
type MyComparable interface {
	comparable
}

// Interface with any constraint
type AnyType interface {
	any
}

func funcMain() {
	fmt.Println("=== interface_types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Empty interface
	fmt.Println("  ✓ ast.InterfaceType (empty):")
	var a Any = 42
	var b Any = "string"
	var c Any = true
	fmt.Printf("    Any can hold int: %v\n", a)
	fmt.Printf("    Any can hold string: %v\n", b)
	fmt.Printf("    Any can hold bool: %v\n", c)

	// Simple interface
	fmt.Println("  ✓ ast.InterfaceType (single method):")
	var r Reader = StringReader{content: "Hello, World!"}
	fmt.Printf("    Reader.Read(): %s\n", r.Read())

	// Multiple methods
	fmt.Println("  ✓ ast.InterfaceType (multiple methods):")
	var rw ReadWriter = &Buffer{data: "initial"}
	fmt.Printf("    ReadWriter.Read(): %s\n", rw.Read())
	err := rw.Write("modified")
	fmt.Printf("    ReadWriter.Write(): error=%v, result=%s\n", err, rw.Read())

	// Interface with parameter names
	fmt.Println("  ✓ ast.InterfaceType with ast.Field (named parameters):")
	var calc Calculator = BasicCalculator{}
	fmt.Printf("    Calculator.Add(5, 3): %d\n", calc.Add(5, 3))
	fmt.Printf("    Calculator.Subtract(10, 4): %d\n", calc.Subtract(10, 4))
	fmt.Printf("    Calculator.Multiply(6, 7): %d\n", calc.Multiply(6, 7))

	// Embedded interfaces
	fmt.Println("  ✓ ast.InterfaceType (embedded):")
	var rwc ReadWriteCloser = &FileHandle{
		content: "file content",
		open:    true,
	}
	fmt.Printf("    ReadWriteCloser.Read(): %s\n", rwc.Read())
	_ = rwc.Write("new content")
	fmt.Printf("    After Write: %s\n", rwc.Read())
	_ = rwc.Close()
	fmt.Printf("    ReadWriteCloser.Close(): file closed\n")

	// Variadic method
	fmt.Println("  ✓ ast.InterfaceType (variadic method):")
	var logger Logger = SimpleLogger{}
	logger.Log("Message: %s, Number: %d", "test", 42)

	// Type assertion
	fmt.Println("  ✓ Interface type assertion:")
	var iface interface{} = "hello"
	str, ok := iface.(string)
	fmt.Printf("    Type assertion (string): %q, ok=%v\n", str, ok)
	num, ok := iface.(int)
	fmt.Printf("    Type assertion (int): %d, ok=%v\n", num, ok)

	// Type switch
	fmt.Println("  ✓ Interface type switch:")
	describeType(42)
	describeType("hello")
	describeType(true)
	describeType(3.14)

	// Interface value with nil
	fmt.Println("  ✓ Interface nil handling:")
	var nilReader Reader = nil
	fmt.Printf("    nil interface: %v, is nil: %v\n", nilReader, nilReader == nil)

	// Interface with nil concrete value
	var nilPtr *StringReader = nil
	var readerWithNil Reader = nilPtr
	fmt.Printf("    Interface with nil concrete: is nil: %v\n", readerWithNil == nil)

	// Dynamic dispatch
	fmt.Println("  ✓ Interface dynamic dispatch:")
	readers := []Reader{
		StringReader{content: "first"},
		StringReader{content: "second"},
		StringReader{content: "third"},
	}
	for i, reader := range readers {
		fmt.Printf("    Reader[%d]: %s\n", i, reader.Read())
	}

	// Interface satisfaction (implicit implementation)
	fmt.Println("  ✓ Interface implicit implementation:")
	var w Writer = &Buffer{data: "test"}
	_ = w.Write("implicit implementation works")
	fmt.Printf("    Writer interface satisfied implicitly\n")

	// Comparable constraint (Go 1.18+)
	fmt.Println("  ✓ ast.InterfaceType (comparable constraint):")
	fmt.Printf("    comparable constraint defined\n")

	// Type sets (Go 1.18+)
	fmt.Println("  ✓ ast.InterfaceType (type sets/unions):")
	fmt.Printf("    Ordered interface with type union defined\n")

	fmt.Println("Summary: Comprehensive interface type AST node coverage")
	fmt.Println("Primary AST Nodes: ast.InterfaceType, ast.Field, ast.FieldList")
	fmt.Println("Features: empty interfaces, method sets, embedding, generics, type constraints")
	fmt.Println("========================================")
}

// Implementations for testing

type StringReader struct {
	content string
}

func (s StringReader) Read() string {
	return s.content
}

type Buffer struct {
	data string
}

func (b *Buffer) Read() string {
	return b.data
}

func (b *Buffer) Write(s string) error {
	b.data = s
	return nil
}

type BasicCalculator struct{}

func (c BasicCalculator) Add(a, b int) int {
	return a + b
}

func (c BasicCalculator) Subtract(x, y int) int {
	return x - y
}

func (c BasicCalculator) Multiply(a, b int) int {
	return a * b
}

type FileHandle struct {
	content string
	open    bool
}

func (f *FileHandle) Read() string {
	return f.content
}

func (f *FileHandle) Write(s string) error {
	f.content = s
	return nil
}

func (f *FileHandle) Close() error {
	f.open = false
	return nil
}

type SimpleLogger struct{}

func (l SimpleLogger) Log(format string, args ...interface{}) {
	fmt.Printf("    Log: "+format+"\n", args...)
}

func describeType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("    Type switch: int(%d)\n", v)
	case string:
		fmt.Printf("    Type switch: string(%q)\n", v)
	case bool:
		fmt.Printf("    Type switch: bool(%v)\n", v)
	default:
		fmt.Printf("    Type switch: unknown(%v)\n", v)
	}
}

func main() {
	funcMain()
}
