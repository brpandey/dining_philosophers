package dining

type Semaphore struct {
        ch chan int
}

func NewSemaphore(count int) *Semaphore {
        var s Semaphore
        s.ch = make(chan int, count)
        return &s
}

func (s *Semaphore) Acquire() {
        x := 1
        s.ch <- x // send onto buffered channel which will block if full
}

func (s *Semaphore) Release() {
        <-s.ch // receive expression, discarding result, will block if empty
}
