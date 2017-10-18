package main

//create new functions in seperate files under fileUtils
//fileUtils/reverse.go is to get an idea of how to create a new function
//Test output of functions in here
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"./fileUtils"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	buffer.WriteString(wd)
	buffer.WriteString("\\LICENSE")
	//windows file seperator replace
	wd = strings.Replace(buffer.String(), "\\", "/", -1)

	fmt.Println(fileUtils.GetFileSize(wd))
	fmt.Println(fileUtils.GetMD5(wd))

}
