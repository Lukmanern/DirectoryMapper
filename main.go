package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var baseDir string = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"

func main() {
	root := baseDir
	printDirectoryMap(root, "   ", ".git")
	// err := filepath.Walk(root, visit)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func printDirectoryMap(path string, prefix string, exclude string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for i, file := range files {
		if exclude == file.Name() {
			continue
		}
		if file.IsDir() {
			fmt.Println(prefix + "├── " + file.Name() + " -- 1")
			printDirectoryMap(filepath.Join(path, file.Name()), prefix+"│   ", exclude)
		} else {
			fmt.Println(prefix + "├── " + file.Name() + " -- 2")
		}
		if i == len(files)-1 {
			fmt.Println(prefix + "└── " + file.Name() + " -- 3")
			printDirectoryMap(filepath.Join(path, file.Name()), prefix+"    ", exclude)
		}
	}
	return nil
}

// func ErrorHandler(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func runner() {
// 	fileInfo, err := os.Stat(baseDir)
// 	ErrorHandler(err)
// 	if ! fileInfo.IsDir() {
// 		err = errors.New(baseDir + " isn't Folder/ Directory.")
// 		ErrorHandler(err)
// 	}
// }