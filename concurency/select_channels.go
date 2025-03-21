package concurency

import (
	"fmt"
	"time"
)

func SomeChannels() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Channel 1. Passed 200 ms"
			time.Sleep(time.Millisecond * 200)

		}
	}()
	go func() {
		for {
			message2 <- "Channel 2. Passed 1 s"
			time.Sleep(time.Second)

		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)

		}
	}
}
