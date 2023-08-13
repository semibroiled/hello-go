package main //Declare a main package, made up of all files in same directory as a way to group function

import (
	"fmt" //import fmt package, with functionsto format text as part of stl
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	//Set properties of logger
	log.SetPrefix("greetings: ") //Sets prefix
	log.SetFlags(0) //Disables printing time, source file and line number
	// Implement main function that runs on default when maain package runs
	fmt.Println("Hello, World!")

	//call new import fcn
	fmt.Println(quote.Opt())

	// A slice of names.

	names := []string{"Fatma", "Hobbi", "Ba7ebbek", "Basbousa", "Hayati"}


	//Handle error and call on hellos
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	//handle error and call on hello
	message, err := greetings.Hello("Hobbi")
	if err != nil {
		log.Fatal(err)
	}
	//Print from local import if no error
	fmt.Println(message)
	fmt.Println(messages)
}