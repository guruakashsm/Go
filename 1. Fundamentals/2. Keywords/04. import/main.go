package main

import (
	"fmt"
	"math"
	ran "math/rand"
)

/*
=============================
IMPORT KEYWORD IN GO
=============================

--- 1. Definition of `import` ---
	`import` is a keyword in Go that is used to bring external packages or libraries into your Go program.
	It allows you to access functions, types, and variables defined in other Go packages.

--- 2. Purpose of `import` ---
	`import` is used to include reusable code from other packages into your current program.
	This enables you to use predefined libraries or third-party packages to extend the functionality of your program without having to write everything from scratch.

--- 3. Syntax ---
	The basic syntax for importing a package is:
	```go
	import "package-name"
	```
	You can also import multiple packages by grouping them:
	```go
	import (
		"package1"
		"package2"
	)
	```

--- 4. Standard Library and External Packages ---
	- Go provides a rich standard library, which includes packages for I/O, networking, text processing, etc.
	- You can also import external packages from repositories like GitHub or other sources.

--- 5. Aliasing Imported Packages ---
	You can alias an imported package using the `import` statement, which allows you to refer to it by a custom name:
	```go
	import customAlias "package-name"
	```

--- 6. Blank Identifier for Unused Imports ---
	If you need to import a package for its initialization side effects but do not use it directly in your code,
	you can use the blank identifier `_`:
	```go
	import _ "package-name"
	```

--- 7. Use Cases ---
	- Importing Go’s standard library (e.g., `fmt`, `os`, `net/http`) to use predefined functionality.
	- Importing third-party libraries for added functionality.
	- Organizing your code by splitting it into multiple packages and importing them into a main program.

=============================
EXAMPLES OF `IMPORT` IN GO
=============================

--- 8. Importing a Standard Package ---
	Using the `fmt` package to print text.

--- 9. Importing Multiple Packages ---
	Importing multiple packages in one statement.

--- 10. Aliasing Imported Packages ---
	Using an alias for an imported package.

*/

func main() {
	// Example 1: Importing a standard package
	fmt.Println("Hello, Go!")

	// Example 2: Importing multiple packages
	importExample()

	// Example 3: Aliasing imported packages
	// Here we import the "math/rand" package with an alias
	randExample()
}

func importExample() {
	// Importing "math" and using its functions
	// Here we are using the "math" package to calculate the square root of a number
	importedNumber := 16
	result := fmt.Sprintf("The square root of %d is %v", importedNumber, math.Sqrt(float64(importedNumber)))
	fmt.Println(result)
}

func randExample() {
	// Importing the "math/rand" package with an alias and using it for random number generation
	fmt.Printf("Random number example (using math/rand): %d \n", int(ran.Float32()*100))
}
