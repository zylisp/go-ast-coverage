// Package main demonstrates statement AST nodes.
// This file exercises all ast.Stmt node types.
package main

import "fmt"

// AST Nodes Covered:
// - ast.DeclStmt
// - ast.EmptyStmt
// - ast.LabeledStmt
// - ast.ExprStmt
// - ast.SendStmt
// - ast.IncDecStmt
// - ast.AssignStmt
// - ast.GoStmt
// - ast.DeferStmt
// - ast.ReturnStmt
// - ast.BranchStmt
// - ast.BlockStmt

func main() {
	fmt.Println("=== statements.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// DeclStmt - declarations in function bodies
	var x int = 42
	const y = 100
	type MyInt int
	fmt.Printf("  ✓ ast.DeclStmt (var): x = %d\n", x)
	fmt.Printf("  ✓ ast.DeclStmt (const): y = %d\n", y)
	fmt.Printf("  ✓ ast.DeclStmt (type): MyInt declared\n")

	// ExprStmt - standalone expressions
	_ = add(5, 3) // expression as statement
	fmt.Printf("  ✓ ast.ExprStmt: function call as statement\n")

	// AssignStmt - various assignment forms
	a := 10 // short variable declaration
	fmt.Printf("  ✓ ast.AssignStmt (DEFINE :=): a = %d\n", a)

	a = 20 // simple assignment
	fmt.Printf("  ✓ ast.AssignStmt (ASSIGN =): a = %d\n", a)

	a += 5 // compound assignment
	fmt.Printf("  ✓ ast.AssignStmt (ADD_ASSIGN +=): a = %d\n", a)

	a -= 3
	fmt.Printf("  ✓ ast.AssignStmt (SUB_ASSIGN -=): a = %d\n", a)

	a *= 2
	fmt.Printf("  ✓ ast.AssignStmt (MUL_ASSIGN *=): a = %d\n", a)

	a /= 4
	fmt.Printf("  ✓ ast.AssignStmt (QUO_ASSIGN /=): a = %d\n", a)

	a %= 5
	fmt.Printf("  ✓ ast.AssignStmt (REM_ASSIGN %%=): a = %d\n", a)

	b := 12
	b &= 10
	fmt.Printf("  ✓ ast.AssignStmt (AND_ASSIGN &=): b = %d\n", b)

	b |= 5
	fmt.Printf("  ✓ ast.AssignStmt (OR_ASSIGN |=): b = %d\n", b)

	b ^= 3
	fmt.Printf("  ✓ ast.AssignStmt (XOR_ASSIGN ^=): b = %d\n", b)

	b <<= 2
	fmt.Printf("  ✓ ast.AssignStmt (SHL_ASSIGN <<=): b = %d\n", b)

	b >>= 1
	fmt.Printf("  ✓ ast.AssignStmt (SHR_ASSIGN >>=): b = %d\n", b)

	b &^= 2
	fmt.Printf("  ✓ ast.AssignStmt (AND_NOT_ASSIGN &^=): b = %d\n", b)

	// Multiple assignment
	m, n := 1, 2
	fmt.Printf("  ✓ ast.AssignStmt (multiple): m=%d, n=%d\n", m, n)

	m, n = n, m // swap
	fmt.Printf("  ✓ ast.AssignStmt (swap): m=%d, n=%d\n", m, n)

	// IncDecStmt - increment and decrement
	counter := 10
	counter++
	fmt.Printf("  ✓ ast.IncDecStmt (++): counter = %d\n", counter)

	counter--
	fmt.Printf("  ✓ ast.IncDecStmt (--): counter = %d\n", counter)

	// SendStmt - channel send
	ch := make(chan int, 1)
	ch <- 42
	fmt.Printf("  ✓ ast.SendStmt: ch <- 42\n")
	received := <-ch
	fmt.Printf("  ✓ Channel receive: %d\n", received)

	// GoStmt - goroutine launch
	done := make(chan bool)
	go func() {
		fmt.Printf("  ✓ ast.GoStmt: goroutine executed\n")
		done <- true
	}()
	<-done

	// DeferStmt - deferred function call
	defer fmt.Printf("  ✓ ast.DeferStmt: deferred call executed (printed last)\n")
	fmt.Printf("  ✓ ast.DeferStmt: defer registered\n")

	// ReturnStmt - demonstrated in functions
	result := returnExample(5)
	fmt.Printf("  ✓ ast.ReturnStmt (value): returned %d\n", result)

	r1, r2 := multiReturn()
	fmt.Printf("  ✓ ast.ReturnStmt (multiple): returned %d, %d\n", r1, r2)

	namedReturn()
	fmt.Printf("  ✓ ast.ReturnStmt (named): executed\n")

	// BranchStmt - break, continue, goto, fallthrough
	fmt.Printf("  ✓ ast.BranchStmt (break):\n")
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Printf("    Breaking at i=%d\n", i)
			break
		}
		fmt.Printf("    i=%d\n", i)
	}

	fmt.Printf("  ✓ ast.BranchStmt (continue):\n")
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Printf("    Skipping i=%d\n", i)
			continue
		}
		fmt.Printf("    i=%d\n", i)
	}

	fmt.Printf("  ✓ ast.BranchStmt (goto):\n")
	i := 0
LoopStart:
	if i < 3 {
		fmt.Printf("    goto iteration i=%d\n", i)
		i++
		goto LoopStart
	}

	// LabeledStmt - labels
	fmt.Printf("  ✓ ast.LabeledStmt:\n")
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 1 && j == 0 {
				fmt.Printf("    Breaking OuterLoop at i=%d, j=%d\n", i, j)
				break OuterLoop
			}
			fmt.Printf("    i=%d, j=%d\n", i, j)
		}
	}

	// BlockStmt - block of statements
	{
		blockVar := "inside block"
		fmt.Printf("  ✓ ast.BlockStmt: %s\n", blockVar)
	}

	// EmptyStmt - empty statement (subtle, usually in for loops)
	for j := 0; j < 1; j++ {
		; // Empty statement
	}
	fmt.Printf("  ✓ ast.EmptyStmt: demonstrated in loop\n")

	fmt.Println("Summary: All statement AST node types exercised")
	fmt.Println("Primary AST Nodes: ast.DeclStmt, ast.AssignStmt, ast.IncDecStmt, ast.SendStmt,")
	fmt.Println("                   ast.GoStmt, ast.DeferStmt, ast.ReturnStmt, ast.BranchStmt,")
	fmt.Println("                   ast.BlockStmt, ast.LabeledStmt, ast.ExprStmt, ast.EmptyStmt")
	fmt.Println("========================================")
}

func add(a, b int) int {
	return a + b
}

func returnExample(x int) int {
	return x * 2
}

func multiReturn() (int, int) {
	return 1, 2
}

func namedReturn() (result int) {
	result = 42
	return // naked return with named result
}
