package main

import (
	"fmt"
)

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Implement the String() method for Person
// func (p Person) String() string {
// 	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
// }

func main() {
	p := Person{Name: "Alice", Age: 30}

	// fmt.Println calls p.String() automatically
	fmt.Println(p)

	// fmt.Sprintf also uses String()
	// str := fmt.Sprintf("Formatted: %s", p)
	// fmt.Println(str)
}
