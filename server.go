package main

import "net"
import "fmt"
import "bufio"
import (
	"strings"
	"math/rand"
	"strconv"
) // only needed below for sample processing

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":6969")
	ln.Addr().Network()

	// accept connection on port
	conn, _ := ln.Accept()
	rand := strconv.Itoa(rand.Intn(1000000000))

	// demander authentification
	fmt.Print(rand + "\n")
	// run loop forever (or until ctrl-c)
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message), "from ")
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))

}