package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	number                        int
	leftChopstick, rightChopstick *ChopStick
}

type Host struct {
	eatingPhilosophers map[int]Philosopher
}

func (philosopher Philosopher) eat() {
	philosopher.leftChopstick.Lock()
	philosopher.rightChopstick.Lock()

	fmt.Printf("starting to eat %v\n", philosopher.number)
	fmt.Printf("finished eating %v\n", philosopher.number)

	philosopher.leftChopstick.Unlock()
	philosopher.rightChopstick.Unlock()
}

func (host Host) requestEatPermission(philosopher Philosopher) bool {
	if len(host.eatingPhilosophers) >= 2 {
		return false
	}

	var _, hasPhilosopher = host.eatingPhilosophers[philosopher.number]

	if hasPhilosopher {
		return false
	} else {
		host.eatingPhilosophers[philosopher.number] = philosopher
		return true
	}
}

func (host Host) releaseEatingPhilosopher(philosopher Philosopher) {
	delete(host.eatingPhilosophers, philosopher.number)
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

	// each philosopher should only eat three times
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
}
