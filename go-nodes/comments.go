// Package main demonstrates comment AST nodes.
// This file exercises ast.Comment and ast.CommentGroup nodes.
//
// Multi-line package documentation comment.
// This is line 3 of the package doc.
package main

import "fmt"

// AST Nodes Covered:
// - ast.Comment
// - ast.CommentGroup
// - ast.File.Comments

/*
This is a block comment for demonstrating ast.Comment.
It can span multiple lines.
Block comments are useful for longer explanations.
*/

// Constant with documentation comment (ast.CommentGroup)
const DocumentedConst = 42

/*
MultilineDocConst has a block comment.
This demonstrates ast.CommentGroup with block style.
*/
const MultilineDocConst = 100

// Variable with doc comment
var DocumentedVar = "value"

// Type with documentation
type DocumentedType int

/*
DocumentedStruct is a struct with documentation.
*/
type DocumentedStruct struct {
	// Field1 has a line comment
	Field1 int

	/*
		Field2 has a block comment.
		It spans multiple lines.
	*/
	Field2 string

	Field3 bool // Inline comment after field
}

// DocumentedFunc is a function with documentation.
// It demonstrates ast.CommentGroup for functions.
//
// This function has multi-paragraph documentation.
// The second paragraph provides more detail.
func DocumentedFunc(x int) int {
	// Internal comment (not part of doc)
	y := x * 2

	/*
		Block comment inside function.
		These are regular comments, not documentation.
	*/
	return y
}

/*
BlockDocFunc has block-style documentation.

It can include blank lines.
*/
func BlockDocFunc() string {
	return "documented"
}

func funcMain() {
	fmt.Println("=== comments.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	fmt.Println("  ✓ ast.Comment (line comments //)")
	fmt.Println("  ✓ ast.Comment (block comments /* */)")
	fmt.Println("  ✓ ast.CommentGroup (documentation comments)")
	fmt.Println("  ✓ Package-level doc comments")
	fmt.Println("  ✓ Constant doc comments")
	fmt.Println("  ✓ Variable doc comments")
	fmt.Println("  ✓ Type doc comments")
	fmt.Println("  ✓ Struct field comments")
	fmt.Println("  ✓ Function doc comments")
	fmt.Println("  ✓ Inline comments")
	fmt.Println("  ✓ Internal comments (non-doc)")
	fmt.Println("  ✓ Multi-paragraph doc comments")

	// Line comment
	x := 10 // Inline comment

	// Another line comment
	y := 20

	/*
		Block comment in code.
	*/
	z := x + y

	// Comment before control structure
	if z > 0 {
		// Comment inside if block
		fmt.Printf("    z = %d\n", z)
	}

	// Comment before loop
	for i := 0; i < 3; i++ {
		// Comment inside loop
		_ = i
	}

	// Comment before switch
	switch z {
	case 30: // Comment on case
		// Comment inside case
		fmt.Printf("    z is 30\n")
	default:
		// Default case comment
	}

	// Comment groups at different levels
	fmt.Println("  ✓ Comment positions:")
	fmt.Println("    - Package level")
	fmt.Println("    - Declaration level")
	fmt.Println("    - Field level")
	fmt.Println("    - Statement level")
	fmt.Println("    - Inline")

	// Special comment formats
	// TODO: This is a TODO comment
	// FIXME: This is a FIXME comment
	// NOTE: This is a NOTE comment
	// BUG: This is a BUG comment
	fmt.Println("  ✓ Special comment markers (TODO, FIXME, NOTE, BUG)")

	// Comments with special characters
	// Comment with symbols: !@#$%^&*()
	// Comment with URL: https://example.com
	// Comment with code: `backticks`
	fmt.Println("  ✓ Comments with special characters")

	// Documentation example
	// Example:
	//   result := DocumentedFunc(10)
	//   fmt.Println(result)
	fmt.Println("  ✓ Documentation with examples")

	result := DocumentedFunc(5)
	str := BlockDocFunc()
	fmt.Printf("    DocumentedFunc(5) = %d\n", result)
	fmt.Printf("    BlockDocFunc() = %s\n", str)

	// Struct with comments
	s := DocumentedStruct{
		Field1: 1,    // Comment on initialization
		Field2: "hi", // Another init comment
		Field3: true,
	}
	fmt.Printf("    DocumentedStruct: %+v\n", s)

	fmt.Println("Summary: Comprehensive comment AST node coverage")
	fmt.Println("Primary AST Nodes: ast.Comment, ast.CommentGroup")
	fmt.Println("Features: line comments, block comments, doc comments, inline comments")
	fmt.Println("========================================")
}

// Method with documentation
func (d DocumentedType) Method() int {
	// Method implementation comment
	return int(d) * 2
}

/*
BlockMethod has block-style method documentation.
*/
func (d DocumentedType) BlockMethod() string {
	return "block"
}

func main() {
	funcMain()
}
