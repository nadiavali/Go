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
	fmt.Println(sqrt(4), sqrt(-4))

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
//this is not working out of func(non-declaration statement outside function body)
//man := Woman

func globalNonGlobalVars()(bool, bool, string){
	var i bool = to
	var b bool = Go
	bed := "bed"
	return i , b, bed
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}