package main

import "fmt"

//jwtToken := 3000000//here it will error because no var operator is used inside method only.It can be written
//as var jwtToken int =300000 or var jwtToken =300000
const LoginToken string = "gfgfwfbbf" // any variable ot constant starts with cap letter it is " PUBLIC"

func main() {
	var username string = "Zoaib"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type : %T \n", isLoggedin)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type : %T \n", smallval)

	var smallfloat float32 = 255.01802378 //if we use float32 you'll get only minimum values after (.)
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type : %T \n", smallfloat)

	var smallfloat1 float64 = 255.01802378
	fmt.Println(smallfloat1)
	fmt.Printf("variable is of type : %T \n", smallfloat1)

	//Default values and alaises

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n", anotherVariable)

	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("variable is of type : %T \n", anotherVariable1)

	//implicit type
	var name = "zoaib" // here lexer will take care of type
	fmt.Println(name)
	//name=2//here it will give an error because it has already decided it string to int can;t be assign

	//no var style [:=] the symbol is used to when var is not used
	//walrus operator, it will give error if we use it out side method
	NumberOfUser := 300000
	fmt.Println(NumberOfUser)
	//public access
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
