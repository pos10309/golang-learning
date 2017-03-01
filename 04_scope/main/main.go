package main

import (
	"github.com/williammarkpainter/golang-learning/04_scope/vis"
	"fmt"
)

func wrapper() func() int {
	x:=100
	return func() int{
		x++
		return x
	}
}

func main(){
	println(vis.MyNameVissible)
	vis.PrintMeVisible()

	x := 0

	increment := func() int{
		x++
		return x
	}

	fmt.Println(increment())


	tryingThis := wrapper()
	fmt.Println(tryingThis())
	fmt.Println(tryingThis())


}

