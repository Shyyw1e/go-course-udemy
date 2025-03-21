package main

import (
	// "fmt"
	// "sync"
	// "time"

	"go_course/concurency"
	"go_course/concurency/gpt_tasks"
)

func main() {
	concurency.WaitGroups()
	concurency.ChannelRelease()
	concurency.BufferChanel()
	concurency.SomeChannels()
	concurency.RaceCondition()
	gpt_tasks.Task1()
}
