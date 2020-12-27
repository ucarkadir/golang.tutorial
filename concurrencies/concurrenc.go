package concurrencies

import (
	"fmt"
	"sync"
	"time"
)

/*
	* Concurrency : Eşzamanlılık

	1. Goroutines: Go Rutinleri
	2. Channels : Kanallar

	Eşzamanlılık dostu bir dil olarak tanımlanır. Bunun nedeni, iki
	güçlü mekanizma için basit bir sözdizimi sağlamasıdır:

	* go rutinleri
	* kanallar


*/

// Goroutine is
func Goroutine() {
	fmt.Println("Start")
	go process() // goroutine go anantar cümleciği ile başlatılır
	time.Sleep(time.Microsecond * 10)
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}

/////////////////////////////////////////////////////////

var counter = 0

// GoRoutineForLoop is a for loop gorutine
func GoRoutineForLoop() {
	for i := 0; i < 20; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	counter++
	fmt.Println(counter)
}

/////////////////////////////////////////////////////////

var (
	counterMutex = 0
	lock         sync.Mutex
)

// GoRoutineMutex is a Mutex go routine
func GoRoutineMutex() {
	for i := 0; i < 20; i++ {
		go incrMutex()
	}
}

func incrMutex() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
