package main

import (
	"day02/internal/finder"
	"fmt"
)

func main() {
	err := finder.FindProcess()
	if err != nil {
		fmt.Println(err)
	}
	return
}
