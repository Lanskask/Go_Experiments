package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("Here")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("And There")
	wg.Wait()
	// time.Sleep(time.Second)
}
