package main

import (
	"DirectoryMapper/banner"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()
	banner.ShowBanner()
	var baseDir string = "C:/Users/Lenovo/go/src/DB_CLI"
	// baseDir = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"
	printDirectoryMap(0, baseDir, "   ", ".git")
	finish := time.Since(start)
	fmt.Println("\n Mapping duration: " + finish.String())
}

func printDirectoryMap(counter uint, path, prefix, exclude string) {
	// max 100+ stacks
	if counter >= 100 {
		fmt.Println("::::::::::MAX PRINT::::::::::")
		return
	}
	// Read all the files/ folders in the current directory.
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
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