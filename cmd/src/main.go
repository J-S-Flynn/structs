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
	//printf allows us to format the string, %+v will not only the value of the variables, but also the name of them
	fmt.Printf("\n %+v \n %+v \n %+v \n", sebastian, erin, toby)

	//imporant opiratos & and *. the oparator & tels the compiler to give access to the
	//address of the variable in memory so in this case seb is the address of the variable sebastian
	seb := &sebastian
	seb.updatName("seb")

	sebastian.printPerson()

	//that a little cumbersom though to go has the option to call a reciver with a pionte to type by variable
	//of that type and assumes pionter. this will work with a var of type person, but would with a non maching type
	toby.updatName("Tobob")

	toby.printPerson()

	//and for the sake of completion
	erin.printPerson()
}

//it should be noted the is a difrence between *person and (*p).fName. *person denotes we are dealing with a pionter to a type
//in this case its a pionter of type person so the reciver is expecting a pionter to type person
func (p *person) updatName(newName string) {

	//the *p tells the compiler we want to look at the value of the variable so in this case we want to change the
	//value of the variable from the old name to the new one.
	(*p).fName = newName
}

//simple reciver to allow for printing of structs of type person
func (p person) printPerson() {

	fmt.Println()
	fmt.Println("fName : ", p.fName)
	fmt.Println("lName : ", p.lName)
	fmt.Println("email : ", p.info.email)
	fmt.Println("postCode : ", p.info.postCode+"\n")

}
