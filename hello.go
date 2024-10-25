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
	fmt.Println(swap("World", "hello"))
	fmt.Println(spliting(20))
	fmt.Println(globalNonGlobalVars())

}

func add(x int,y int)int{
	return x + y
}

func print() {
	fmt.Println(add(42, 78))
}

func swap(a, b string)(string, string){
	return b, a
}


// return values and types schuld be spcified here
func spliting(number int)(right_side, left_side int){
	right_side = number * 4 / 9
	left_side= number - right_side
	return // naked return(no argas)

} 


var Go, to, bed bool = true, false, true

func globalNonGlobalVars()(bool, bool){
	var i bool = to
	var b bool = Go
	return i , b
}

