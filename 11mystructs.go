package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	zoaib := User{"Mulla", "zoaib@gmail.com", true, 33}
	fmt.Println(zoaib)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
