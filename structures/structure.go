package structures

import "fmt"

// StructureMethodOne is
func StructureMethodOne() {
	fmt.Println("Yapılar")

	syn := Saiyan{
		Name:  "Kadir",
		Power: 5,
	}

	syn2 := Saiyan{}

	syn3 := Saiyan{Name: "Veli"}
	syn3.Power = 9000

	fmt.Printf("%s isimli kullanıcının gücü: %d", syn.Name, syn.Power)
	fmt.Printf("%s isimli kullanıcının gücü: %d", syn2.Name, syn2.Power)
	fmt.Printf("%s isimli kullanıcının gücü: %d", syn3.Name, syn3.Power)

	syn4 := Saiyan{"Ayşe", 9000}
	Super(syn4)
	fmt.Println(syn4.Power)

	syn5 := &Saiyan{"And", 9000}
	Super2(syn5)
	fmt.Printf("%s isimli kullanıcının gücü: %d", syn5.Name, syn5.Power)

	Super3(syn5)
	fmt.Printf("%s isimli kullanıcının gücü: %d", syn5.Name, syn5.Power)

	syn6 := &Saiyan{"Bu son", 9001}
	syn6.Super4()
	fmt.Printf("%s isimli kullanıcının gücü: %d", syn6.Name, syn6.Power)

}

// Saiyan is a struct. It is help us.
type Saiyan struct {
	Name  string
	Power int
}

// Super is a method. Culculate our values
func Super(s Saiyan) {
	s.Power += 10000
}

// Super2 is a method. Culculate our values
// Dışarıdan gelen nesne üzerinden işlem yap
func Super2(s *Saiyan) {
	s.Power += 10000
}

// Super3 is a method.
func Super3(s *Saiyan) {
	s = &Saiyan{"Gohan", 1000}
}

// Super4 is a method.
func (s *Saiyan) Super4() {
	s.Power += 10000
}
