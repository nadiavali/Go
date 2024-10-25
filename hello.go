package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("helo")
	fmt.Println(time.Now().UTC().Date())
	fmt.Println(rand.Intn(3))
	fmt.Println(math.Pi)
	print()  // Only to realise how to call a function in go

}

func add(x int,y int)int{
	return x + y
}

func print() {
	fmt.Println(add(42, 78))
}