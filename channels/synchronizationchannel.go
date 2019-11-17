package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- false
}

func workerone(done chan bool) {
	fmt.Print("workingone ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func workertwo(done chan bool) {
	fmt.Print("workingtwo ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	go workerone(done)
	go workertwo(done)
	<-done
}
