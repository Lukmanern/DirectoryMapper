package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	var baseDir string = "C:/xampp/htdocs/mengundang/app"
	// baseDir = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"
	printDirectoryMap(1, baseDir, "   ", ".git")
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printDirectoryMap(counter int, path, prefix, exclude string) {
	if counter >= 8 {
		fmt.Println("::::::::::MAX PRINT::::::::::")
		return
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = errors.New(fmt.Sprint(err.Error(), "at counter:", counter))
		ErrorHandler(err)
	}
	for i, file := range files {
		if exclude == file.Name() {
			continue
		}
		// 3 white space
		space := "   "
		if i == len(files)-1 {
			fmt.Println(prefix + "└── " + file.Name())
			space = " " + space
		} else {
			fmt.Println(prefix + "├── " + file.Name())
			space = "│" + space
		}
		if file.IsDir() {
			printDirectoryMap(counter+1, filepath.Join(path, file.Name()), prefix+space, exclude)
		}
	}
}