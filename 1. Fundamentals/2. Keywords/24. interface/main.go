package main

import "fmt"

/*
=============================
INTERFACE KEYWORD IN GO
=============================

--- 1. Definition of `interface` ---
	`interface` is a keyword in Go used to define a type that specifies a set of methods. 
	Any type that implements those methods satisfies the interface, without explicitly declaring its intent to do so.

--- 2. Purpose of `interface` ---
	Interfaces allow for defining behaviors (sets of methods) that can be shared among different types. 
	This provides a way to achieve polymorphism in Go, as any type that implements the methods of an interface can be treated as that interface type.

--- 3. Syntax ---
	To declare an interface, use the `interface` keyword followed by a set of method signatures:
	```go
	type MyInterface interface {
		Method1()
		Method2() string
	}
	```
	Any type that implements all methods of the interface implicitly satisfies the interface.

--- 4. Empty Interface ---
	The empty interface `interface{}` can hold values of any type. It’s commonly used to represent any value when you don’t know the specific type.
	```go
	var x interface{}
	x = 42
	x = "hello"
	```

--- 5. Implementing an Interface ---
	A type implements an interface by providing definitions for all the methods declared by the interface. There is no need for an explicit declaration of implementation.
	For example, the `Cat` type below implicitly implements the `Animal` interface:
	```go
	type Animal interface {
		Speak() string
	}

	type Cat struct{}

	func (c Cat) Speak() string {
		return "Meow"
	}
	```

--- 6. Type Assertion ---
	You can use type assertions to retrieve the underlying value from an interface.
	```go
	var i interface{} = 42
	x := i.(int) // Type assertion
	```

--- 7. Use Cases ---
	- Achieving polymorphism by allowing different types to be treated as the same interface.
	- Defining flexible APIs and ensuring that different types can be used interchangeably as long as they implement the necessary methods.
	- Using empty interfaces to accept any type of value.

=============================
EXAMPLES OF `INTERFACE` IN GO
=============================

--- 8. Defining and Implementing an Interface ---
	Example showing how a type implements an interface.

--- 9. Empty Interface ---
	Using the empty interface `interface{}` to hold any type.

--- 10. Type Assertion ---
	Example of using type assertion to retrieve the underlying value from an interface.

*/

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	// Example 1: Defining and Implementing an Interface
	var animal Animal
	animal = Dog{}
	fmt.Println("Dog says:", animal.Speak())

	animal = Cat{}
	fmt.Println("Cat says:", animal.Speak())

	// Example 2: Empty Interface
	var anything interface{}
	anything = 42
	fmt.Println("Anything is:", anything)
	anything = "Hello"
	fmt.Println("Anything is now:", anything)

	// Example 3: Type Assertion
	var i interface{} = 42
	if value, ok := i.(int); ok {
		fmt.Println("The value is an int:", value)
	} else {
		fmt.Println("The value is not an int.")
	}
}
