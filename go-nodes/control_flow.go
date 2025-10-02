// Package main demonstrates control flow AST nodes.
// This file exercises control flow statement nodes.
package main

import "fmt"

// AST Nodes Covered:
// - ast.IfStmt
// - ast.SwitchStmt
// - ast.TypeSwitchStmt
// - ast.SelectStmt
// - ast.ForStmt
// - ast.RangeStmt
// - ast.CaseClause
// - ast.CommClause

func main() {
	fmt.Println("=== control_flow.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// IfStmt - if statements
	x := 10
	if x > 5 {
		fmt.Printf("  ✓ ast.IfStmt (simple): x=%d is greater than 5\n", x)
	}

	// If with initialization
	if y := 20; y > 15 {
		fmt.Printf("  ✓ ast.IfStmt (with init): y=%d is greater than 15\n", y)
	}

	// If-else
	if x < 5 {
		fmt.Printf("  This won't print\n")
	} else {
		fmt.Printf("  ✓ ast.IfStmt (else): x=%d is not less than 5\n", x)
	}

	// If-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Printf("  Grade: A\n")
	} else if score >= 80 {
		fmt.Printf("  ✓ ast.IfStmt (else if): Grade B for score=%d\n", score)
	} else if score >= 70 {
		fmt.Printf("  Grade: C\n")
	} else {
		fmt.Printf("  Grade: F\n")
	}

	// ForStmt - traditional for loop
	fmt.Printf("  ✓ ast.ForStmt (traditional):\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("    iteration %d\n", i)
	}

	// For loop with only condition (while-style)
	fmt.Printf("  ✓ ast.ForStmt (while-style):\n")
	count := 0
	for count < 3 {
		fmt.Printf("    count=%d\n", count)
		count++
	}

	// Infinite for loop with break
	fmt.Printf("  ✓ ast.ForStmt (infinite with break):\n")
	n := 0
	for {
		if n >= 2 {
			break
		}
		fmt.Printf("    n=%d\n", n)
		n++
	}

	// RangeStmt - range over slice
	fmt.Printf("  ✓ ast.RangeStmt (slice):\n")
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("    index=%d, value=%d\n", i, v)
	}

	// Range over map
	fmt.Printf("  ✓ ast.RangeStmt (map):\n")
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("    key=%s, value=%d\n", k, v)
	}

	// Range over string
	fmt.Printf("  ✓ ast.RangeStmt (string):\n")
	for i, r := range "Hi" {
		fmt.Printf("    index=%d, rune=%c\n", i, r)
	}

	// Range over channel
	fmt.Printf("  ✓ ast.RangeStmt (channel):\n")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for v := range ch {
		fmt.Printf("    received=%d\n", v)
	}

	// Range with blank identifier
	fmt.Printf("  ✓ ast.RangeStmt (blank identifier):\n")
	for _, v := range []int{5, 6} {
		fmt.Printf("    value=%d\n", v)
	}

	// SwitchStmt - expression switch
	fmt.Printf("  ✓ ast.SwitchStmt (expression):\n")
	day := 3
	switch day {
	case 1:
		fmt.Printf("    Monday\n")
	case 2:
		fmt.Printf("    Tuesday\n")
	case 3:
		fmt.Printf("    Wednesday (day=%d)\n", day)
	case 4, 5:
		fmt.Printf("    Thursday or Friday\n")
	default:
		fmt.Printf("    Weekend\n")
	}

	// Switch with initialization
	fmt.Printf("  ✓ ast.SwitchStmt (with init):\n")
	switch hour := 14; {
	case hour < 12:
		fmt.Printf("    Morning\n")
	case hour < 18:
		fmt.Printf("    Afternoon (hour=%d)\n", hour)
	default:
		fmt.Printf("    Evening\n")
	}

	// Switch without expression (tagless switch)
	fmt.Printf("  ✓ ast.SwitchStmt (tagless):\n")
	temp := 25
	switch {
	case temp < 0:
		fmt.Printf("    Freezing\n")
	case temp < 20:
		fmt.Printf("    Cold\n")
	case temp < 30:
		fmt.Printf("    Comfortable (temp=%d)\n", temp)
	default:
		fmt.Printf("    Hot\n")
	}

	// CaseClause with fallthrough
	fmt.Printf("  ✓ ast.CaseClause (with fallthrough):\n")
	num := 1
	switch num {
	case 1:
		fmt.Printf("    Case 1\n")
		fallthrough
	case 2:
		fmt.Printf("    Case 2 (via fallthrough or direct)\n")
	case 3:
		fmt.Printf("    Case 3\n")
	}

	// TypeSwitchStmt - type switch
	fmt.Printf("  ✓ ast.TypeSwitchStmt:\n")
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Printf("    Integer: %d\n", v)
	case string:
		fmt.Printf("    String: %q\n", v)
	case bool:
		fmt.Printf("    Boolean: %v\n", v)
	default:
		fmt.Printf("    Unknown type\n")
	}

	// Type switch without assignment
	fmt.Printf("  ✓ ast.TypeSwitchStmt (no assignment):\n")
	var j interface{} = 42
	switch j.(type) {
	case int:
		fmt.Printf("    It's an int\n")
	case string:
		fmt.Printf("    It's a string\n")
	default:
		fmt.Printf("    Unknown\n")
	}

	// SelectStmt - select for channels
	fmt.Printf("  ✓ ast.SelectStmt and ast.CommClause:\n")
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)
	ch1 <- 100
	ch2 <- "message"

	select {
	case val := <-ch1:
		fmt.Printf("    Received from ch1: %d\n", val)
	case msg := <-ch2:
		fmt.Printf("    Received from ch2: %s\n", msg)
	default:
		fmt.Printf("    No channel ready\n")
	}

	// Select with send
	ch3 := make(chan int, 1)
	select {
	case ch3 <- 99:
		fmt.Printf("    ✓ ast.CommClause (send): Sent to ch3\n")
	default:
		fmt.Printf("    Channel full\n")
	}

	// Select with timeout pattern
	timeout := make(chan bool, 1)
	timeout <- true
	select {
	case <-timeout:
		fmt.Printf("    ✓ ast.CommClause (receive): Timeout occurred\n")
	}

	// Nested control flow
	fmt.Printf("  ✓ Nested control flow:\n")
	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			switch i {
			case 0:
				fmt.Printf("    Nested: i=%d (even, zero)\n", i)
			}
		}
	}

	fmt.Println("Summary: All control flow AST node types exercised")
	fmt.Println("Primary AST Nodes: ast.IfStmt, ast.ForStmt, ast.RangeStmt, ast.SwitchStmt,")
	fmt.Println("                   ast.TypeSwitchStmt, ast.SelectStmt, ast.CaseClause, ast.CommClause")
	fmt.Println("========================================")
}
