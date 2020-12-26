package slices

import "fmt"

// SliceMethodZero is a method
func SliceMethodZero() {
	fmt.Println("Dilimler")

	// Slices
	slicesScores := []int{1, 4, 293, 4, 9}
	fmt.Println(slicesScores)

	// Slices Make ile uzunluğu (dilmin boyutu) 10 olan dilim oluşturulur.
	slicesScoresMake := make([]int, 10)
	fmt.Println(slicesScoresMake)

	// Uzunluğu 0, Kapasitesi 10 olan bir dilim oluşturmu oluruz.
	slicesScoresMakeLength := make([]int, 0, 10)
	fmt.Println(slicesScoresMakeLength)

	slicesScoresMakeExOne := make([]int, 1, 10)
	slicesScoresMakeExOne[0] = 9033
	fmt.Println(slicesScoresMakeExOne)

	slicesScoresMakeExTwo := make([]int, 0, 10)
	slicesScoresMakeExTwo = append(slicesScoresMakeExTwo, 5)
	fmt.Println(slicesScoresMakeExTwo)
	slicesScoresMakeExTwo = slicesScoresMakeExTwo[0:8]
	slicesScoresMakeExTwo[7] = 9033
	fmt.Println(slicesScoresMakeExTwo)

}

// SliceMethodOne is
func SliceMethodOne() {
	fmt.Println("Dilimler Geniş Boyutlu")

	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)

	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}

	scoresNew := make([]int, 5)
	scoresNew = append(scoresNew, 9322)
	fmt.Println(scoresNew)
}
