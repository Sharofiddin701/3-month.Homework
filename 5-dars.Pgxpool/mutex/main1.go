package main

import (
	"fmt"
	"log"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (acc *BankAccount) deposit(amount int, wg *sync.WaitGroup) {
	defer acc.mutex.Unlock()

	acc.mutex.Lock()
	acc.balance += amount
	wg.Done()
}

func (acc *BankAccount) withdraw(amount int, wg *sync.WaitGroup) {
	defer acc.mutex.Unlock()

	acc.mutex.Lock()
	if acc.balance < amount {
		log.Println("Insufficient funds")
	}
	acc.balance -= amount
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	acc := &BankAccount{balance: 500}

	expected := 9
	for i := 0; i < expected; i++ {
		wg.Add(2)
		go acc.deposit(100, &wg)
		go acc.withdraw(150, &wg)
	}

	wg.Wait()
	fmt.Println("Expected:", expected)
	fmt.Println("Actual balance:", acc.balance)
}
