package main

import "fmt"

func main(){

	x:= 13/3
	y:= 13 % 3
	fmt.Println("Division value: ", x)
	fmt.Println("Remainder: ", y)

	for (i:= 0; i< 100; i++) {
		if i % 2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}

}