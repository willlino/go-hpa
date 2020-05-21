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
	valueInString := r.URL.Query().Get("value")

	validateQueryString(valueInString)

	value, _ := strconv.ParseFloat(valueInString, 64)

	result := doSqrt(value)
	fmt.Printf("%v", result)
	printMessage()
}

func doSqrt(value float64) float64 {
	return math.Sqrt(value)
}

func printMessage() {
	fmt.Printf("William Rocks!")
}

func validateQueryString(input string) []string {
	errors := []string{}

	if input == "" {
		errors = append(errors, "The input 'value' cannot be empty")
	}

	if _, err := strconv.ParseFloat(input, 64); err != nil {
		errors = append(errors, "The input 'value' must be a float")
	}

	return errors
}
