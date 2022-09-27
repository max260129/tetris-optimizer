package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// check if the Os Args format is valid
	if len(os.Args) != 2 {
		fmt.Println("Wrong Usage !")
		return
	}

	// read file 
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Wrong file format !")
		return
	}

	// check pieces format
	ok, tetriminos := validationFile(data)
	if !ok {
		fmt.Println("Error of pieces format !")
		return
	}

	// solve
	solution := solve(tetriminos)

	// print solution
	solution.print()
}