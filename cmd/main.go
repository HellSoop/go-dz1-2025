package main

import (
	"fmt"
	"os"
	"dz1/internal/task_1"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)


func main() {
	// Check task1
	fmt.Println("Task 1 tests:")
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
	fmt.Printf("FilterCommonDigits(123, 123) = (%d, %d, %v), expected: (0, 0, ErrEmptyNum)\n", firstNumber, secondNumber, err)

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
	
	// Example call 9:
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123456789, 987654321)
	fmt.Printf("FilterCommonDigits(123456789, 987654321) = (%d, %d, %v), expected: (0, 0, ErrEmptyNum)\n", firstNumber, secondNumber, err)

	// Example call 10:
	firstNumber, secondNumber, err = task1.FilterCommonDigits(123, 0)
	fmt.Printf("FilterCommonDigits(123, 0) = (%d, %d, %v), expected: (123, 0, nil)\n", firstNumber, secondNumber, err)

	// Example call 11:
	firstNumber, secondNumber, err = task1.FilterCommonDigits(1023, 145)
	fmt.Printf("FilterCommonDigits(1023, 145) = (%d, %d, %v), expected: (23, 45, nil)\n", firstNumber, secondNumber, err)

	// task 2 tests
	fmt.Println("\n\nTask 2 tests:")
	result_file := "internal/task_2/files/test.txt"
	// simple test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/devil_trigger.txt", "internal/task_2/files/bury_the_light.txt")
	data, _ := os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/devil_trigger.txt\", \"internal/task_2/files/bury_the_light.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// single file test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/case_test1.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/case_test1.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// three file test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/devil_trigger.txt", "internal/task_2/files/bury_the_light.txt", "internal/task_2/files/no_common1.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/devil_trigger.txt\", \"internal/task_2/files/bury_the_light.txt\", \"internal/task_2/files/no_common1.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// no files test
	err = task2.FindCommonWords(result_file)
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s) = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

		// identical files test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/case_test1.txt", "internal/task_2/files/case_test1.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/case_test1.txt\", \"internal/task_2/files/case_test1.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// empty files test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/empty1.txt", "internal/task_2/files/empty2.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/empty1.txt\", \"internal/task_2/files/empty2.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// no common files test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/no_common1.txt", "internal/task_2/files/no_common2.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/no_common1.txt\", \"internal/task_2/files/no_common2.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// different cases file test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/case_test1.txt", "internal/task_2/files/case_test2.txt")
	data, _ = os.ReadFile(result_file)
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/case_test1.txt\", \"internal/task_2/files/case_test2.txt\") = (%v), expected (nil)\n", result_file, err)
	fmt.Printf("%s file contains: \"%s\"\n", result_file, data)

	// invalid file test
	err = task2.FindCommonWords(result_file, "internal/task_2/files/invalid1.txt", "internal/task_2/files/anoter_invalid.txt")
	fmt.Printf("FindCommonWords(%s, \"internal/task_2/files/invalid1.txt\", \"internal/task_2/files/anoter_invalid.txt\") = (%v), expected (ErrOpenFile)\n", result_file, err)

	// invalid result file test
	err = task2.FindCommonWords("invalid/result/file", "internal/task_2/files/devil_trigger.txt", "internal/task_2/files/bury_the_light.txt")
	fmt.Printf("FindCommonWords(\"invalid/result/file\", \"internal/task_2/files/devil_trigger.txt\", \"internal/task_2/files/bury_the_light.txt\") = (%v), expected (ErrOpenFile)\n", err)

	//task 3 test
	fmt.Println("\n\nTask 3 tests:")
	
	// simple test
	a := []int{1, 2, 3}
	fmt.Printf("ScaleSlice(%v, 3) = ", a)
	err = task3.ScaleSlice(&a, 3)
	fmt.Printf("(%v), excepted (nil)\nSclaed slice: %v\n", err, a)

	// scaleFactor = 1
	a = a[:3]
	fmt.Printf("ScaleSlice(%v, 1) = ", a)
	err = task3.ScaleSlice(&a, 1)
	fmt.Printf("(%v), excepted (nil)\nSclaed slice: %v\n", err, a)

	// scaleFactor = 0
	fmt.Printf("ScaleSlice(%v, 0) = ", a)
	err = task3.ScaleSlice(&a, 0)
	fmt.Printf("(%v), excepted (nil)\nSclaed slice: %v\n", err, a)

	// empty slice
	fmt.Printf("ScaleSlice(%v, 10) = ", a)
	err = task3.ScaleSlice(&a, 10)
	fmt.Printf("(%v), excepted (nil)\nSclaed slice: %v\n", err, a)

	// overflow test
	a = make([]int, 4200000)
	err = task3.ScaleSlice(&a, 4200000)
	fmt.Printf("ScaleSlice(<slice of 4200000 zeros>, 4200000) = (%v), excepted (ErrOverflow)\n", err)
	a = nil  // free memory

	// nil slice
	err = task3.ScaleSlice(nil, 10)
	fmt.Printf("ScaleSlice(nil, 10) = (%v), excepted (nil)\n", err)
}
