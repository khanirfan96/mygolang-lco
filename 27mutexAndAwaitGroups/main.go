package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - Google.com")

	waitg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var score = []int{0}

	waitg.Add(3)

	go func(waitg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One GoRoutine")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		waitg.Done()
	}(waitg, mutex)
	// waitg.Add(1)
	go func(waitg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two GoRoutine")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		waitg.Done()
	}(waitg, mutex)

	go func(waitg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three GoRoutine")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		waitg.Done()
	}(waitg, mutex)

	waitg.Wait()

	fmt.Println(score)
}
