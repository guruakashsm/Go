package main

/*
=============================
PACKAGE IN GO
=============================

--- 1. Definition of `package` ---
	The `package` keyword in Go is used to define a group of related Go source files. All Go source files belong to a package, and the `package` declaration must appear at the top of each Go source file. The first line of a Go file typically declares the package it belongs to.

--- 2. Purpose of `package` ---
	Go packages are used to organize code and group related functions, types, variables, and constants together. By organizing code into packages, you can reuse it across multiple projects and maintain a clean, modular structure.

--- 3. The `main` Package ---
	When building executable programs, you use the `main` package. It defines the entry point of the program with the `main()` function. All Go programs must contain a `main` package if they are to be run.

--- 4. Importing Packages ---
	To use a package in your program, you need to import it using the `import` keyword. This allows you to access the functions and types defined within that package.

--- 5. Standard Library and Custom Packages ---
	Go provides a rich set of standard libraries (like `fmt`, `math`, `net/http`), but you can also create your own custom packages. Custom packages must reside in directories that match their names and should be imported by their path.

--- 6. Package Name Convention ---
	The name of the package should reflect its purpose. When importing a package, you can use a shorthand alias for the package name. By convention, Go uses lowercase names for packages.

=============================
EXAMPLES OF `PACKAGE` IN GO
=============================

--- 7. Declaring a Package ---
	The `package` declaration defines the package to which the Go file belongs.

--- 8. Importing and Using a Package ---
	You can import standard or custom packages using the `import` statement and then use the functions, types, or variables declared in the imported package.

*/

import (
	"fmt"
	"log"
	"package/modle"
)

func main() {
	// Using the fmt package to print a message to the console
	fmt.Println("Hello, Go!")

	// Importing a custom package needs mod file

	req := modle.NewRequest("john", "doe")
	log.Println("Request created successfully ", req)
	log.Printf("Username : %v ", req.GetUsername())
	log.Printf("Password : %v ", req.GetPassword())

	res := modle.NewResponse("success", "200")
	log.Println("Response created successfully", res)

	res.SetMessage("failed")

	log.Printf("Response : %v ", res)
}
