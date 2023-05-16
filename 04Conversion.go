package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza App")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Pizza app:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", input)
	// conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}

}
