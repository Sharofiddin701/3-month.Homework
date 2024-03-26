package main

import (
	"fmt"
	"sync"
)

var counterGlobal = 0
var mutex sync.Mutex

func inctement1(wg *sync.WaitGroup) {
	mutex.Lock()
	counterGlobal++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	expectedCounter := 10

	for i := 0; i < expectedCounter; i++ {
		wg.Add(1)
		go inctement1(&wg)
	}

	wg.Wait()
	fmt.Println("Expected Counter:",
		expectedCounter, "Actual Cpunter:", counterGlobal)

	if expectedCounter != counterGlobal {
		fmt.Println("Race condition detected")
	} else {
		fmt.Println("No race condition detected.")
	}

}
