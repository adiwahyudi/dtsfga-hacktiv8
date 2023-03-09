package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true
	var k float64 = 123.456

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%t \n", j)
	fmt.Printf("\u042F (ya)\n")
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n", 'Я')
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)

}
