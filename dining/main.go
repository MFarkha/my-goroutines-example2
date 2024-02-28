// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilemma.

// Philosopher is a struct which stores information about a philosopher.
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	Name      string
	RightFork int
	LeftFork  int
}

var philosophers = []Philosopher{
	{Name: "Plato", LeftFork: 4, RightFork: 0},
	{Name: "Socrates", LeftFork: 0, RightFork: 1},
	{Name: "Aristotel", LeftFork: 1, RightFork: 2},
	{Name: "Pascal", LeftFork: 2, RightFork: 3},
	{Name: "Locke", LeftFork: 3, RightFork: 4},
}

var hunger = 3 // how many times a philosopher eats
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

type TableOrder struct {
	Names []string
	mutex *sync.Mutex
}

var tableOrder = TableOrder{
	Names: make([]string, 0, len(philosophers)),
	mutex: &sync.Mutex{},
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers)) // zero means everyone is done eating

	seated := &sync.WaitGroup{} // zero means everyone is seated
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()
}

func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	// seat philosopher at the table
	fmt.Printf("The philosopher %s seated at the table\n", p.Name)
	seated.Done()

	seated.Wait() // make sure all philosopers are the table
	//eat three times
	for i := hunger; i > 0; i-- {
		// solving a logical race condition by introducing an order of fork selection
		if p.LeftFork > p.RightFork {
			forks[p.RightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", p.Name)
			forks[p.LeftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", p.Name)
		} else {
			forks[p.LeftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", p.Name)
			forks[p.RightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", p.Name)
		}
		fmt.Printf("\t%s takes the both forks and is eating\n", p.Name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking\n", p.Name)
		time.Sleep(thinkTime)

		forks[p.LeftFork].Unlock()
		forks[p.RightFork].Unlock()
		fmt.Printf("\t%s released the both forks\n", p.Name)
	}
	tableOrder.mutex.Lock()
	tableOrder.Names = append(tableOrder.Names, p.Name)
	tableOrder.mutex.Unlock()
	fmt.Printf("The philosopher:`%s` left the table \n", p.Name)
}

func main() {
	thinkTime = 0 * time.Second
	eatTime = 0 * time.Second
	// print out welcome message
	fmt.Println("Dining Philosopher's problem")
	fmt.Println("The table is empty")

	//start the meal
	dine()

	//print out finished message
	fmt.Println("The table is empty")
	// time.Sleep(sleepTime)
	fmt.Printf("Table Order: %s.\n", strings.Join(tableOrder.Names, ", "))
}
