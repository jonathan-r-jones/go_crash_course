package main

import "fmt"

func main(){
	// Arrays
	var fruitArr [3]string

	fruitArr[0]="Apple"
	fruitArr[1]="Orange"
	fruitArr[2]="Banana"

	fmt.Println(fruitArr)
	
	vegetableArry := [4]string{"Carrot", "Pea", "Lettuce", "Tomato"}
	
	fmt.Println(vegetableArry)
	fmt.Println(len(vegetableArry))
	fmt.Println(vegetableArry[1:3])
}