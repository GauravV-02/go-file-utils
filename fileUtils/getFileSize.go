// Package fileUtils contains utility functions for working with files.
package fileUtils

import (
	"os"
	"log"
)

func GetFileSize(pathToFile string) int64 {
	file, err := os.Stat(pathToFile);

	//if file can't be found create an err message and close program
	if os.IsNotExist(err) {	
		log.Fatal("File does not exsist at: ", pathToFile)
	}
	//if any filesystem error occurs close the program with err details
	if err != nil {
		log.Fatal(err)
	}

	// get and return the size of the file
	return file.Size()
}