package archiver

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	archiveFlag = flag.Bool("a", false, "Archive directory path")
)

func RotateProcess() error {
	flag.Parse()

	archiveDir, logFiles := checkArgs()

	var wg sync.WaitGroup

	for _, logFile := range logFiles {
		wg.Add(1)
		go archiveProcess(logFile, archiveDir, &wg)
	}

	wg.Wait()
	return nil
}

func checkArgs() (string, []string) {
	files := flag.Args()
	var dirPath string
	if (*archiveFlag && len(files) < 3) || len(files) == 0 {
		fmt.Println("Usage: ./myRotate [-a archive_dir] log_file1 log_file2 ...")
		return dirPath, nil
	} else if *archiveFlag && len(files) > 2 {
		_, err := os.Stat(files[0])
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist")
			return dirPath, nil
		} else if err != nil {
			fmt.Println(err)
			return dirPath, nil
		}
		dirPath = files[0]
		files = files[1:]
	}
	return dirPath, files
}

func archiveProcess(filePath string, dirPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := filepath.Base(filePath)
	fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]
	if dirPath == "" {
		dirPath = filepath.Dir(filePath)
	}
	archiveFilePath := fmt.Sprintf("%s/%s_%s.%s", dirPath, fileName, strconv.FormatInt(fileInfo.ModTime().Unix(), 10), "tar.gz")

	file, err := os.Create(archiveFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	fileDescriptor, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	header := &tar.Header{
		Name: filepath.Base(filePath),
		Size: fileInfo.Size(),
		Mode: int64(fileInfo.Mode()),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := io.Copy(tarWriter, fileDescriptor); err != nil {
		fmt.Println(err)
		return
	}

}
