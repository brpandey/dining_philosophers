package main

import (
	"fmt"
	"sync"
	"time"
)

var NUM_GUESTS = 5

type Philosopher struct {
	chair_index int
}

type Table struct {
	chopsticks []sync.Mutex
}

func (t *Table) Setup() {
	t.chopsticks = make([]sync.Mutex, NUM_GUESTS)
}

func (t *Table) Sit(p *Philosopher, wg *sync.WaitGroup) {
	fmt.Println("New table arrival! P", p.chair_index)

	left := p.chair_index
	cur := p.chair_index
	right := (p.chair_index + 1) % NUM_GUESTS

	for {
		// attempt to lock chopstick to left and if successful, lock chopstick to the right
		if t.chopsticks[left].TryLock() {
			if t.chopsticks[right].TryLock() {
				fmt.Println("EATING -- Tasty food residue on chopsticks, P", cur)

				t.chopsticks[right].Unlock()

				wg.Done()

				t.chopsticks[left].Unlock()
				return
			}

			t.chopsticks[left].Unlock() // don't hold on to first lock if second lock taking unsuccessful

		} else {
			fmt.Println("THINKING -- Pontificating brownian motion, P", cur)
			time.Sleep(100)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var t Table
	var guests []Philosopher

	t.Setup()

	for i := 0; i < NUM_GUESTS; i++ {
		wg.Add(1)
		p := Philosopher{chair_index: i}
		guests = append(guests, p)
	}

	for i := 0; i < NUM_GUESTS; i++ {
		go t.Sit(&guests[i], &wg)
	}

	wg.Wait()
	fmt.Println("FINISHED - Dining Philosophers have finished eating their meals")
}
