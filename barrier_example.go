package main

import "time"

func waitOnBarrier(name string, timeToSleep int, barrier *Barrier) {
	for {
		println(name, " is running")
		time.Sleep(time.Duration(timeToSleep) * time.Second)
		println(name, "is waiting")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go waitOnBarrier("red", 4, barrier)
	go waitOnBarrier("blue", 10, barrier)
	time.Sleep(time.Duration(100) * time.Second)
}
