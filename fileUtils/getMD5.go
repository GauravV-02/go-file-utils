package fileUtils

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

func GetMD5(pathToFile string) []byte {
	file, err := os.Open(pathToFile)

	//if file can't be found create an err message and close program
	if os.IsNotExist(err) {
		log.Fatal("File does not exsist at: ", pathToFile)
	}
	//if any filesystem error occurs close the program with err details
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	h := md5.New()

	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}
