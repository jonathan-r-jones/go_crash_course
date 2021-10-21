package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(param1 int, param2 int) int {
	return param1 + param2
}

func main(){
	fmt.Println(greeting("Bill"))
	fmt.Println(getSum(5, 25))
}