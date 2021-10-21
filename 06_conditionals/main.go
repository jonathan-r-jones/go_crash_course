package main

import "fmt"

func main(){
	x := 5
	y := 10

	if x < y {
		fmt.Printf("If: %d is less than %d\n", x, y)
	} else {
		fmt.Printf("Else: %d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Printf("color is red\n")
	} else if color == "blue" {
		fmt.Printf("color is blue\n")
	} else {
		fmt.Printf("color is something else\n")
	}

	//switch
	switch color {
	case "red":
		fmt.Printf("color is red\n")
	case "blue":
		fmt.Printf("color is blue\n")
	default:
		fmt.Printf("color is something else\n")
	}

}