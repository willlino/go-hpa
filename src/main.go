package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":6800", nil)
	
}

func handler(w http.ResponseWriter, r *http.Request){
	timesInString := r.URL.Query().Get("times")
	
	times, _ := strconv.Atoi(timesInString)
	
	doSqrt(times);
	printMessage();
}

func doSqrt(times int){
	x := 0.0001
	for i := 0; i <= times; i ++ {
		x += math.Sqrt(x)
		fmt.Printf("%v" , x)
	}
}

func printMessage(){
	fmt.Printf("William Rocks!")
}