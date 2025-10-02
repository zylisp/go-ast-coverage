// Package main demonstrates basic literal AST nodes.
// This file exercises ast.BasicLit for all primitive types.
package main

import "fmt"

// AST Nodes Covered:
// - ast.BasicLit (INT, FLOAT, IMAG, CHAR, STRING)
// - ast.Ident
// - ast.ValueSpec
// - ast.GenDecl

func main() {
	fmt.Println("=== basic_literals.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Integer literals (ast.BasicLit INT)
	var intDec int = 42
	var intHex int = 0x2A
	var intOct int = 052
	var intBin int = 0b101010
	fmt.Printf("  ✓ ast.BasicLit (INT decimal): %d\n", intDec)
	fmt.Printf("  ✓ ast.BasicLit (INT hex): 0x%X = %d\n", intHex, intHex)
	fmt.Printf("  ✓ ast.BasicLit (INT octal): 0%o = %d\n", intOct, intOct)
	fmt.Printf("  ✓ ast.BasicLit (INT binary): 0b%b = %d\n", intBin, intBin)

	// Integer type variations
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	fmt.Printf("  ✓ ast.BasicLit (INT8): %d\n", i8)
	fmt.Printf("  ✓ ast.BasicLit (INT16): %d\n", i16)
	fmt.Printf("  ✓ ast.BasicLit (INT32): %d\n", i32)
	fmt.Printf("  ✓ ast.BasicLit (INT64): %d\n", i64)
	fmt.Printf("  ✓ ast.BasicLit (UINT8): %d\n", u8)
	fmt.Printf("  ✓ ast.BasicLit (UINT16): %d\n", u16)
	fmt.Printf("  ✓ ast.BasicLit (UINT32): %d\n", u32)
	fmt.Printf("  ✓ ast.BasicLit (UINT64): %d\n", u64)

	// Floating point literals (ast.BasicLit FLOAT)
	var f32 float32 = 3.14159
	var f64 float64 = 2.718281828459045
	var floatExp float64 = 1.23e10
	var floatHex float64 = 0x1.fp+1
	fmt.Printf("  ✓ ast.BasicLit (FLOAT32): %f\n", f32)
	fmt.Printf("  ✓ ast.BasicLit (FLOAT64): %f\n", f64)
	fmt.Printf("  ✓ ast.BasicLit (FLOAT exp): %e\n", floatExp)
	fmt.Printf("  ✓ ast.BasicLit (FLOAT hex): %f\n", floatHex)

	// Complex number literals (ast.BasicLit IMAG)
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3.14 + 2.71i
	var imagOnly = 5i
	fmt.Printf("  ✓ ast.BasicLit (COMPLEX64): %v\n", c64)
	fmt.Printf("  ✓ ast.BasicLit (COMPLEX128): %v\n", c128)
	fmt.Printf("  ✓ ast.BasicLit (IMAG): %v\n", imagOnly)

	// Character literals (ast.BasicLit CHAR)
	var char rune = 'A'
	var charUnicode rune = '\u0041'
	var charEscape rune = '\n'
	var byteChar byte = 'Z'
	fmt.Printf("  ✓ ast.BasicLit (CHAR): '%c' = %d\n", char, char)
	fmt.Printf("  ✓ ast.BasicLit (CHAR unicode): '%c' = U+%04X\n", charUnicode, charUnicode)
	fmt.Printf("  ✓ ast.BasicLit (CHAR escape): %q\n", charEscape)
	fmt.Printf("  ✓ ast.BasicLit (BYTE): '%c'\n", byteChar)

	// String literals (ast.BasicLit STRING)
	var str1 string = "Hello, World!"
	var str2 string = `Raw string
with newlines`
	var strUnicode string = "Hello, 世界"
	var strEscape string = "Line1\nLine2\tTabbed"
	fmt.Printf("  ✓ ast.BasicLit (STRING interpreted): %q\n", str1)
	fmt.Printf("  ✓ ast.BasicLit (STRING raw): %q\n", str2)
	fmt.Printf("  ✓ ast.BasicLit (STRING unicode): %q\n", strUnicode)
	fmt.Printf("  ✓ ast.BasicLit (STRING escape): %q\n", strEscape)

	// Boolean literals (ast.Ident for true/false)
	var boolTrue bool = true
	var boolFalse bool = false
	fmt.Printf("  ✓ ast.Ident (bool true): %v\n", boolTrue)
	fmt.Printf("  ✓ ast.Ident (bool false): %v\n", boolFalse)

	// Nil literal (ast.Ident)
	var nilPtr *int = nil
	fmt.Printf("  ✓ ast.Ident (nil): %v\n", nilPtr)

	// Underscores in numeric literals (Go 1.13+)
	var largeNum int = 1_000_000
	var largeFloat float64 = 1_234.567_89
	fmt.Printf("  ✓ ast.BasicLit (INT with underscores): %d\n", largeNum)
	fmt.Printf("  ✓ ast.BasicLit (FLOAT with underscores): %f\n", largeFloat)

	fmt.Println("Summary: 35+ unique literal variations exercised")
	fmt.Println("Primary AST Nodes: ast.BasicLit, ast.Ident, ast.ValueSpec, ast.GenDecl")
	fmt.Println("========================================")
}
