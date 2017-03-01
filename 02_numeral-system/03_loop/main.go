package main

import "fmt"

func main(){

	for i := 100000; i < 101000; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

}
