package main

import "fmt"

const (
	_ = iota
	KB = 1 << (iota * 10) //1 with a 10 (iota = 1) to the left bitshift
	MB = 1 << (iota * 10) //1 with a 20 (iota = 2) to the left bitshift
)

func main(){
	fmt.Println("\tbinary\t\t\tdecimal")
	fmt.Print("KB:\t")
	fmt.Printf("%b\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Print("MB:\t")
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}