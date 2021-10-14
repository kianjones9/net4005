package reservation

import (
	"errors"
	"fmt"
	"strconv"
)

type Reservation struct {
	Name, Class string
	Seat        int
	Sold        bool
}

func (r *Reservation) String() string {
	return fmt.Sprintf("%s %s %d", r.Name, r.Class, r.Seat)
}

// Flight struct represents a flight.
type Flight struct {
	database [30]Reservation // private
}

// List methods returns available seats within each passenger class
func (f *Flight) List(payload Reservation, reply *[]string) error {

	// LOOK AT ME TMR! Messed up zero indexing for ticket prices and availability

	//business class
	lower, upper := 0, 0
	threshold := 3
	sold := priceAvailTickets(f.database[0:5])

	if sold < threshold {
		lower, upper = threshold-sold, 5-threshold
	} else {
		lower, upper = 0, 5-sold
	}

	*reply = append(*reply, strconv.Itoa(lower)+" tickets available at $500")
	*reply = append(*reply, strconv.Itoa(upper)+" tickets available at $800")

	// economy class
	threshold = 10

	sold = priceAvailTickets(f.database[5:])
	if sold <= 10 {
		*reply = append(*reply, strconv.Itoa(10-sold)+" tickets available at $200")
		*reply = append(*reply, strconv.Itoa(10)+" tickets available at $300")
		*reply = append(*reply, strconv.Itoa(5)+" tickets available at $450")
	} else if sold > 10 && sold <= 20 {
		*reply = append(*reply, strconv.Itoa(0)+" tickets available at $200")
		*reply = append(*reply, strconv.Itoa(20-sold)+" tickets available at $300")
		*reply = append(*reply, strconv.Itoa(5)+" tickets available at $450")
	} else if sold > 20 {
		*reply = append(*reply, strconv.Itoa(0)+" tickets available at $200")
		*reply = append(*reply, strconv.Itoa(0)+" tickets available at $300")
		*reply = append(*reply, strconv.Itoa(25-sold)+" tickets available at $450")
	}

	return nil
}

func (f *Flight) Reserve(payload Reservation, reply *[]string) error {

	if payload.Seat < 1 || payload.Seat > 30 {
		*reply = append(*reply, "Failed to reserve seat")
		return errors.New("invalid seat selection")
	}

	if payload.Seat <= 5 {
		if payload.Class != "business" {
			*reply = append(*reply, "Failed to reserve seat")
			return errors.New("invalid seat selection")
		}
	}

	if payload.Seat > 5 {
		if payload.Class != "economy" {
			*reply = append(*reply, "Failed to reserve seat")
			return errors.New("invalid seat selection")
		}
	}

	if f.database[payload.Seat-1].Sold {
		*reply = append(*reply, "Failed to reserve seat")
		return errors.New("seat was already sold")
	} else {
		f.database[payload.Seat-1] = payload
		f.database[payload.Seat-1].Sold = true
		*reply = append(*reply, "Successful reservation: ")
		*reply = append(*reply, payload.String())
	}

	return nil
}

func (f *Flight) Passengerlist(payload Reservation, reply *[]string) error {

	for i := range f.database {

		if f.database[i].Sold {
			*reply = append(*reply, f.database[i].String())
		}
	}

	return nil
}

func priceAvailTickets(subDB []Reservation) int {

	var sold = 0

	for i := range subDB {

		if subDB[i].Sold {
			sold++
		}
	}

	return sold

}

// NewFlight function returns a new instance of Flight (pointer).
func NewFlight() *Flight {
	f := Flight{
		database: [30]Reservation{},
	}

	return &f
}
