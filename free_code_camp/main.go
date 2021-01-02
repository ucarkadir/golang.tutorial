package freeCodeCamp

import (
	"fmt"
	"strconv"
)

var (
	actorName    string = "kadir"
	doctorNumber int    = 3
)

// const genellikle package boyuntunda tanımlanır
const myConst = iota

const (
	erorSpecialist = iota
)

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota // ignore firs vaule by assignigng to blank identifier
	// KB is a const
	KB = 1 << (10 * iota)
	// MB is a const
	MB
	// GB is a const
	GB
	// TB is a const
	TB
	// PB is a const
	PB
	// EB is a const
	EB
	// ZB is a const
	ZB
	// YB is a const
	YB
)

// Tutorial is a method
func main() {

	i := 12
	fmt.Printf("%v, %T", i, i) // value, type

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T", j, j) // int to float

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T", k, k) // int to string

	s := " this is a strign"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	r := 'a'
	fmt.Printf("%v %T", r, r) // 97, int32

	rune1 := '\a'
	fmt.Printf("Rune 1: %c; Unicode: %U;", rune1, rune1)

	// Constants
	const myConst int = 42
	fmt.Printf("%v, %T", myConst, myConst)

	const myConst2 = 42
	fmt.Printf("%v, %T", myConst2, myConst2)

	fmt.Printf("%v \n", catSpecialist)   // 6
	fmt.Printf("%v \n", dogSpecialist)   // 7
	fmt.Printf("%v \n", snakeSpecialist) // 8

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/float64(GB)) //3.73GB
}
