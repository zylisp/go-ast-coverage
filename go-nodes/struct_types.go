// Package main demonstrates struct type AST nodes in detail.
// This file exercises ast.StructType and ast.Field nodes.
package main

import "fmt"

// AST Nodes Covered:
// - ast.StructType
// - ast.Field
// - ast.FieldList
// - ast.Tag

// Simple struct
type Person struct {
	Name string
	Age  int
}

// Struct with tags (ast.Tag)
type JSONPerson struct {
	Name  string `json:"name" xml:"person_name"`
	Age   int    `json:"age" xml:"person_age"`
	Email string `json:"email,omitempty" validate:"email"`
}

// Struct with embedded fields
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person          // embedded struct
	Address         // embedded struct
	EmployeeID int
	Department string
}

// Struct with anonymous fields
type Mixed struct {
	int         // anonymous field
	string      // anonymous field
	Named int
}

// Struct with pointer fields
type Node struct {
	Value int
	Next  *Node // self-referential
	Prev  *Node
}

// Empty struct
type Empty struct{}

// Struct with various field types
type Complex struct {
	// Basic types
	IntField    int
	StringField string
	BoolField   bool

	// Pointer
	PtrField *int

	// Array and slice
	ArrayField [3]int
	SliceField []string

	// Map
	MapField map[string]int

	// Channel
	ChanField chan int

	// Function
	FuncField func(int) string

	// Interface
	IfaceField interface{}

	// Nested struct
	NestedField struct {
		X int
		Y int
	}
}

// Struct with unexported fields
type Encapsulated struct {
	Public  string
	private int
}

func funcMain() {
	fmt.Println("=== struct_types.go AST Node Coverage ===")
	fmt.Println("Exercising AST Nodes:")

	// Simple struct
	fmt.Println("  ✓ ast.StructType (simple):")
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("    Person: %+v\n", p)

	// Struct with tags
	fmt.Println("  ✓ ast.StructType with ast.Tag:")
	jp := JSONPerson{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
	}
	fmt.Printf("    JSONPerson: %+v\n", jp)

	// Embedded structs
	fmt.Println("  ✓ ast.StructType (embedded fields):")
	e := Employee{
		Person: Person{
			Name: "Charlie",
			Age:  35,
		},
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			Zip:    "02101",
		},
		EmployeeID: 1001,
		Department: "Engineering",
	}
	fmt.Printf("    Employee.Name (via Person): %s\n", e.Name)
	fmt.Printf("    Employee.City (via Address): %s\n", e.City)
	fmt.Printf("    Employee.Department: %s\n", e.Department)

	// Anonymous fields
	fmt.Println("  ✓ ast.StructType (anonymous fields):")
	m := Mixed{
		int:    42,
		string: "hello",
		Named:  100,
	}
	fmt.Printf("    Mixed.int: %d\n", m.int)
	fmt.Printf("    Mixed.string: %s\n", m.string)
	fmt.Printf("    Mixed.Named: %d\n", m.Named)

	// Self-referential struct (linked list)
	fmt.Println("  ✓ ast.StructType (self-referential):")
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	fmt.Printf("    Linked list: %d -> %d -> %d\n", n1.Value, n1.Next.Value, n1.Next.Next.Value)

	// Empty struct
	fmt.Println("  ✓ ast.StructType (empty):")
	var empty Empty
	_ = empty
	fmt.Printf("    Empty struct: %+v\n", empty)
	fmt.Printf("    Empty struct size: %d bytes\n", 0) // Empty structs are 0 bytes

	// Complex struct with all field types
	fmt.Println("  ✓ ast.StructType (complex with ast.Field variations):")
	num := 42
	c := Complex{
		IntField:    10,
		StringField: "test",
		BoolField:   true,
		PtrField:    &num,
		ArrayField:  [3]int{1, 2, 3},
		SliceField:  []string{"a", "b", "c"},
		MapField:    map[string]int{"one": 1},
		ChanField:   make(chan int, 1),
		FuncField:   func(i int) string { return fmt.Sprintf("%d", i) },
		IfaceField:  "interface value",
		NestedField: struct{ X, Y int }{X: 5, Y: 10},
	}
	fmt.Printf("    Complex.IntField: %d\n", c.IntField)
	fmt.Printf("    Complex.PtrField: %d\n", *c.PtrField)
	fmt.Printf("    Complex.ArrayField: %v\n", c.ArrayField)
	fmt.Printf("    Complex.SliceField: %v\n", c.SliceField)
	fmt.Printf("    Complex.MapField: %v\n", c.MapField)
	fmt.Printf("    Complex.FuncField(99): %s\n", c.FuncField(99))
	fmt.Printf("    Complex.NestedField: %+v\n", c.NestedField)

	// Anonymous struct literal
	fmt.Println("  ✓ ast.StructType (anonymous):")
	anon := struct {
		X int
		Y string
	}{
		X: 100,
		Y: "anonymous",
	}
	fmt.Printf("    Anonymous struct: %+v\n", anon)

	// Struct field access
	fmt.Println("  ✓ ast.Field access patterns:")
	fmt.Printf("    Direct access: p.Name = %s\n", p.Name)
	fmt.Printf("    Pointer access: (&p).Age = %d\n", (&p).Age)
	pPtr := &p
	fmt.Printf("    Via pointer: pPtr.Name = %s\n", pPtr.Name)

	// Struct with methods (methods defined below)
	fmt.Println("  ✓ ast.StructType with methods:")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("    Rectangle area: %d\n", rect.Area())
	fmt.Printf("    Rectangle perimeter: %d\n", rect.Perimeter())

	// Struct comparison
	fmt.Println("  ✓ Struct operations:")
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 25}
	fmt.Printf("    p1 == p2: %v\n", p1 == p2)
	fmt.Printf("    p1 == p3: %v\n", p1 == p3)

	fmt.Println("Summary: Comprehensive struct type AST node coverage")
	fmt.Println("Primary AST Nodes: ast.StructType, ast.Field, ast.FieldList, ast.Tag")
	fmt.Println("========================================")
}

// Struct for method demonstration
type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	funcMain()
}
