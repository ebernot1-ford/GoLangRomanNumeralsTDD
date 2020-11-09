package main

import "fmt"

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

func convert(number int) string {
	convertedNumeral := ""

	if number - 100 >= 0 {
		convertedNumeral += "C"
		number -= 100
	} else if number - 90 >= 0 {
		convertedNumeral += "XC"
		number -= 90
	} else if number - 50 >= 0 {
		convertedNumeral += "L"
		number -= 50
	} else if number - 40 >= 0 {
		convertedNumeral += "XL"
		number -= 40
	} else if number - 10 >= 0 {
		convertedNumeral += "X"
		number -= 10 
	} else if number - 9 >= 0 {
		convertedNumeral += "IX"
		number -= 9
	} else if number - 5 >= 0 {
		convertedNumeral += "V"
		number -= 5
	} else if number - 4 >= 0 {
		convertedNumeral += "IV"
		number -= 4
	} else if number - 1 >= 0 {
		convertedNumeral += "I"
		number -= 1
	} 

	if number > 0 {
		convertedNumeral += convert(number)
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


