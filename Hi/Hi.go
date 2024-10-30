package main

import (
	"fmt"
	"time"
)

func main(){
	week()
	Greet()
	DearPointer()
	CallCollection()
	inttpe()
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


/* Python uses references, not pointers in Go I should handle the memory */

func DearPointer() {
	x := 67
	i := &x
	
	pow(&x) // This modifies x by squaring it
	/*pow(i) : This does the same as above because i and &x point to the same location
	This means x is squared twice*/
	
	fmt.Println("This is i:", *i) // Dereference i to get the value of x
	fmt.Println("This is x:", x)   // Print x directly, which holds the modified value
}

func pow(p *int) {
	*p = *p * *p // Dereference p to access and modify the value it points to
}


type collection struct{
	x string
	y string
	z string
}

func CallCollection(){
	c := collection{"I","love","go"}
	c.x = "everyone"
	c.y = c.y + "s"
	p := &c
	(*p).z = "GO" // equal to p.z
	fmt.Println(c)
	v := collection{}
	fmt.Println(v)
}



type Vertex struct {
	m int
	n interface{}

}

func inttpe(){

	var(
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{n: "3"}  
	v3 = Vertex{}     
	p  = &Vertex{1, 65.0005} // has type *Vertex
	)
	p.n = "I can be any typey"
	
	fmt.Println(v1,*p,v2,v3)
}
