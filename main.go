package main
//create new functions in seperate files under fileUtils
//fileUtils/reverse.go is to get an idea of how to create a new function
//Test output of functions in here
import (
	"fmt"

	"./fileUtils"
)

func main() {
	fmt.Printf(fileUtils.Reverse("!oG ,olleH"))
}