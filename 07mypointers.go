package main

import "fmt"

func main() {

	var p *int
	fmt.Printf("value if pointer %v\n", p) //default value is nil

	myNumber := 23

	var ptr = &myNumber                            //assinging the address of mynum to ptr using & symbol
	fmt.Println("The actual value of ", ptr)       //here we get the address because we assing ptr to myNumber using &
	fmt.Println("The actual value of ", *ptr)      // here * is used to show acctual value pointer of which we have assiged
	*ptr = *ptr + 2                                //here we are manuplilating the value of myNumber using its refrens with acttual value by * symbol
	fmt.Println("New value of myNumber", myNumber) // Printing new value

}
