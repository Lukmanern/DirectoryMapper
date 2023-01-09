package main

import (
	"DirectoryMapper/banner"
	"fmt"
	"log"
	"os"
)

var baseDir string = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Directory Mapper"

func main() {
	banner.ShowBanner()
	fmt.Println(isDirectoryExist(baseDir))
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isDirectoryExist(dir string) bool {
	p, err := os.Stat(dir)
	ErrorHandler(err)
	fmt.Println(p.IsDir())
	return true
}