// Package main demonstrates import AST nodes.
// This file exercises ast.ImportSpec and import patterns.
package main

// AST Nodes Covered:
// - ast.GenDecl (import)
// - ast.ImportSpec
// - Various import patterns

// Single import
import "fmt"

// Grouped imports
import (
	"os"
	"strings"
	"time"
)

// Import with alias
import (
	fmtAlias "fmt"
	str "strings"
)

// Import for side effects (blank identifier)
import (
	_ "image/png" // Registers PNG decoder
)

// Dot import (imports into current namespace)
import (
	. "math"
)

// Multiple import blocks
import "io"
import "bytes"

func funcMain() {
	fmtAlias.Println("=== imports.go AST Node Coverage ===")
	fmtAlias.Println("Exercising AST Nodes:")

	// Standard library imports
	fmt.Println("  ✓ ast.ImportSpec (standard single import):")
	fmt.Println("    fmt package imported and used")

	// Grouped imports
	fmt.Println("  ✓ ast.GenDecl (grouped imports):")
	fmt.Printf("    os.Stdout type: %T\n", os.Stdout)
	fmt.Printf("    strings.ToUpper: %s\n", strings.ToUpper("test"))
	fmt.Printf("    time.Now: %s\n", time.Now().Format("2006-01-02"))

	// Aliased imports
	fmt.Println("  ✓ ast.ImportSpec (aliased import):")
	fmtAlias.Printf("    Using fmtAlias (alias for fmt)\n")
	result := str.Contains("hello", "ll")
	fmt.Printf("    Using str (alias for strings): Contains = %v\n", result)

	// Blank import (side effects only)
	fmt.Println("  ✓ ast.ImportSpec (blank import _):")
	fmt.Println("    image/png imported for side effects")
	fmt.Println("    (Registers PNG decoder without explicit use)")

	// Dot import
	fmt.Println("  ✓ ast.ImportSpec (dot import .):")
	// Can use math functions without qualifier
	piValue := Pi
	sqrtValue := Sqrt(16)
	fmt.Printf("    Pi (from math via .): %f\n", piValue)
	fmt.Printf("    Sqrt(16) (from math via .): %f\n", sqrtValue)

	// Multiple import blocks
	fmt.Println("  ✓ Multiple import blocks:")
	var buf bytes.Buffer
	buf.WriteString("io and bytes packages imported")
	fmt.Printf("    %s\n", buf.String())

	// Import usage patterns
	fmt.Println("  ✓ Import usage patterns:")

	// Package qualifier
	fmt.Println("    Package qualifier: fmt.Println")

	// Aliased package
	fmtAlias.Println("    Aliased package: fmtAlias.Println")

	// Imported types
	var t time.Time = time.Now()
	fmt.Printf("    Imported type: time.Time = %T\n", t)

	// Imported constants
	const layout = time.RFC3339
	fmt.Printf("    Imported constant: time.RFC3339 = %s\n", layout)

	// Imported functions
	upper := strings.ToUpper("hello")
	fmt.Printf("    Imported function: strings.ToUpper = %s\n", upper)

	// Nested package paths (standard library)
	fmt.Println("  ✓ Nested package imports:")
	fmt.Printf("    io package: io.Reader type = %T\n", (io.Reader)(nil))
	fmt.Println("    image/png package (nested: image/png)")

	// Import path patterns
	fmt.Println("  ✓ Import path patterns:")
	fmt.Println("    Single-level: fmt, os, io")
	fmt.Println("    Multi-level: strings, bytes")
	fmt.Println("    Nested: image/png")

	fmt.Println("Summary: Comprehensive import AST node coverage")
	fmt.Println("Primary AST Nodes: ast.GenDecl, ast.ImportSpec")
	fmt.Println("Features: single, grouped, aliased, blank, dot imports")
	fmt.Println("========================================")
}

func main() {
	funcMain()
}
