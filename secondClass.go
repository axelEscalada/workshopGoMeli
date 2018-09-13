package main

import (
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{value, nil, nil}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func walk(t *tree) {
	if t == nil {
		return
	}

	walk(t.left)
	fmt.Println(t.value)

	walk(t.right)

}

func populateTree(nums []int) *tree {
	var t *tree
	for _, value := range nums {
		t = add(t, value)
	}

	return t
}

func parseToInt(sliceString []string) []int {
	sliceInt := make([]int, len(sliceString))
	fmt.Println(sliceString)
	for index, value := range sliceString {
		i, err := strconv.Atoi(value)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		sliceInt[index] = i
	}
	return sliceInt
}

func main() {
	nums := os.Args[1:]
	sliceInt := parseToInt(nums)
	t := populateTree(sliceInt)
	walk(t)
}
