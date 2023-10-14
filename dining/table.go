package dining

import (
	"fmt"
	"sync"
	"time"
)

var NUM_TABLE_GUESTS = 5 // table only has 5 seats
var NUM_FOOD_PER_GUEST = 3 // 3 portions in the meal, salad, soup, and main course

type Table struct {
	chopsticks []sync.Mutex
}

func (t *Table) Setup() {
	t.chopsticks = make([]sync.Mutex, NUM_TABLE_GUESTS)
}

func (t *Table) Sit(p *Philosopher, sem *Semaphore, wg *sync.WaitGroup) {
	fmt.Println("New table arrival! P", p.id, p.chair_index)
        defer wg.Done()
        defer sem.Release()

	left := p.chair_index
	right := (p.chair_index + 1) % NUM_TABLE_GUESTS

	for {
		// attempt to lock chopstick to left and if successful, lock chopstick to the right
		if t.chopsticks[left].TryLock() {
			if t.chopsticks[right].TryLock() {
                                p.Eat()
				t.chopsticks[right].Unlock()
				t.chopsticks[left].Unlock()

                                // if philosopher has finished all pieces of food, done eating!
				if p.pieces_eaten == NUM_FOOD_PER_GUEST  { return } else { time.Sleep(20) }
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
