package main

import (
	"log"
	"net"
)

func server() {
	laddr, err := net.ResolveUDPAddr("udp", ":5555")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	go handleConn(conn)
}

func handleConn(conn *net.UDPConn) {
	buf := make([]byte, 64*1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Message from client: %v\n", string(buf[:n]))
		conn.WriteToUDP(buf[:n], addr)
	}
}
