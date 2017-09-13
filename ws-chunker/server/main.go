package main

import (
	"flag"
	"log"
	"net/http"

	"crypto/rand"
	"fmt"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

const chunkSize = 64 * 1024 // 64 KiB
var blob []byte

func download(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	chnk := make([]byte, 0, chunkSize)
	for currentByte := 0; currentByte < len(blob); currentByte += chunkSize {
		if currentByte+chunkSize > len(blob) {
			chnk = blob[currentByte:]
		} else {
			chnk = blob[currentByte : currentByte+chunkSize]
		}
		err = c.WriteMessage(websocket.BinaryMessage, chnk)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	blob = make([]byte, 128*1024*1024) // 128MiB
	rand.Read(blob)

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/file", download)
	fmt.Printf("Listening %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
