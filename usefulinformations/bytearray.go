package usefulinformations

import "fmt"

// ByteArray is
func ByteArray() {

	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)

	fmt.Println(strb)

	count := 10
	fmt.Println(int64(count))
}
