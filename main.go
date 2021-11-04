package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


// a function that creates a new iso file, writes and reads from it
func main() {
	// create a new iso file
	file, err := os.Create("isostore.iso")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// write some data to the file
	file.WriteString("Hello World")

	// read the data from the file
	data, err := ioutil.ReadFile("isostore.iso")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// create a new iso file
	file, err = os.Create("isostore.iso")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// write some data to the file
	file.WriteString("Hello World")

	// read the data from the file
	data, err = ioutil.ReadFile("isostore.iso")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
