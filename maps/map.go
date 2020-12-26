package maps

import "fmt"

// MapExOne is a method
func MapExOne() {
	fmt.Println("maps")

	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exits := lookup["vegeta"]
	fmt.Println(power, exits)

	total := len(lookup)
	fmt.Println(total)

	delete(lookup, "goku")
	fmt.Println(lookup)

	lookupOne := make(map[string]int, 100)
	fmt.Println(lookupOne)

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	// goku.Friends["krillin"]  ??
	fmt.Println(goku)

	lookupTwo := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	for key, value := range lookupTwo {
		fmt.Println(key, value)
	}
}

// Saiyan as a struct
type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}
