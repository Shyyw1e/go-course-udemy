package concurency

import (
	"fmt"
	"time"
)

func ChannelRelease() {
	msg := make(chan string)

	time.Sleep(time.Second)
	go func() { msg <- "Ninja channel 12312123123" }()
	go func() { msg <- "Ninja channel 2" }()
	go func() { msg <- "Ninja channel 1" }()
	go func() { msg <- "Ninja channel" }()

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)

}
