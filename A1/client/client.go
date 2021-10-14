package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: client <server_IP> <port> <filename>")
		os.Exit(1)
	}

	filename := os.Args[3]
	server := net.JoinHostPort(os.Args[1], os.Args[2])

	con, err := net.Dial("tcp", server)

	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	_, err = con.Write([]byte(filename))

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(string(filename))

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(con)

	// Get number of bytes
	numBytes := make([]byte, 1)

	_, err = reader.Read(numBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Preliminary message (number of bytes to process before first chunk of file, file found or not, statistics)
	prelim := make([]byte, numBytes[0])

	_, err = reader.Read(prelim)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(prelim))

	fmt.Printf("Downloading file: %s\n", filename)

	// Read file data
	fileBuffer := make([]byte, 1024)
	for {
		_, err = reader.Read(fileBuffer)
		if err == io.EOF {
			fmt.Println("Download complete")
			con.Close()
			return
		} else if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(fileBuffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}
