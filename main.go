package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	var baseDir string = "C:/xampp/htdocs/mengundang/app"
	// baseDir = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"
	printDirectoryMap(1, baseDir, "   ", "node_modules")
	finish := time.Since(start)
	fmt.Println("\n Mapping duration: " + finish.String())
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printDirectoryMap(counter int, path, prefix, exclude string) {
	// max 100 stacks
	if counter >= 100 {
		fmt.Println("::::::::::MAX PRINT::::::::::")
		return
	}
	// Read all the files/ folders in the current directory.
	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = errors.New(fmt.Sprint(err.Error(), "at counter:", counter))
		ErrorHandler(err)
	}

	for i, file := range files {
		// Skip exclude files
		if exclude == file.Name() {
			continue
		}
		var space string
		// Determine the appropriate prefix to use based on 
		// whether the current file is the last in the list.
		if i == len(files)-1 {
			fmt.Println(prefix + "└── " + file.Name())
			// 4 white space
			space = "    "
		} else {
			fmt.Println(prefix + "├── " + file.Name())
			// line + 3 white space
			space = "│   "
		}
		// If the current file is a directory, 
		// recursively call printDirectoryMap on it.
		if file.IsDir() {
			printDirectoryMap(counter+1, filepath.Join(path, file.Name()), prefix+space, exclude)
		}
	}
}