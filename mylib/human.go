package mylib

import "fmt"

var Public string = "pub"
var private string = "pri"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
