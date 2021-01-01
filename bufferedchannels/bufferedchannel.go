package bufferedchannels

import (
	"fmt"
	"math/rand"
	"time"
)

// RandonNumber is a func.
func RandonNumber() {
	c := make(chan int, 100)

	for {
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Microsecond * 50)
	}
}
