package main

import "fmt"

func passingValue( x int){
	x = 0
	fmt.Println("The value of x in passingValue is set to :", x)
	fmt.Println("x has the address: ", &x)

}

func passingRef(x *int){
	*x = 0
	fmt.Println("The value of *x in passingRef is set to :", *x)
	fmt.Println("*x has the value: ", x)
}

func main(){

	a:= 43

	fmt.Println("The value of variable a is: ", a)
	fmt.Println("at the address: ", &a)

	var b *int = &a
	fmt.Println("The value of variable b is a pointer: ", b)

	fmt.Print("The type is shown as: ")
	fmt.Printf("%T \n", b)

	fmt.Print("We can use the * to dereference the pointer and get the value: ")
	fmt.Println(*b)

	fmt.Print("Now using pointers to change the value *b and then print value for varible a: ")
	*b = 42
	fmt.Println(a)

	fmt.Println("We are now passing a into the function passingValue with x int as a paramater with the address", &a)
	passingValue(a)
	fmt.Println("Back in main the value for a is still: ", a, " with the address still ", &a)

	fmt.Println("We are now passing &a into the function passingRef with x *int as a paramater")
	passingRef(&a)
	fmt.Println("Back in main the value for a has changed to: ", a)



}