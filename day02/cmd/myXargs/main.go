package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := os.Args[1:]

	scanner := bufio.NewScanner(os.Stdin)

	var args []string
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}

	cmd := exec.Command(command[0], append(command[1:], args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
