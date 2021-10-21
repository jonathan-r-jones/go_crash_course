package main

import "fmt"

func main(){
	// A map in Go is simply a key-value pair.
	emails:=make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
}