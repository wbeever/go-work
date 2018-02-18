package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanUp() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in Cleanup:", r)
	}
}

func say(s string) {
	defer cleanUp()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dead, a 2")
		}
	}
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}
