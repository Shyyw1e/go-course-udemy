package concurency

import (
	"fmt"
	//"time"
)

func BufferChanel() {
	msg := make(chan string, 5)

	msg <- "Buffer Channel"
	msg <- "Buffer Channel 2"
	msg <- "Buffer Channel 3"
	msg <- "Buffer Channel 4"
	msg <- "Buffer Channel 5"
	close(msg)

	// 	for {
	// 		value, ok := <-msg
	// 		if !ok {
	// 			fmt.Println("Channel closed")
	// 			break
	// 		}
	// 		fmt.Println(value)
	// 	}

	for m := range msg {
		fmt.Println(m)
	}

}
