package main

import "fmt"

func main() {
		var result = convert(1)
		var expectedResult = "I"

		if result == "I" {
			fmt.Println("Expected result ", expectedResult, "actual result was ", result)
			fmt.Println("TEST PASSED")
		} else {
			fmt.Println("Expected result ", expectedResult, "actual result was ", result)
			fmt.Println("TEST FAILED")
		}

		result = convert(2)
		expectedResult = "II"

		if result == "II" {
			fmt.Println("Expected result ", expectedResult, "actual result was ", result)
			fmt.Println("TEST PASSED")
		} else {
			fmt.Println("Expected result ", expectedResult, "actual result was ", result)
			fmt.Println("TEST FAILED")
		}
		
}

func convert(number int) string {
	var convertedNumeral string
	if number == 1 {
		convertedNumeral = "I"
	} else if number == 2 {
		convertedNumeral = "II"
	}
	return convertedNumeral
}

