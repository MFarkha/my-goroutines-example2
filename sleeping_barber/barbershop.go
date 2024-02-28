package main

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string // push names of clients into this channel
	Open            bool
}

func (b *Barbershop) Add(barber string) {
	b.NumberOfBarbers++
	go func() {
		isSleeping := true
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			if len(b.ClientsChan) == 0 {
				isSleeping = true
				color.Yellow("There is nothing to do: %s takes a nap", barber)
			}
			client, shopOpen := <-b.ClientsChan // check if channel is open == means the barbershop is closed
			if shopOpen {
				if isSleeping {
					isSleeping = false
					color.Yellow("`%s` wakes `%s` up", client, barber)
				}
				// cut hair
				b.CutHair(barber, client)
			} else {
				// shop closed
				b.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (b *Barbershop) CutHair(barber string, client string) {
	//
	color.Green("`%s` is cutting hair of `%s`", barber, client)
	time.Sleep(b.HairCutDuration)
	color.Green("`%s` finished hair cutting of `%s`", barber, client)
}

func (b *Barbershop) SendBarberHome(barber string) {
	//
	color.Cyan("`%s` sent home", barber)
	b.BarbersDoneChan <- true
}

func (b *Barbershop) CloseShopForDay() {
	color.Cyan("Closing the shop for day")
	b.Open = false
	close(b.ClientsChan)
	// waiting for barbers to finish
	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChan
	}
	close(b.BarbersDoneChan)
	color.Green("The barbershop closed for day!")
}

func (b *Barbershop) AddClient(client string) {
	color.Green("*** `%s` arrives", client)
	if b.Open {
		select {
		case b.ClientsChan <- client:
			color.Yellow("`%s` takes a seat in the waiting room", client)
		default:
			color.Red("The waiting room is full and `%s` leaves", client)
		}

	} else {
		color.Red("The shop is closed and `%s` has to leave", client)
	}
}
