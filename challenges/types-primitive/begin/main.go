// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func printGlobal() {
	fmt.Printf("%T, global val = %v\n", val, val)
}

func updateGlobal() {
	val = "updated global"
}

func main() {
	// create a local variable "val" and assign it an int value
	val := 45

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printGlobal()

	// update the package-level variable "val" and print it again
	updateGlobal()
	printGlobal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 200
	fmt.Printf("%T, local val = %v\n", val, val)
}
