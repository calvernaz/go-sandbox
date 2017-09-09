package main

import (
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go server()
	go client()

	time.AfterFunc(time.Duration(30)*time.Second, func() {
		wg.Done()
	})

	wg.Wait()
}
