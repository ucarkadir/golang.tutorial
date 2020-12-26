package slices

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Saiyan is a struct.
type Saiyan struct {
	Name  string
	Power int
}

/*
	1. names := []string{"leto", "jessica", "paul"}
	2. checks := make([]bool, 10)
	3. var names []string
	4. scores := make([]int, 0, 20)
*/

//
// SliceMethodTwo is a method.
// Dizide istedi
// ğimiz değerlerini Önceden Bildiğimizde
func SliceMethodTwo() {
	names := []string{"leto", "jessica", "paul"}
	fmt.Println(names)
}

//
// SliceMethodThree is a method.
// Bir dilimin
// belirli indeksine yazarken kullanır
func SliceMethodThree(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}

//
// SliceMethodFour is a method.
//
//
func SliceMethodFour() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores) // [1, 2, 999, 4, 5]

	haystack := "the spice must flow"
	index := strings.Index(haystack[7:], " ") // 7. indexten sonuna kadar
	fmt.Println(index)

	scores = scores[:len(scores)-1]
	fmt.Println(scores)

	scores = removeAtIndex(scores, 3)
	fmt.Println(scores)
}

//
// SliceMethodFive is a method
//
func SliceMethodFive() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
}

// removeAtIndex is a method
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}
