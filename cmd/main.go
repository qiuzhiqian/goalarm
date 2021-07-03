package main

import (
	"log"
	"time"

	"github.com/qiuzhiqian/goalarm"
)

type Manager struct {
	Monitor *goalarm.Monitor
	/*action  *Action*/
}

func (m *Manager) Start() {
	event := make(chan time.Time, 100)
	m.Monitor.AddEvent(event)

	go m.Monitor.Start()

	for t := range event {
		log.Println("done ", t)
	}

	m.Monitor.Done()
}

func main() {
	log.Println("go-alarm")

	var monitor goalarm.Monitor

	monitor.SetConfig(&goalarm.Config{
		Year:  []int{2020, 2021, 2022},
		Month: []int{2, 5, 7, 8, 11},
		Day:   []int{1, 2, 3, 15, 20, 28},
		Hour:  []int{10, 13, 17, 18, 21, 23},
		/*Minute: []int{10, 13, 17, 20, 28, 33, 37, 42, 46, 51, 56},*/
		Second: []int{23, 33, 45, 57},
	})

	m := Manager{
		Monitor: &monitor,
	}

	m.Start()
}
