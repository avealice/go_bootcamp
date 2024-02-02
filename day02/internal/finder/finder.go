package finder

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	fileFlag    = flag.Bool("f", false, "Print only files")
	dirFlag     = flag.Bool("d", false, "Print only directories")
	symlinkFlag = flag.Bool("sl", false, "Print only symlinks")
	extFlag     = flag.String("ext", "", "Print only files with the specified extension")
)

func FindProcess() error {
	flag.Parse()

	dirPath := flag.Arg(0)

	if dirPath == "" {
		fmt.Println("Usage: ./myFind [-f] [-d] [-sl] [-ext extension] /path/to/dir")
		return nil
	}

	if !*fileFlag && !*dirFlag && !*symlinkFlag && *extFlag == "" {
		*fileFlag, *dirFlag, *symlinkFlag = true, true, true
	} else if *extFlag != "" && !*fileFlag {
		fmt.Println("Use -ext with -f")
		return nil
	}

	err := filepath.Walk(dirPath, walkProcess)

	return err
}

func walkProcess(path string, info os.FileInfo, err error) error {
	if os.IsPermission(err) {
		return filepath.SkipDir
	} else if err != nil {
		return err
	}

	if info.IsDir() && *dirFlag {
		if path == flag.Arg(0) {
			return nil
		}
		fmt.Println(path)
	} else if info.Mode().IsRegular() && *fileFlag {
		if filepath.Ext(path) == "."+*extFlag || *extFlag == "" {
			fmt.Println(path)
		}
	} else if info.Mode()&os.ModeSymlink != 0 && *symlinkFlag {
		dest, err := filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Println(path, "->", "[broken]")
		} else {
			fmt.Println(path, "->", dest)
		}
	}

	return nil
}
