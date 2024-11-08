package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	number                        int
	leftChopstick, rightChopstick *ChopStick
}

var host = make(chan int, 2)

var wg sync.WaitGroup

func (philosopher Philosopher) eat(c chan int) {
	c <- philosopher.number

	philosopher.leftChopstick.Lock()
	philosopher.rightChopstick.Lock()

	fmt.Printf("starting to eat %v\n", philosopher.number)

	time.Sleep(500 * time.Millisecond)

	fmt.Printf("finished eating %v\n", philosopher.number)

	philosopher.leftChopstick.Unlock()
	philosopher.rightChopstick.Unlock()

	<-c
	wg.Done()
}

func main() {
	chopSticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i + 1,
			chopSticks[i],
			chopSticks[(i+1)%5]}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			wg.Add(1)
			go philosophers[j].eat(host)
		}
	}

	wg.Wait()
}
