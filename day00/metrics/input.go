package metrics

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Input(reader io.Reader) []int {
	var numbers []int
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		num, err := strconv.Atoi(strings.TrimSpace(line))

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		if -100000 > num || num > 100000 {
			fmt.Println("Number out of range. Please enter an integer between -100000 and 100000.")
			continue
		}

		numbers = append(numbers, num)
	}
	sort.Ints(numbers)

	return numbers
}
