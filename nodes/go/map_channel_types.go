// Package main demonstrates map and channel AST nodes.
// This file exercises ast.MapType and ast.ChanType nodes.
package main

import (
	"fmt"
	"time"
)

// AST Nodes Covered:
// - ast.MapType
// - ast.ChanType (bidirectional, send-only, receive-only)
// - ast.SendStmt
// - ast.UnaryExpr (channel receive <-)

func funcMain() {
	fmt.Println("=== map_channel_types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Map declarations
	fmt.Println("  ✓ ast.MapType (various key/value types):")

	// Nil map
	var nilMap map[string]int
	fmt.Printf("    Nil map: %v (nil: %v)\n", nilMap, nilMap == nil)

	// Empty map with make
	emptyMap := make(map[string]int)
	fmt.Printf("    Empty map: %v (nil: %v, len: %d)\n", emptyMap, emptyMap == nil, len(emptyMap))

	// Map literal
	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Printf("    map[string]int: %v\n", map1)

	// Map with different types
	map2 := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}
	fmt.Printf("    map[int]string: %v\n", map2)

	map3 := map[bool]float64{
		true:  1.1,
		false: 2.2,
	}
	fmt.Printf("    map[bool]float64: %v\n", map3)

	// Map indexing and assignment
	fmt.Println("  ✓ Map operations (ast.IndexExpr):")
	fmt.Printf("    map1[\"two\"]: %d\n", map1["two"])
	map1["four"] = 4
	fmt.Printf("    After map1[\"four\"]=4: %v\n", map1)

	// Map lookup with ok idiom
	val, ok := map1["two"]
	fmt.Printf("    map1[\"two\"]: value=%d, ok=%v\n", val, ok)
	val2, ok2 := map1["nonexistent"]
	fmt.Printf("    map1[\"nonexistent\"]: value=%d, ok=%v\n", val2, ok2)

	// Delete from map
	delete(map1, "two")
	fmt.Printf("    After delete(map1, \"two\"): %v\n", map1)

	// Map iteration
	fmt.Println("  ✓ Map iteration:")
	for k, v := range map1 {
		fmt.Printf("    map1[%s]=%d\n", k, v)
	}

	// Nested maps
	fmt.Println("  ✓ Nested maps (ast.MapType):")
	nested := map[string]map[string]int{
		"outer1": {"inner1": 1, "inner2": 2},
		"outer2": {"inner3": 3, "inner4": 4},
	}
	fmt.Printf("    Nested map: %v\n", nested)
	fmt.Printf("    nested[\"outer1\"][\"inner2\"]: %d\n", nested["outer1"]["inner2"])

	// Map with struct values
	type Person struct {
		Name string
		Age  int
	}
	mapStruct := map[string]Person{
		"alice": {Name: "Alice", Age: 30},
		"bob":   {Name: "Bob", Age: 25},
	}
	fmt.Printf("    map[string]Person: %v\n", mapStruct)

	// Map with slice values
	mapSlice := map[string][]int{
		"primes": {2, 3, 5, 7, 11},
		"evens":  {2, 4, 6, 8, 10},
	}
	fmt.Printf("    map[string][]int: %v\n", mapSlice)

	// Map capacity hint
	largeMap := make(map[int]int, 100)
	largeMap[1] = 1
	fmt.Printf("    Map with capacity hint: len=%d\n", len(largeMap))

	// Channel types
	fmt.Println("  ✓ ast.ChanType (bidirectional):")
	ch1 := make(chan int)
	ch2 := make(chan string, 5) // buffered
	fmt.Printf("    Unbuffered chan int: %T\n", ch1)
	fmt.Printf("    Buffered chan string: %T (cap: %d)\n", ch2, cap(ch2))

	// Buffered channel operations
	ch2 <- "first"
	ch2 <- "second"
	fmt.Printf("    After sends: len=%d, cap=%d\n", len(ch2), cap(ch2))
	msg := <-ch2
	fmt.Printf("    Received: %s, remaining len=%d\n", msg, len(ch2))

	// Send-only channel (ast.ChanType with Dir=SEND)
	fmt.Println("  ✓ ast.ChanType (send-only chan<-):")
	chSend := make(chan int, 1)
	sendOnlyExample(chSend)
	received := <-chSend
	fmt.Printf("    Received from send-only channel: %d\n", received)

	// Receive-only channel (ast.ChanType with Dir=RECV)
	fmt.Println("  ✓ ast.ChanType (receive-only <-chan):")
	chRecv := make(chan int, 1)
	chRecv <- 42
	receiveOnlyExample(chRecv)

	// Channel close
	fmt.Println("  ✓ Channel close and range:")
	chClose := make(chan int, 3)
	chClose <- 1
	chClose <- 2
	chClose <- 3
	close(chClose)
	for val := range chClose {
		fmt.Printf("    Received: %d\n", val)
	}

	// Channel with select
	fmt.Println("  ✓ Channel with select (ast.SelectStmt):")
	ch3 := make(chan int, 1)
	ch4 := make(chan string, 1)
	ch3 <- 100
	ch4 <- "message"

	select {
	case v := <-ch3:
		fmt.Printf("    Received from ch3: %d\n", v)
	case m := <-ch4:
		fmt.Printf("    Received from ch4: %s\n", m)
	default:
		fmt.Println("    No channel ready")
	}

	// Goroutine with channel
	fmt.Println("  ✓ Goroutine with channel (ast.GoStmt):")
	done := make(chan bool)
	go func() {
		fmt.Println("    Goroutine executing")
		done <- true
	}()
	<-done
	fmt.Println("    Goroutine completed")

	// Channel timeout pattern
	fmt.Println("  ✓ Channel timeout pattern:")
	timeout := time.After(1 * time.Millisecond)
	select {
	case <-timeout:
		fmt.Println("    Timeout occurred")
	}

	// Nil channel behavior
	fmt.Println("  ✓ Nil channel:")
	var nilChan chan int
	fmt.Printf("    Nil channel: %v (nil: %v)\n", nilChan, nilChan == nil)

	// Channel of channels
	fmt.Println("  ✓ Channel of channels:")
	chOfCh := make(chan chan int, 1)
	innerCh := make(chan int, 1)
	innerCh <- 99
	chOfCh <- innerCh
	receivedCh := <-chOfCh
	receivedVal := <-receivedCh
	fmt.Printf("    Received from channel of channels: %d\n", receivedVal)

	// Directional channel conversion
	fmt.Println("  ✓ Channel direction conversion:")
	biCh := make(chan int, 1)
	var sendCh chan<- int = biCh  // convert to send-only
	var recvCh <-chan int = biCh  // convert to receive-only
	sendCh <- 77
	val77 := <-recvCh
	fmt.Printf("    Directional channels work: %d\n", val77)

	fmt.Println("Summary: Comprehensive map and channel AST node coverage")
	fmt.Println("Primary AST Nodes: ast.MapType, ast.ChanType, ast.SendStmt")
	fmt.Println("Features: maps, channels (buffered/unbuffered, directional), select, goroutines")
	fmt.Println("========================================")
}

// Helper function demonstrating send-only channel parameter
func sendOnlyExample(ch chan<- int) {
	ch <- 999
}

// Helper function demonstrating receive-only channel parameter
func receiveOnlyExample(ch <-chan int) {
	val := <-ch
	fmt.Printf("    Received in receive-only function: %d\n", val)
}

func main() {
	funcMain()
}
