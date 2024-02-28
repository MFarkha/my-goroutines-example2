// a purpose: how much many someone is going to make in 52 week
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// bank balance
	var bankBalance int
	var balance sync.Mutex
	// print out starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)
	// define weekly revenue
	incomes := []Income{
		{Source: "main job", Amount: 500},
		{Source: "gifts", Amount: 10},
		{Source: "part-time job", Amount: 50},
		{Source: "investments", Amount: 100},
	}
	// loop through 52 weeks and print out how to much he made
	for i, income := range incomes {
		wg.Add(1)
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; i < 52; i++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	// print out final balance
	fmt.Printf("Final account balance is $%d.00\n", bankBalance)
}
