package main

import (
	"fmt"
)

func main(){
	// Variable:
	// Decleration with type
	var foo int
	// Initialization of value
	foo = 1000
	// Short variable declaration
	// using 'walrus operator'
	bar := "Ten"
	fmt.Printf("%d %s",foo,bar)
}
