package main

import (
	"fmt" 
	"strings"
)

var (
	hundred = Numeral{
		romanNumeral: "C",
		integerValue: 100,
	}
	ninety = Numeral{
		romanNumeral: "XC",
		integerValue: 90,
	}
	fifty = Numeral{
		romanNumeral: "L",
		integerValue: 50,
	}
	forty = Numeral{
		romanNumeral: "XL",
		integerValue: 40,
	}
	ten = Numeral{
		romanNumeral: "X",
		integerValue: 10,
	}
	nine = Numeral{
		romanNumeral: "IX",
		integerValue: 9,
	}
	five = Numeral{
		romanNumeral: "V",
		integerValue: 5,
	}
	four = Numeral{
		romanNumeral: "IV",
		integerValue: 4,
	}
	one = Numeral{
		romanNumeral: "I",
		integerValue: 1,
	}
)

func main() {

	var testsFailed = 0

		test(1, "I", &testsFailed)
		test(2, "II", &testsFailed)
		test(3, "III", &testsFailed)
		test(4, "IV", &testsFailed)
		test(5, "V", &testsFailed)
		test(8, "VIII", &testsFailed)
		test(9, "IX", &testsFailed)
		test(10, "X", &testsFailed)
		test(37, "XXXVII", &testsFailed)
		test(49, "XLIX", &testsFailed)
		test(86, "LXXXVI", &testsFailed)
		test(90, "XC", &testsFailed)
		test(399, "CCCXCIX", &testsFailed)

	printAllTestsResult(testsFailed)
}

func test(subject int, expectedResult string, testsFailed *int) {
	var result = convert(subject)
	if !testConversion(expectedResult, result) {
		*testsFailed++
	}
}

type Numeral struct {
	romanNumeral string
	integerValue int
	lowerNumeral string
	lowerInteger int
}

func convert(number int) string {
	convertedNumeral := ""
	numerals := [9]Numeral{hundred, ninety, fifty, forty, ten, nine, five, four, one}

	for _, numeral := range numerals {
		quotient := number / numeral.integerValue 
		number = number % numeral.integerValue
		convertedNumeral += strings.Repeat(numeral.romanNumeral, quotient)
	}

	return convertedNumeral
}

func testConversion(expectedResult string, actualResult string) bool {
	if expectedResult == actualResult {
		fmt.Println("Expected result ", expectedResult, "actual result was ", actualResult)
		fmt.Println("TEST PASSED")
		return true
	} else {
		fmt.Println("Expected result ", expectedResult, "actual result was ", actualResult)
		fmt.Println("TEST FAILED")
		return false
	}
}

	func printAllTestsResult(testsFailed int) {
		if testsFailed == 0 {
			fmt.Println("ALL TESTS PASSING")
		} else {
			fmt.Println(testsFailed, " tests failing")
		}
	}


