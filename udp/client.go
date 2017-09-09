package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func client() {
	raddr, err := net.ResolveUDPAddr("udp", ":5555")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := 0
	buf := make([]byte, 64*1024)
	for {
		time.Sleep(time.Duration(2) * time.Second)

		_, err = conn.Write([]byte(fmt.Sprintf("Message %d", c)))
		if err != nil {
			log.Fatal(err)
		}

		n, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Fatalf("Error reading from server %v", err)
		}
		log.Printf("Reply from server: %v\n", string(buf[:n]))
		c++
	}

}
