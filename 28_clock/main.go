// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// A TCP server that periodically writes the time, updated as an exercise to provide time in different time zones when started on different ports

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var portFlag = flag.String("p", "8000", "specify a port")

var portMap = map[string]string{
	"8010": "US/Eastern",
	"8020": "Asia/Tokyo",
	"8030": "Europe/London",
}

func main() {
	flag.Parse()

	timeZone := portMap[*portFlag]
	listenAddress := fmt.Sprintf("localhost:%v", *portFlag)
	listener, err := net.Listen("tcp", listenAddress)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port %v", *portFlag)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		log.Print("Connection accepted")
		go handleConn(conn, timeZone)
	}
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()

	for {
		loc, _ := time.LoadLocation(timeZone)

		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}
