package arrays

import "fmt"

// ArrayMethodOne is
func ArrayMethodOne() {
	print("Diziler")

	var scores [10]int
	scores[0] = 339
	fmt.Println(scores)

	// Arrays
	newScores := [4]int{9001, 9333, 212, 33}
	fmt.Println(newScores)

	for index, value := range newScores {
		fmt.Printf("%d . sıradaki kayıt: %d \n", index+1, value)
	}
}
