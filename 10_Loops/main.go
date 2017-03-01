package main

import "fmt"

func main(){

	fmt.Println("For 'condition' statement")

	var i = 0

	for i < 100{
		fmt.Println(i)
		i++
	}

	fmt.Println("For a normal for statement")

	for x:=100; x< 200; x++ {
		fmt.Println(x)
	}

	fmt.Println("For a 'infinate' loop")

	var z = 200

	for {
		fmt.Println(z)
		z++
		if z > 299 {
			break
		}
	}

	fmt.Println("Nested Loops")

	for i:= 0; i < 10; i++{
		for x:=0; x<10; x++ {
			fmt.Print((i * 10) + x, " ")
		}
		fmt.Print("\n")
	}



}
