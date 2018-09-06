package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	result := sumParameters(text)
	fmt.Println(result)
}

func sumParameters(inputConsole string) (total int) {

	number, err := strconv.Atoi(inputConsole)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	for i := 0; i < number; i++ {
		if (i%3) == 0 || (i%5) == 0 {
			total += i
		}
	}

	return
}
