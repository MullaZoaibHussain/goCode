package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map concept")

	languages := make(map[string]string) //make(map[key]value)
	languages["JS"] = "JavaScript"       //key = value
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("The list of languages is :", languages["JS"])

	//finding a value using key
	fmt.Println("Fulform of JS is", languages["JS"])
	//to delete some value using key
	delete(languages, "RB")
	fmt.Println("The list after deleting one value", languages)

	//using for loop
	for key, value := range languages {
		fmt.Printf("for the key %v, value is %v", key, value)
	}

}
