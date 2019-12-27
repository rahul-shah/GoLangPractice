package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	//This will block
	//Trying to read from channel without writing to it
	/*<-channel
	fmt.Println("Here")*/

	go func() {
		channel <- 353
	}()

	val := <-channel
	fmt.Printf("Got Value %d\n", val)

	//Sending values
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			channel <- i
			time.Sleep(time.Second)
		}
		//Signal that sending is completed by closing the channel
		close(channel)
	}()

	//Loop over channel values
	for j := range channel {
		fmt.Printf("Read %d\n", j)
	}
}
