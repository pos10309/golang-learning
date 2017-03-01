package main

import "fmt"

var q string

func main(){

	a:= 10
	b:= "golang"
	c:= 4.17
	d:= true

	var  e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n",e,e)
	fmt.Printf("%v %T \n",f,f)
	fmt.Printf("%v %T \n",g,g)
	fmt.Printf("%v %T \n",h,h)

	var message string
	message = "Hello World"

	fmt.Println(message)

	var i,j,k int
	a = 1

	fmt.Println(i,j,k)

	var l,m,n = 3,4,5

	fmt.Println(l,m,n)

	fmt.Println(q)
	assignQ()
	fmt.Println(q)


}

func assignQ(){
	q = "Function Run"
}