package main

import "fmt"

func main(){
	ids := []int{33, 76, 54, 23, 11, 2}
	for i, id:=range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	for _, id:=range ids {
		fmt.Printf("ID: %d\n", id)
	}
	sum:=0
	for _, id:= range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	emails2:=map[string]string{"tim": "tim@gmail.com", "Jerry": "Jerry@gmail.com"}

	for k, v := range emails2 {
		fmt.Printf("%s: %s\n", k, v)
	}
}