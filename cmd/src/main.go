package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {

	//three ways of declaring and initalising struct

	//1.
	sebastian := person{"Sebastian", "Flynn"}

	//2.
	erin := person{fName: "Erin", lName: "Flynn"}

	//3.
	var toby person

	toby.fName, toby.lName = "Toby", "Flynn"

	//same print out for all

	fmt.Println(sebastian, erin, toby)
}
