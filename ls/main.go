package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {

	if e != nil {
		panic(e)
	}

}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing dir name")
		os.Exit(1)
	}

	dirName := args[1]

	files, err := ioutil.ReadDir(dirName)

	check(err)

	for _, file := range files {

		fmt.Println(file.Name())

	}

}
