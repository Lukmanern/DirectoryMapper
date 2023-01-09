package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var baseDir string = "C:/xampp/htdocs/mengundang/app"

func main() {
	baseDir = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"
	root := baseDir
	printDirectoryMap(1, root, "   ", ".git")
}

func printDirectoryMap(counter int, path, prefix, exclude string) error {
	if counter >= 8 {
		return nil
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for i, file := range files {
		if exclude == file.Name() {
			continue
		}
		if i == len(files)-1 {
			fmt.Println(prefix + "└── " + file.Name())
			if file.IsDir() {
				printDirectoryMap(counter+1, filepath.Join(path, file.Name()), prefix+"    "+"", exclude)
			}
		} else {
			fmt.Println(prefix + "├── " + file.Name())
			if file.IsDir() {
				printDirectoryMap(counter+1, filepath.Join(path, file.Name()), prefix+"│   "+"", exclude)
			}
		}
	}

	return nil
}