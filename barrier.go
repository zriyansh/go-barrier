package main

import "sync"

type Barrier struct {
	total int         // total threads participating in barrier
	count int         // count of thread that can call wait()
	mutex *sync.Mutex // we are accesing variable inside our struct from multiple threads
	cond  *sync.Cond  // so that we block the thread when they are waiting on the other threads to call the wait operation and wake it up again
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}           // pointer to a mutex
	condToUse := sync.NewCond(lockToUse) // return a ptr to cond variable
	return &Barrier{size, size, lockToUse, condToUse}
}

func (b *Barrier) Wait() {
	b.mutex.Lock() // so that we safely count the count variable
	b.count -= 1   // 1 less thread waiting on this therad
	if b.count == 0 {
		b.count = b.total  // all thread have called wait()
		b.cond.Broadcast() // so that all thread wake up and continue execution
	} else {
		b.cond.Wait() // block a thread from executing
	}
	b.mutex.Unlock()
}

// full implemenation of barrier
