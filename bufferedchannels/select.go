package bufferedchannels

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectMethod is a method.
func SelectMethod() {
	c := make(chan int)
	for {
		select {
		case c <- rand.Int():
			// buraya kodu yazılabilri
		default:
			// burası boş olar da barıkılabilir, kanalın bırakıldığına dair bir şey söylemek istenmediği zamanlarda
			fmt.Println("dropped")
		}
		time.Sleep(time.Microsecond * 50)
	}
}
