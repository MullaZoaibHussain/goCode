package main

import "fmt"

func main() {
	fmt.Println("Welcome to my array")
	var fruitlist = [4]string{"PineApple", "Banana", "peach"}
	fmt.Println("Fruit list is :", fruitlist)
	fmt.Println("Fruit list is :", len(fruitlist))

	var veglist [5]string
	veglist[0] = "potato"
	veglist[1] = "pea" //index 2 has no value so that in output you will get a space after index one
	veglist[3] = "mushroom"
	veglist[4] = "beans"
	fmt.Println("The veglist is :", veglist)
	fmt.Printf("The type veglist is %T\n", veglist)
	fmt.Println("The veglist is :", len(veglist))

}
