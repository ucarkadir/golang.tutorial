package compositiontraitmixins

import "fmt"

// CompositionTraitMixinMethodOne is
func CompositionTraitMixinMethodOne() {
	fmt.Println("### İçerme - Composition ###")

	goku := &Saiyan{
		Person: &Person{
			Name: "Goku",
		},
		Power: 9001,
	}

	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)

	goku.Introduce()
}

// Person is a struct.
type Person struct {
	Name string
}

// Saiyan is a struct.
type Saiyan struct {
	*Person
	Power int
}

// Introduce is a method
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}
