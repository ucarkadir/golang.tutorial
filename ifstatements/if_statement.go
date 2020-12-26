package ifstatements

import (
	"fmt"
	"os"
)

// IfStatementMethodOne is
func IfStatementMethodOne() {
	var name string = "Lego2"
	var age int32 = 4
	check := true

	if name == "Lego" {
		fmt.Print("the spice must flow")
	}

	if (name == "Lego2" && age > 3) || (check) {
		fmt.Print("super saiyan")
	}

	/*if len(os.Args) != 2 {
		os.Exit(1)
	}*/

	fmt.Println(os.Args[0])

	var power int
	power = 9000

	fmt.Printf("It's over %d\n", power)

	power2 := getPower()
	fmt.Println(power2)

	power2 = 9000
	fmt.Print(power2)

	name2, power2 := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name2, power2)

	power3, name2 := 9000, "Goku"
	fmt.Printf("%s's power is over %d\n", name2, power3)

	value, exits := getPower2("goku")

	if exits == false {
		fmt.Printf("%d not support\n", value)
	}

	_, exits2 := getPower2("kadir")
	if exits2 == false {
		fmt.Println(exits)
	}
}

func getPower() int {
	return 9001
}

func log(message string) {

}

func add(name string) int {
	return 0
}

func getPower2(name string) (int, bool) {
	return 123, false
}
