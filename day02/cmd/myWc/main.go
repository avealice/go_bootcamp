package main

import (
	"day02/internal/counter"
	"fmt"
)

func main() {
	err := counter.CountProcess()
	if err != nil {
		fmt.Println(err)
	}

	return
}
