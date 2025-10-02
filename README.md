# Go AST Coverage Test Suite

A comprehensive test suite that demonstrates and exercises every AST (Abstract Syntax Tree) node type defined in Go's `go/ast` package. This project serves as both a reference implementation and a testing framework for Go AST parsing tools.

## Overview

This project provides:

- **Complete AST node coverage**: Test files covering all expression, statement, declaration, type, and structural nodes
- **AST analyzer**: Tool to inspect and enumerate AST nodes in Go source files
- **Coverage reporting**: Comprehensive reports showing which AST nodes are covered and which are missing
- **Executable examples**: Every test file is independently runnable and produces meaningful output

## Project Structure

```
go-ast-coverage/
├── README.md                    # This file
├── LICENSE                      # Project license
├── go.mod                       # Go module definition
├── main.go                      # Main orchestrator
├── go-nodes/                    # AST node test files
│   ├── basic_literals.go        # BasicLit, primitive types
│   ├── identifiers.go           # Ident, Object, Scope
│   ├── expressions.go           # Expression nodes (Binary, Unary, Call, etc.)
│   ├── statements.go            # Statement nodes (Assign, If, For, etc.)
│   ├── control_flow.go          # Control flow (If, Switch, For, Range, Select)
│   ├── declarations.go          # Declaration nodes (GenDecl, FuncDecl, etc.)
│   ├── types.go                 # Type nodes overview
│   ├── struct_types.go          # StructType, Field, Tag
│   ├── interface_types.go       # InterfaceType, method sets
│   ├── function_types.go        # FuncType, FuncLit, closures
│   ├── array_slice_types.go     # ArrayType, SliceExpr
│   ├── map_channel_types.go     # MapType, ChanType
│   ├── comments.go              # Comment, CommentGroup
│   ├── imports.go               # ImportSpec, import patterns
│   ├── generics.go              # IndexListExpr, type parameters (Go 1.18+)
│   └── edge_cases.go            # Edge cases and special constructs
├── ast-analyzer/
│   └── analyzer.go              # AST inspection and analysis utilities
└── coverage-report/
    └── report.go                # Coverage report generation
```

## Quick Start

### Prerequisites

- Go 1.21 or later (for generics and modern Go features)
- Standard Go toolchain

### Installation

```bash
# Clone the repository
git clone https://github.com/zylisp/go-ast-coverage.git
cd go-ast-coverage

# Verify Go module
go mod tidy
```

### Running the Test Suite

```bash
# Run everything (tests + analysis + report)
go run main.go

# Run only the test files
go run main.go -run

# Generate coverage report only
go run main.go -report

# Analyze AST nodes only
go run main.go -analyze

# Verbose output
go run main.go -verbose

# Save report as JSON
go run main.go -report -json
```

### Running Individual Test Files

Each test file in `go-nodes/` can be run independently:

```bash
# Run a specific test file
go run go-nodes/basic_literals.go
go run go-nodes/expressions.go
go run go-nodes/generics.go

# Run any test file
cd go-nodes
go run <filename>.go
```

## AST Node Coverage

This test suite achieves **94.64% coverage (53 of 56 node types)** of all AST node types defined in Go's `go/ast` package. The three uncovered nodes are error recovery nodes that require invalid syntax (see note below).

### Expression Nodes (ast.Expr)

- ✗ `*ast.BadExpr` - Error recovery node (not covered - see note below)
- ✓ `*ast.Ident` - Identifiers
- ✓ `*ast.Ellipsis` - `...` in variadic functions
- ✓ `*ast.BasicLit` - Basic literals (int, float, string, char, etc.)
- ✓ `*ast.FuncLit` - Function literals (closures)
- ✓ `*ast.CompositeLit` - Composite literals (struct, array, slice, map)
- ✓ `*ast.ParenExpr` - Parenthesized expressions
- ✓ `*ast.SelectorExpr` - Field/method selectors (x.f)
- ✓ `*ast.IndexExpr` - Index expressions (x[i])
- ✓ `*ast.IndexListExpr` - Type parameter instantiation (Go 1.18+)
- ✓ `*ast.SliceExpr` - Slice expressions (x[i:j])
- ✓ `*ast.TypeAssertExpr` - Type assertions (x.(T))
- ✓ `*ast.CallExpr` - Function calls
- ✓ `*ast.StarExpr` - Pointer types and dereference
- ✓ `*ast.UnaryExpr` - Unary expressions (+, -, !, ^, &, <-)
- ✓ `*ast.BinaryExpr` - Binary expressions (+, -, *, /, etc.)
- ✓ `*ast.KeyValueExpr` - Key-value pairs in composite literals

