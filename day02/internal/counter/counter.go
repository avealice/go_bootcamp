package counter

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"unicode/utf8"
)

type CountRes struct {
	FileName string
	Count    int
}

var (
	linesFlag      = flag.Bool("l", false, "Count lines")
	charactersFlag = flag.Bool("m", false, "Count characters")
	wordsFlag      = flag.Bool("w", false, "Count words")
)

func CountProcess() error {
	flag.Parse()

	countFlag := 0

	if *linesFlag {
		countFlag++
	}
	if *charactersFlag {
		countFlag++
	}
	if *wordsFlag {
		countFlag++
	}

	if countFlag == 0 {
		*wordsFlag = true
	} else if countFlag > 1 {
		fmt.Println("Error: Please select one counting option (-l, -m or -w).")
		return nil
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Using: ./myWc [-l] [-w] [-m] file1 file2 ...")
		return nil
	}

	var wg sync.WaitGroup

	res := make(chan CountRes)

	for _, file := range files {
		wg.Add(1)
		localFile := file
		go counting(localFile, res, &wg)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Printf("%d\t%s\n", r.Count, r.FileName)
	}

	return nil
}

func counting(file string, result chan<- CountRes, wg *sync.WaitGroup) {
	defer wg.Done()

	fileDescriptor, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileDescriptor.Close()

	count := CountRes{
		FileName: file,
		Count:    0,
	}

	scanner := bufio.NewScanner(fileDescriptor)
	if *wordsFlag {
		scanner.Split(bufio.ScanWords)
	}
	for scanner.Scan() {
		if *charactersFlag {
			count.Count += utf8.RuneCountInString(scanner.Text())
		} else {
			count.Count++
		}
	}

	result <- count
}
