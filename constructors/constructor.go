package constructors

import "fmt"

// ConstructorMethodOne is
func ConstructorMethodOne() {

	consts := NewSaiyan("Kadir", 100)
	fmt.Printf("%s isimli kullanıcının gücü: %d\n", consts.Name, consts.Power)

	newSaiyan := new(Saiyan)
	newSaiyan.Name = "Ali"
	newSaiyan.Power = 5

	fmt.Printf("%s isimli kullanıcının gücü: %d\n", newSaiyan.Name, newSaiyan.Power)

	newSaiyan2 := &Saiyan{
		Name:  "Veli",
		Power: 9,
	}

	fmt.Printf("%s isimli kullanıcının gücü: %d\n", newSaiyan2.Name, newSaiyan2.Power)
}

// Saiyan is a struct.
type Saiyan struct {
	Name  string
	Power int
}

// NewSaiyan is a const. method.
func NewSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
