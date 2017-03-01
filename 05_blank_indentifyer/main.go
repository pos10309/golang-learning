package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)




func main(){
	//the http.Get function return two variables, a result and an error function
	//since we will not use the error function in this example, we need to put in a
	// "_", this is a blank identifier
	res, _ := http.Get("http://www.cnn.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

}
