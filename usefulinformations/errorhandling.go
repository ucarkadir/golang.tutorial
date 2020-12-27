package usefulinformations

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ErrorHandling is
func ErrorHandling() {

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	var asArgs = os.Args[1]
	n, err := strconv.Atoi(asArgs)

	if err != nil {
		fmt.Println("not a valid number")
	} else {
		fmt.Println(n)
	}

}

// Process is a return methdo
func Process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}

// EOFMethod is
func EOFMethod() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	}
}

/*
	panic : bir istisna atmaya benzerken
	recover: catch kullanımına benzer ve nadiren kullanılırlar.
*/
