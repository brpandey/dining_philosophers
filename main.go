package main

import (
        "sync"
	"fmt"
        "github.com/brpandey/dining_philosophers/dining"
)

func main() {
        var NUM_TRAVELING = 12 // wandering band of traveling philosophers
	var wg sync.WaitGroup
	var t dining.Table
	var traveling_guests []dining.Philosopher

	t.Setup()
        sem := dining.NewSemaphore(dining.NUM_TABLE_GUESTS)

	for i := 0; i < NUM_TRAVELING; i++ {
                p := dining.NewPhilosopher(i, i % dining.NUM_TABLE_GUESTS)
		traveling_guests = append(traveling_guests, p)
	}

	for i := 0; i < NUM_TRAVELING; i++ {
                // Bound the concurrency allowing only for NUM_TABLE_GUESTS at a time
                // among the traveling philosophers
                sem.Acquire() 
		wg.Add(1)
		go t.Sit(&traveling_guests[i], sem, &wg)
	}

	wg.Wait() // block until wg value down to 0
	fmt.Println("FINISHED - Dining Philosophers have finished eating their meals")
}
