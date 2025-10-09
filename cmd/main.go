package main

import (
	"fmt"
	"dz1/internal/task_1"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)

func main() {
	// Check task1
	// Example call 1: numbers are not the same
	firstNumber, secondNumber, err := task1.FilterCommonDigits(123, 456)
	fmt.Printf("FilterCommonDigits(123, 456) = (%d, %d, %v), expected: (123, 456, nil)\n", firstNumber, secondNumber, err)

	// Example call 2: numbers are the same with one digit
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123, 345)
	fmt.Printf("FilterCommonDigits(123, 345) = (%d, %d, %v), expected: (12, 45, nil)\n", firstNumber, secondNumber, err)

	// Example call 3: numbers are the same with two digit
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123, 234)
	fmt.Printf("FilterCommonDigits(123, 234) = (%d, %d, %v), expected: (1, 4, nil)\n", firstNumber, secondNumber, err)

	// Example call 4: numbers are the same
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123, 123)
	fmt.Printf("FilterCommonDigits(123, 123) = (%d, %d, %v), expected: (0, 0, ErrEmptyNUm)\n", firstNumber, secondNumber, err)

	// Example call 5: big numbers
	firstNumber, secondNumber, err = task1.FilterCommonDigits(12300405502345602, 9988699870045882)
	fmt.Printf("FilterCommonDigits(12300405502345602, 9988699870045882) = (%d, %d, %v), expected: (133, 998899878, nil)\n", firstNumber, secondNumber, err)

	// Example call 6: negative numbers
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123, -345)
	fmt.Printf("FilterCommonDigits(123, -345) = (%d, %d, %v), expected: (0, 0, ErrNegNums)\n", firstNumber, secondNumber, err)

	// Example call 7: negative numbers
	firstNumber, secondNumber, err = task1.FilterCommonDigits(-123, 345)
	fmt.Printf("FilterCommonDigits(-123, 345) = (%d, %d, %v), expected: (0, 0, ErrNegNums)\n", firstNumber, secondNumber, err)

	// Example call 8: negative numbers
	firstNumber, secondNumber, err = task1.FilterCommonDigits(-123, -345)
	fmt.Printf("FilterCommonDigits(-123, -345) = (%d, %d, %v), expected: (0, 0, ErrNegNums)\n", firstNumber, secondNumber, err)

	err = task2.FindCommonWords("intal/ts/test.txt", "internal/task_2/files/case_test1.txt", "internal/task_2/files/case_test2.txt")
	if err != nil {
		fmt.Println(err)
	}

	var a []int
	err = task3.ScaleSlice(&a, 3)
	fmt.Println(err)
}
