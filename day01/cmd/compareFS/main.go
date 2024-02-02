// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"s21/internal/fscompare"
// )

// func main() {
// 	oldFilePath := flag.String("old", "", "Path to the old filesystem dump")
// 	newFilePath := flag.String("new", "", "Path to the new filesystem dump")
// 	flag.Parse()

// 	if *oldFilePath == "" || *newFilePath == "" {
// 		fmt.Println("Usage: ./compareFS --old <snapshot1.txt> --new <snapshot2.txt>")
// 		os.Exit(1)
// 	}

// 	// err := fscompare.CompareFileSets(*oldFilePath, *newFilePath)
// 	fileSet, err := fscompare.ReadAndCompress(*oldFilePath)

// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		os.Exit(1)
// 	}

// }
package main

import (
	"flag"
	"fmt"
	"s21/internal/fscompare"
)

func main() {
	oldFilePath := flag.String("old", "", "Path to the old filesystem dump")
	newFilePath := flag.String("new", "", "Path to the new filesystem dump")
	flag.Parse()

	if *oldFilePath == "" || *newFilePath == "" {
		fmt.Println("Usage: ./compareFS --old <snapshot1.txt> --new <snapshot2.txt>")
		return
	}

	hashMap, err := fscompare.GenerateHashMap(*oldFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := fscompare.CompareFS(hashMap, *newFilePath, *oldFilePath); err != nil {
		fmt.Println(err)
		return
	}
	return
}
