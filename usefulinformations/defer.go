package usefulinformations

import (
	"fmt"
	"os"
)

// DeferMethod is
func DeferMethod() {
	file, err := os.Open("a_file_toread")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close()
	defer file.Close()
}
