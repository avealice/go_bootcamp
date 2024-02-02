package main

import (
	"day02/internal/archiver"
	"fmt"
)

func main() {
	// archiveDir := flag.String("a", "", "Archive directory path")
	// flag.Parse()

	// if *archiveDir == "" {
	// 	fmt.Println("Usage: ./myRotate [-a archive_dir] log_file1 log_file2 ...")
	// 	os.Exit(1)
	// }

	// logFiles := flag.Args()

	// if len(logFiles) == 0 {
	// 	fmt.Println("No log files specified.")
	// 	os.Exit(1)
	// }

	// var wg sync.WaitGroup

	// for _, logFile := range logFiles {
	// 	wg.Add(1)
	// 	go archiver.RotateAndArchive(logFile, *archiveDir, &wg)
	// }

	// wg.Wait()
	err := archiver.RotateProcess()
	if err != nil {
		fmt.Println(err)
	}
}
