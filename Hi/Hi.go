package main

import (
	"fmt"
	"time"
)

func main(){
	week()
	Greet()
}

func week() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
		fmt.Println(time.Now().Weekday() + 0)
	}
}

func Greet(){
	t := time.Now()
	var name string
	println("What is your name?")
	fmt.Scanln(&name)
	switch{
	case t.Hour() < 12:
		fmt.Println("Good Morning", name)
	case t.Hour() < 17:
		fmt.Println("Good Afternoon", name)
	default:
		fmt.Println("Good Night", name)
	}
}