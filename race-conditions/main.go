package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func raceCondition(totalGoroutines int) {
	totalValue := 0

	var wg sync.WaitGroup
	wg.Add(totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {
		go func() {
			v := totalValue
			time.Sleep(1)
			v++
			totalValue = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("[race condition] - Total value:", totalValue)
}

func mutex(totalGoroutines int) {
	totalValue := 0

	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			mutex.Lock()
			v := totalValue
			time.Sleep(1)
			v++
			totalValue = v
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("[mutex] - Total value:", totalValue)
}

func atomicRaceCondition(totalGoroutines int) {
	var totalValue int32
	totalValue = 0

	var wg sync.WaitGroup
	wg.Add(totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {
		go func() {
			atomic.AddInt32(&totalValue, 1)
			time.Sleep(1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("[atomic race condition] - Total value:", totalValue)

}

func loop(t int, ch chan<- int) {
	for i := 0; i < t; i++ {
		ch <- i
	}
	close(ch) // avoid all goroutines are asleep - deadlock!
}

func printLoop(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	totalCPUs := runtime.NumCPU()
	fmt.Println("CPUs:", totalCPUs)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// raceCondition(totalCPUs)
	// mutex(totalCPUs)
	// atomicRaceCondition(totalCPUs)

	ch := make(chan int)
	go loop(10, ch)
	printLoop(ch)
}
