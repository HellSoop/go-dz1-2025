package main

import (
	"fmt"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)

func main() {
	err := task2.FindCommonWords("test.txt", "../internal/task_2/files/devil_trigger.txt", "../internal/task_2/files/bury_the_light.txt")
	if err != nil {
		fmt.Println(err)
	}

	a := []int{1, 2, 3}
	task3.ScaleSlice(&a, 5)
	fmt.Println(a)
}
