package main

import (
	"fmt"
	"sync"
	"time"
)

var NUM_GUESTS = 5
var NUM_FOOD_PER_GUEST = 3

type Philosopher struct {
	chair_index int
        food_pieces int
}

func (p *Philosopher) Eat() {
        p.food_pieces += 1
        fmt.Printf("EATING -- Portion (%d)", p.food_pieces)
        fmt.Println("-- Tasty leftover food residue on utensils, P", p.chair_index)

}

func (p *Philosopher) Think(b bool) {
        cur := p.chair_index

        if b {
                fmt.Println("THINKING -- Pontificating brownian motion, P", cur)
        } else {
                fmt.Println("ThInKiNg -- Pontificating brownian motion, P", cur)
        }
        time.Sleep(100)
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
	right := (p.chair_index + 1) % NUM_GUESTS

	for {
		// attempt to lock chopstick to left and if successful, lock chopstick to the right
		if t.chopsticks[left].TryLock() {
			if t.chopsticks[right].TryLock() {
                                p.Eat()
				t.chopsticks[right].Unlock()
				t.chopsticks[left].Unlock()

                                // if philosopher has finished all pieces of food, done eating!
				if p.food_pieces == NUM_FOOD_PER_GUEST  {
                                        wg.Done()
                                        return
                                }
			} else {
                                t.chopsticks[left].Unlock() // don't hold on to first lock if second lock taking unsuccessful
                                p.Think(true)
                        }
		} else {
                        p.Think(false)
                }
	}
}

func main() {
	var wg sync.WaitGroup
	var t Table
	var guests []Philosopher

	t.Setup()

	for i := 0; i < NUM_GUESTS; i++ {
		p := Philosopher{chair_index: i}
		guests = append(guests, p)
	}

	for i := 0; i < NUM_GUESTS; i++ {
		wg.Add(1)
		go t.Sit(&guests[i], &wg)
	}

	wg.Wait() // block until wg value down to 0
	fmt.Println("FINISHED - Dining Philosophers have finished eating their meals")
}
