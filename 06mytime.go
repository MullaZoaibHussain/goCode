package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time topic")

	//	present := time.Now()//this will give time secounds but in diffrent formart
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.April, 11, 12, 20, 30, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
