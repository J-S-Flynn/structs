package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}
type person struct {
	fName string
	lName string
	info  contactInfo
}

func main() {

	//three ways of declaring and initalising struct

	//1.
	sebastian := person{"Sebastian", "Flynn", contactInfo{"seb@someplace.com", "ls9 2ho"}}

	//2.
	erin := person{fName: "Erin", lName: "Flynn", info: contactInfo{email: "erin@somplace.com", postCode: "ls9 2ho"}}

	//3.
	var toby person

	toby.fName, toby.lName, toby.info.email, toby.info.postCode = "Toby", "Flynn", "toby@someplace.com", "ls9 2ho"

	//same print out for all

	fmt.Println("\n", sebastian, "\n", erin, "\n", toby)
}
