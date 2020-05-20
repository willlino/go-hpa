package main

import (
	"fmt"
	"math"
)

func main() {
	
	doSqrt(10);
	printMessage();
}

func doSqrt(times int){
	x := 0.0001
	for i := 0; i <= times; i ++ {
		x += math.Sqrt(x)
	}
}

func printMessage(){
	fmt.Printf("William Rocks!")
}