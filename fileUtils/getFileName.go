// Package fileUtils contains utility functions for working with files.
package fileUtils

import (
	"log"
	"os"
)

func GetFileName(pathToFile string) string {
	file, err := os.Stat(pathToFile)

	if os.IsNotExist(err) {
		log.Fatal("File does not exsist at: ", pathToFile)
	}

	if err != nil {
		log.Fatal(err)
	}
// get and return the name of the file
	return file.Name()
}