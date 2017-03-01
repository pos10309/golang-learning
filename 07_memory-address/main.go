package main

import "fmt"

func main(){

	m := "Hello world"
	n := &m

	a:= 43

	fmt.Println(m,n)
	fmt.Println("The Number " , a, " has address ", &a)

	fmt.Print("In Decimal this address is: ")
	fmt.Printf("%d \n", &a)

}
