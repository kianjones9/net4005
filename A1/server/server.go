package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var reqs = 0
var goodReqs = 0

func main() {

	listener, err := net.Listen("tcp", "localhost:5999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}
		reqs++
		go handleReq(con)
	}
}

func handleReq(con net.Conn) {

	filename := make([]byte, 64)

	_, err := con.Read(filename)

	if err != nil {
		log.Fatal(err)
	}

	// Strip null values from recieved filename
	sanitizedFilename := string(bytes.Trim(filename, "\x00"))

	var sb strings.Builder
	sb.WriteString("File ")
	sb.WriteString(sanitizedFilename)

	// Open file with file name read in from TCP connection
	file, err := os.Open(sanitizedFilename)

	// Check if file exists
	if err != nil {
		if os.IsNotExist(err) {
			sb.WriteString(" Not Found")
			con.Write([]byte(sb.String()))
			con.Close()
			return
		}
		log.Fatal(err)
	}
	defer file.Close()

	sb.WriteString(" Found\nServer Handled ")
	sb.WriteString(strconv.Itoa(reqs))
	sb.WriteString(" requests, ")
	sb.WriteString(strconv.Itoa(goodReqs))
	sb.WriteString(" Were Successful")

	bytes := []byte{byte(len([]byte(sb.String())))}

	fmt.Println(bytes)
	fmt.Println(sb.String())

	con.Write(bytes)
	con.Write([]byte(sb.String()))

	_, err = io.Copy(con, file)
	if err != nil {
		log.Fatal(err)
	}
	goodReqs++
	con.Close()
}
