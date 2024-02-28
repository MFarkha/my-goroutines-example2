package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NUM_OF_PIZZA = 10

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NUM_OF_PIZZA {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++
		fmt.Printf("Making pizza #%d, It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay * int(time.Second)))

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** A cook quit while making the pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("The pizza #%d is ready", pizzaNumber)
		}
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var count = 0
	// run forever or quit as we receive notification
	// try to make pizza

	for {
		currentPizza := makePizza(count)
		count = currentPizza.pizzaNumber
		select {
		case pizzaMaker.data <- *currentPizza:
			//
		case quitChan := <-pizzaMaker.quit:
			close(pizzaMaker.data)
			close(quitChan)
			return
		}
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {
	//seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//print out a message
	color.Cyan("The Pizzeria is open for business!")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the backgroud
	go pizzeria(pizzaJob)

	//create and run a consumer
	for d := range pizzaJob.data {
		if d.pizzaNumber <= NUM_OF_PIZZA {
			if d.success {
				color.Green(d.message)
				color.Green("Order #%d is out for delivery", d.pizzaNumber)
			} else {
				color.Red(d.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("We are done for today")
			_ = pizzaJob.Close()
		}
	}
	//print out the ending message
	color.Cyan("We made %d pizzas, but failed to make %d, with attempts %d in total", pizzaMade, pizzaFailed, total)
}
