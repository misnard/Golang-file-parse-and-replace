package main

import (
	"awesomeProject/helpers"
	"fmt"
)



//This program can be improved by support {oldWord}. / {oldWord}, ect..
func main() {
	fmt.Printf("File search and replace\n")
	occ, lines, err := helpers.FindReplaceFile("./input.txt", "Go", "python")
	if err != nil {
		fmt.Printf("An error occured %s", err)
	}

	fmt.Printf("Find, %d occurences at lines : %v", occ, lines)
}


