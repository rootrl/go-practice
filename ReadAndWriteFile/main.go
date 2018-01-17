package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// to byte format
	content := []byte("hello \n just for test")

	// if test.txt file not exists, will create it with 0644 permission
	err := ioutil.WriteFile("test.txt", content, 0644)

	if err != nil {
		panic(err)
	}

	// data is  byte format
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}

	// to string
	fmt.Println(string(data))

}
