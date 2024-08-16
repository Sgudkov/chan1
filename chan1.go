package main

import (
	"fmt"
	"hash/fnv"
	"time"
)

func main() {

	chan1 := make(chan uint32, 1)

	//Run in a goroutine to calculate the hash
	//and write to the channel
	go runchannels(chan1)

	//Read from the channel and print the results
	go getNumber(chan1)

	time.Sleep(2 * time.Second)

}

func runchannels(in chan<- uint32) {
	for i := 0; i < 500; i++ {
		val := i
		in <- hash(fmt.Sprint(val))
		fmt.Println("Write to chan2", val)
	}

	close(in)
}

func getNumber(in <-chan uint32) {
	for val := range in {
		fmt.Println("get from chan1", val)
	}

}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
