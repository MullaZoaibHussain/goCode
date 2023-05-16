package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitlist = []string{"Banana", "peach", "apple"}
	fmt.Printf("The type of fruitlist is %T\n ", fruitlist)
	fmt.Println("fruitlist before appeding is :", fruitlist)

	fruitlist = append(fruitlist, "mango", "Tomato")
	fmt.Println("fruitlist after appeding is :", fruitlist)
	fruitlist = append(fruitlist[1:]) // here in square brakke is used for range
	fmt.Println("fruitlist after appeding [1:] is :", fruitlist)
	fruitlist = append(fruitlist[:3])                             //ignore
	fmt.Println("fruitlist after appeding [:3] is :", fruitlist)  //ignore
	fruitlist = append(fruitlist[1:3])                            //ignore
	fmt.Println("fruitlist after appeding [1:3] is :", fruitlist) //ignore

	highscores := make([]int, 4)
	highscores[0] = 231
	highscores[1] = 934
	highscores[2] = 433
	highscores[3] = 122
	fmt.Println("The highscore list is:", highscores)
	highscores = append(highscores, 777, 881, 910)
	fmt.Println("The highscore list after appending is:", highscores)
	sort.Ints(highscores) //sorting happens in accending
	//.Println("The highscore list after SORTING is:", highscores)
	//fmt.Println(sort.IntsAreSorted(highscores)) //if it is sorted it will give me true else false

	//how to remove a value from slice using index
	var course = []string{"Java", "reactJs", "Javascript", "Ruby"}
	fmt.Println(course)
	index := 2
	course = append(course[:index], course[index+1:]...) //here we are removing 2nd value using index variable
	fmt.Println(course)

}
