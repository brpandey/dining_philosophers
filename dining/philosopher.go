package dining

import (
        "fmt"
)

type Philosopher struct {
        id int
	chair_index int
        pieces_eaten int
}

func NewPhilosopher(uid int, pos int) Philosopher {
        return Philosopher{id: uid, chair_index: pos}
}

func (p *Philosopher) Eat() {
        p.pieces_eaten += 1
        fmt.Println("EATING -- Portion (", p.pieces_eaten, ")-- Tasty leftover food residue on utensils, P", p.id, p.chair_index)
}

func (p *Philosopher) Think(single bool) {
        pos := p.chair_index

        var text string

        if single {
                text = "THINKING -- Left utensil put down -- Que sera sera, P"
        } else {
                text = "THINKING -- No utensils -- Pontificating brownian motion, P"
        }

        fmt.Println(text, p.id, pos)
}
