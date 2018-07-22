package main

import "fmt"

//struct to hold contact info
type contactInfo struct {
	email    string
	postCode string
}

//struct to hold person info and sub struct of contactInfo
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

	fmt.Printf("\n %+v \n %+v \n", sebastian, erin)

	toby.printPerson()

}

func (p person) printPerson() {

	fmt.Println()
	fmt.Println("fName : ", p.fName)
	fmt.Println("lName : ", p.lName)
	fmt.Println("email : ", p.info.email)
	fmt.Println("postCode : ", p.info.postCode+"\n")

}
