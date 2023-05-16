package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	println("Welcome to math in golang")

	// var mynumberOne int = 5
	// var mynumberTwo float64 = 4.10370470
	// fmt.Println("The addition of two number is ", mynumberOne+int(mynumberTwo)) // by using inbuild conversion we can loose some value after point

	//random number
	//rand.Seed(5)//if we use this will only one random number rather use ------rand.Seed(time.Now().unixNano())

	// // rand.Seed(time.Now().UnixNano())
	// // fmt.Println(rand.Intn(5))

	//random from crypto func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
	RandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(RandomNum)

}