### Statement Nodes (ast.Stmt)

- ✗ `*ast.BadStmt` - Error recovery node (not covered - see note below)
- ✓ `*ast.DeclStmt` - Declarations in function bodies
- ✓ `*ast.EmptyStmt` - Empty statements
- ✓ `*ast.LabeledStmt` - Labeled statements
- ✓ `*ast.ExprStmt` - Expression statements
- ✓ `*ast.SendStmt` - Channel send statements
- ✓ `*ast.IncDecStmt` - Increment/decrement (++ and --)
- ✓ `*ast.AssignStmt` - Assignment statements
- ✓ `*ast.GoStmt` - Go statement (goroutines)
- ✓ `*ast.DeferStmt` - Defer statements
- ✓ `*ast.ReturnStmt` - Return statements
- ✓ `*ast.BranchStmt` - Branch statements (break, continue, goto, fallthrough)
- ✓ `*ast.BlockStmt` - Block statements
- ✓ `*ast.IfStmt` - If statements
- ✓ `*ast.CaseClause` - Case clauses in switch/select
- ✓ `*ast.SwitchStmt` - Switch statements
- ✓ `*ast.TypeSwitchStmt` - Type switch statements
- ✓ `*ast.CommClause` - Communication clauses in select
- ✓ `*ast.SelectStmt` - Select statements
- ✓ `*ast.ForStmt` - For loops
- ✓ `*ast.RangeStmt` - Range loops

### Declaration Nodes (ast.Decl)

- ✗ `*ast.BadDecl` - Error recovery node (not covered - see note below)
- ✓ `*ast.GenDecl` - General declarations (import, const, type, var)
- ✓ `*ast.FuncDecl` - Function declarations

### Spec Nodes (ast.Spec)

- ✓ `*ast.ImportSpec` - Import specifications
- ✓ `*ast.ValueSpec` - Constant and variable specifications
- ✓ `*ast.TypeSpec` - Type specifications

### Type Nodes

- ✓ `*ast.ArrayType` - Array and slice types
- ✓ `*ast.StructType` - Struct types
- ✓ `*ast.FuncType` - Function types
- ✓ `*ast.InterfaceType` - Interface types
- ✓ `*ast.MapType` - Map types
- ✓ `*ast.ChanType` - Channel types

### Other Nodes

- ✓ `*ast.File` - Source file node
- ✓ `*ast.Package` - Package node
- ✓ `*ast.Comment` - Individual comments
- ✓ `*ast.CommentGroup` - Comment groups
- ✓ `*ast.Field` - Struct fields, function parameters
- ✓ `*ast.FieldList` - Field lists

### Coverage Note: Error Recovery Nodes

**Current Coverage: 53 of 56 node types (94.64%)**

Three AST node types are intentionally not covered in this test suite:

- `*ast.BadExpr` - Error recovery node for expressions
- `*ast.BadStmt` - Error recovery node for statements
- `*ast.BadDecl` - Error recovery node for declarations

**Why aren't these covered?**

These "Bad*" nodes are error recovery nodes created by Go's parser when it encounters syntax errors. They allow the parser to continue parsing despite errors, which is useful for tools like IDEs that need to work with incomplete or invalid code.

However, this test suite is designed around a fundamental requirement: **all test files must compile and execute successfully**. Since Bad* nodes only appear in syntactically invalid Go code, they cannot be demonstrated in a compilable test file.

To generate these nodes, you would need:
```go
// This won't compile!
func badExample() {
    x := 5 +      // BadExpr: incomplete expression
    if {          // BadStmt: invalid if statement
    const         // BadDecl: incomplete declaration
}
```

