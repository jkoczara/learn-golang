package main

import (
	"fmt"
	"time"
)

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool)

	//dones[0] = make(chan bool)
	go greet("1", done)

	//dones[1] = make(chan bool)
	go greet("2", done)

	//dones[2] = make(chan bool)
	go slowGreet("3", done)

	//dones[3] = make(chan bool)
	go greet("4", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for doneChan := range done {
		fmt.Println(doneChan)
	}
}

func greet(phrase string, doneCHan chan bool) {
	fmt.Println("Hello!", phrase)
	doneCHan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}
