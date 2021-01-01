package bufferedchannels

import (
	"fmt"
	"math/rand"
	"time"
)

// TimeoutMethod is a method.
func TimeoutMethod() {
	c := make(chan int)
	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Microsecond * 100):
			fmt.Println("time out", t)

		}
		time.Sleep(time.Millisecond * 50)
	}
}

// After is a method
func After(d time.Duration) chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		c <- true
	}()
	return c
}

/*
	İlk kullanılabilir kanal seçilir.
	Birden fazla kanal varsa, rastgele bir kanal seçilir.
	Hiçbir kanal yoksa, varsayılan durum yürütülür.
	Varsayılan durum yoksa select kilitlenir.
*/

// ForSelectMethod is a method.
func ForSelectMethod() {
	c := make(chan int)
	w := Worker{}

	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d go %d \n", w.id, data)
		case <-time.After(time.Microsecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}

// Worker is a struct
type Worker struct {
	id int
}

/* 	Notlar:

Kısa ömürlü bir kilit gerektiren basit bir örnek gördüğünüzde,
bir mutex veya okuma-yazma mutex’i kullannın.

*/
