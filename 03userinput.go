package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "welcome to user input"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin) // This line is used to get input from user
	fmt.Println("Enter rating for pizza:")
	//comma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating!", input)
	fmt.Printf("the type of rating is :%T", input)

}
