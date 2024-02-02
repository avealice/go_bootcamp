package fscompare

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

type Data map[[32]byte]bool

func GenerateHashMap(filename string) (*Data, error) {
	hashMap := make(Data)
	fileDescriptor, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fileDescriptor.Close()

	scanner := bufio.NewScanner(fileDescriptor)
	for scanner.Scan() {
		line := scanner.Text()
		hash := sha256.Sum256([]byte(line))
		hashMap[hash] = false
	}
	return &hashMap, nil
}

func CompareFS(hashMap *Data, newFilePath, oldFilePath string) error {
	err := checkAdded(hashMap, newFilePath)
	if err != nil {
		return err
	}

	err = checkRemoved(hashMap, oldFilePath)

	if err != nil {
		return err
	}

	return nil
}

func checkAdded(hashMap *Data, filePath string) error {
	newSnap, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer newSnap.Close()
	// если мы нашли в хэшмап элемент по хэшу, то мы ставим true что просмотрели его и что он не изменем, если нет значит это новая добавленная строка
	scanner := bufio.NewScanner(newSnap)
	for scanner.Scan() {
		hash := sha256.Sum256(scanner.Bytes())
		if _, found := (*hashMap)[hash]; !found {
			fmt.Printf("ADDED %s\n", scanner.Text())
		} else {
			(*hashMap)[hash] = true
		}
	}

	return nil
}

func checkRemoved(hashMap *Data, filePath string) error {
	oldSnap, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer oldSnap.Close()
	// если индексирование по хэшу false значит в новом снэпе его нет и мы не смогли поставить ему true в функции checkAdded, поэтому он removed
	scanner := bufio.NewScanner(oldSnap)
	for scanner.Scan() {
		hash := sha256.Sum256(scanner.Bytes())
		if !(*hashMap)[hash] {
			fmt.Printf("REMOVED %s\n", scanner.Text())
		}
	}

	return nil
}
