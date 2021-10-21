package main

import "fmt"

func main(){
	// A map in Go is simply a key-value pair.
	emails:=make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete Bob's ass.
	delete(emails, "Bob")
	fmt.Println(emails)

	emails2:=map[string]string{"tim": "tim@gmail.com", "Jerry": "Jerry@gmail.com"}
	fmt.Println(emails2)

}