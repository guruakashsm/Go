package main

import "fmt"

/*
 Identifiers are the user-defined names of the program that are used to identify the variables, functions, classes, modules, etc.
 An identifier can be composed of letters, digits, and the underscore character (_).
 An identifier must start with a letter or an underscore (_), and it cannot start with a digit.
 Identifiers are case-sensitive, i.e., the variable name and the variable Name are different.
 There are some predeclared identifiers available for constants, types, and functions.

 For Constants:
   true, false, iota, nil

 For Types:
	int, int8, int16, int32, int64, uint,
	uint8, uint16, uint32, uint64, uintptr,
	float32, float64, complex128, complex64,
	bool, byte, rune, string, error

 For Functions:
	make, len, cap, new, append, copy, close,
	delete, complex, real, imag, panic, recover

 There are some reserved words in Go that cannot be used as identifiers. These reserved words are called keywords.

 The following are some examples of valid and invalid identifiers:
 Valid identifiers:
	1. var_name
	2. _var
	3. var1
	4. var_1
	5. VAR
Invalid identifiers:
	1. 1var
	2. var-name
	3. var name
	4. var@name
	5. var#name

The identifier represented by the underscore character(_) is known as a blank identifier. It is used to store the values that are not used in the program.

Exported Identifiers:
   An identifier is exported if it starts with an uppercase letter. An exported identifier can be accessed from another package.

*/

// Constants
const (
	trueValue = true // valid constant (not a reserved keyword)
	pi        = 3.14159
)

// Types
type customInt int
type person struct {
	name string
	age  int
}

// Exported Identifiers (start with uppercase letter)
var ExportedVar = "I am exported" // This is accessible from other packages

// Blank Identifier (_)
func useBlankIdentifier() {
	_, err := fmt.Println("This function uses the blank identifier.")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	// Valid Identifiers
	var myVar = "Hello, Go!"
	var myVar1 = 42
	var _var = 3.14
	var VarName = "Exported variable"

	// Invalid Identifiers (uncomment to see compilation errors)
	// var 1var = "invalid"  // Invalid: starts with a digit
	// var var-name = "invalid"  // Invalid: contains hyphen
	// var var name = "invalid"  // Invalid: contains space
	// var var@name = "invalid"  // Invalid: contains special character '@'

	// Use of Blank Identifier
	useBlankIdentifier()

	// Print Valid Identifiers
	fmt.Println(myVar)   // prints: Hello, Go!
	fmt.Println(myVar1)  // prints: 42
	fmt.Println(_var)    // prints: 3.14
	fmt.Println(VarName) // prints: Exported variable

	// Accessing Exported Identifier (ExportedVar)
	fmt.Println(ExportedVar) // prints: I am exported

	// Constants usage
	fmt.Println("Pi is:", pi)             // prints: Pi is: 3.14159
	fmt.Println("True value:", trueValue) // prints: True value: true

	// Using custom types
	var num customInt = 10
	var p person
	p.name = "John Doe"
	p.age = 30

	fmt.Println("Custom int:", num)       // prints: Custom int: 10
	fmt.Println("Person:", p.name, p.age) // prints: Person: John Doe 30
}
