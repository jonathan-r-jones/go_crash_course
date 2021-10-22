package main

import ("fmt"
				"strconv")


type Person struct{
	firstName, lastName, city, gender string
	age int
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m"{
		return
	} else {
		person.lastName = spouseLastName
	}
}

func (person Person) greet() string{
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

func main(){
	person1:=Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	fmt.Println(person1)
	person2:=Person{"Samantha", "Smith", "Boston", "f", 26}
	fmt.Println(person2)
	fmt.Println(person2.age)
	person2.age++
	fmt.Println(person2.age)
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

}