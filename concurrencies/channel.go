package concurrencies

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker is a struct
type Worker struct {
	id int
}

// MakeChannel is
func MakeChannel() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

/*
	Kanal

	1. Kullanımı:
		func worker(c chan int) { ... }

	2. Veri Gönderme:
		CHANNEL <- DATA

	3. Veri Alma:
		VAR := <-CHANNEL

	4. Go Rutinin Yürütülmesi
		Kanaldan Veri Alırken: veri bulunana kadar devam etmez
		Kanala Veri Gönderdiğimizde; veri alınana kadar devam etmez

*/