**Achieving this coverage would require:**
1. Creating intentionally invalid Go files
2. Removing the compilation requirement from the test suite
3. Special parsing logic to handle syntax errors

Since the value of covering error recovery nodes is minimal compared to the cost of abandoning the "all files compile" principle, we've chosen to maintain 94.64% coverage with all valid, executable Go code.

## Features Demonstrated

### Go Language Features

- All primitive types (int variants, uint variants, float, complex, string, bool, byte, rune)
- Variables, constants, type definitions
- Functions (regular, variadic, methods, anonymous, generic)
- Arrays, slices, maps, channels
- Structs, interfaces, embedding
- Pointers and pointer operations
- Control flow (if, switch, for, range, select)
- Goroutines and channels
- Defer, panic, recover
- Generics and type parameters (Go 1.18+)
- Type constraints and type sets
- Comments and documentation

### Modern Go Features (1.18+)

- Generic types and functions
- Type parameters with constraints
- Type inference
- Type sets and unions in interfaces
- `any` and `comparable` built-in constraints

## Usage Examples

### Analyzing a Single File

```go
import "zylisp/go-ast-coverage/ast-analyzer"

result, err := analyzer.AnalyzeFile("path/to/file.go")
if err != nil {
    log.Fatal(err)
}
analyzer.PrintAnalysis(result)
```

### Generating Coverage Report

```go
import "zylisp/go-ast-coverage/coverage-report"

report, err := report.GenerateReport("go-nodes")
if err != nil {
    log.Fatal(err)
}
report.PrintReport(report)
```

## Output Examples

### Test File Output

Each test file produces structured output showing which AST nodes it exercises:

```
=== basic_literals.go AST Node Coverage ===
Exercising AST Nodes:
  ✓ ast.BasicLit (INT decimal): 42
  ✓ ast.BasicLit (INT hex): 0x2A = 42
  ✓ ast.BasicLit (FLOAT64): 3.141590
  ✓ ast.BasicLit (STRING): "Hello, World!"
  ✓ ast.Ident (bool): true
Summary: 35+ unique literal variations exercised
========================================
```

### Coverage Report

```
================================================================================
GO AST COVERAGE REPORT
================================================================================
Generated: 2024-10-02T...

SUMMARY
--------------------------------------------------------------------------------
Total AST Node Types:    68
Covered Node Types:      68
Missing Node Types:      0
Coverage:                100.00%

Progress: [██████████████████████████████████████████████████] 100.00%

🎉 PERFECT COVERAGE! All AST node types are covered!
================================================================================
```

## Development

### Adding New Test Files

1. Create a new `.go` file in `go-nodes/`
2. Ensure it has `package main` and a `main()` function
3. Add documentation comments listing covered AST nodes
4. Make the file independently runnable
5. Include structured output showing which nodes are exercised
6. Test with `go run go-nodes/your_file.go`

### Test File Template

```go
// Package main demonstrates X AST nodes.
// This file exercises ast.NodeType1, ast.NodeType2, etc.
package main

import "fmt"

// AST Nodes Covered:
// - ast.NodeType1
// - ast.NodeType2
// - ast.NodeType3

func main() {
    fmt.Println("=== your_file.go AST Node Coverage ===")
    fmt.Println("Exercising AST Nodes:")

    // Your code here demonstrating AST nodes

    fmt.Println("Summary: X unique AST node types exercised")
    fmt.Println("========================================")
}
```

## Contributing

Contributions are welcome! Please ensure:

- All new files compile and run successfully
- Output follows the established format
- Documentation is comprehensive
- Code follows Go best practices

## License

See LICENSE file for details.

## References

- [Go AST Package Documentation](https://pkg.go.dev/go/ast)
- [Go Language Specification](https://go.dev/ref/spec)
- [Go Generics Proposal](https://go.dev/blog/intro-generics)

## Acknowledgments

This project was created to provide a comprehensive reference for Go AST nodes and to assist developers working with Go's abstract syntax tree.
