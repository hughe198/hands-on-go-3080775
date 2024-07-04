// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val string = "Global"

func printGlobalVal() {
	fmt.Println(val)
}
func updateGlobalVal(a string) {
	val = a
	fmt.Println(val)
}
func main() {
	// create a local variable "val" and assign it an int value
	val := 100

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)
	// print the value of the package-level variable "val"
	printGlobalVal()
	// update the package-level variable "val" and print it again
	updateGlobalVal("New Global Value")
	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val =%v\n", &val, &val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	val = 100
	fmt.Printf("local val = %v\n", val)
}
