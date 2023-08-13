package main //Declare a main package, made up of all files in same directory as a way to group function 

import "fmt" //import fmt package, with functionsto format text as part of stl

import "rsc.io/quote"

func main() {
	// Implement main function that runs on default when maain package runs
	fmt.Println("Hello, World!")

	//call new import fcn
	fmt.Println(quote.Opt())
}