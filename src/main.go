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

func handler(w http.ResponseWriter, r *http.Request) {
	
	printMessage()
	
	var x float64 = 0.0001
	for {
		x += doSqrt(x)
	}
}

func doSqrt(value float64) float64 {
	return math.Sqrt(value)
} 

func printMessage() {
	fmt.Printf("Code.education Rocks!")
}

func validateQueryString(input string) []string {
	errors := []string{}

	if input == "" {
		errors = append(errors, "The input 'times' cannot be empty")
	}

	if _, err := strconv.ParseFloat(input, 64); err != nil {
		errors = append(errors, "The input 'times' must be a float")
	}

	return errors
}
