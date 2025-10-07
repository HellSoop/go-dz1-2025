package main

import (
	"fmt"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)

func main() {
	err := task2.FindCommonWords("test.txt", "../internal/task_2/files/good.txt", "../internal/task_2/files/bad.txt")
	if err != nil {
		fmt.Println(err)
	}

	a := []int{1, 2, 3}
	var p *[]int = &a
	task3.ScaleSlice(p, 0)
	fmt.Println(a)
}
