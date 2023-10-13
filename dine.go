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
        fmt.Println("EATING -- Portion (", p.food_pieces, ")-- Tasty leftover food residue on utensils, P", p.chair_index)
}

func (p *Philosopher) Think(single bool) {
        cur := p.chair_index

        if single {
                fmt.Println("THINKING -- Left utensil put down -- Que sera sera, P", cur)
        } else {
                fmt.Println("THINKING -- No utensils -- Pontificating brownian motion, P", cur)
        }
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
                                } else {
                                        time.Sleep(20)
                                }
			} else {
                                t.chopsticks[left].Unlock() // don't hold on to first lock if second lock taking unsuccessful
                                p.Think(true)
                                time.Sleep(50)
                        }
		} else {
                        p.Think(false)
                        time.Sleep(50)
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
