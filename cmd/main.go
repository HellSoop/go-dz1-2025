package main

import (
	"fmt"
	"dz1/internal/task_2"
)

func main() {
	err := task2.FindCommonWords("test.txt", "../internal/task_2/files/good.txt", "../internal/task_2/files/bad.txt")
	if err != nil {
		fmt.Println(err)
	}
}
