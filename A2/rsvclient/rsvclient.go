package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strconv"

	"net4005/A2/reservation"
)

func main() {

	server := net.JoinHostPort(os.Args[2], os.Args[3])

	// get RPC client by dialing TCP connection
	client, _ := rpc.Dial("tcp", server)

	var reply = []string{}

	if len(os.Args) == 4 {

		// List RPC
		if os.Args[1] == "list" {

			err := client.Call("Flight.List", reservation.Reservation{}, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				for i := range reply {
					fmt.Println(reply[i])
				}
			}

			// Passengerlist RPC
		} else if os.Args[1] == "passengerlist" {
			// client logic
			err := client.Call("Flight.Passengerlist", reservation.Reservation{}, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				for i := range reply {
					fmt.Println(reply[i])
				}
			}

			// otherwise fail
		} else {
			fmt.Println("Usage: client <RPC> <server_IP> <port> <passenger_class> <passenger_name> <seat_number>")
			os.Exit(1)
		}
	} else if len(os.Args) == 7 {

		class := os.Args[4]
		name := os.Args[5]

		seat, err := strconv.Atoi(os.Args[6])
		if err != nil {
			fmt.Println(err)
			fmt.Println("Seat number is invalid format.\nUsage: client <RPC> <server_IP> <port> <passenger_class> <passenger_name> <seat_number>")
			os.Exit(1)
		}

		var res = reservation.Reservation{Name: name, Class: class, Seat: seat}

		err = client.Call("Flight.Reserve", res, &reply)

		for i := range reply {
			fmt.Println(reply[i])
		}

		if err != nil {
			fmt.Println("Error: ", err)
		}

	} else {
		fmt.Println("Usage: client <RPC> <server_IP> <port> <passenger_class> <passenger_name> <seat_number>")
		os.Exit(1)
	}

}
