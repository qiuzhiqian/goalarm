package main

import (
	"log"
	"time"

	"github.com/qiuzhiqian/goalarm"
)

type Manager struct {
	Monitor *goalarm.Monitor
	Actions []Action
}

func (m *Manager) Start() {
	event := make(chan time.Time, 100)
	m.Monitor.AddEvent(event)

	go m.Monitor.Start()

	for t := range event {
		log.Println("done ", t)

		//do all action
		for _, action := range m.Actions {
			err := action.DoAction()
			if err != nil {
				return
			}
		}
	}

	m.Monitor.Done()
}

func main() {
	log.Println("go-alarm")

	var monitor goalarm.Monitor

	monitor.SetConfig(&goalarm.Config{
		Year:  []int{2020, 2021, 2022},
		Month: []int{2, 5, 7, 8, 11},
		Day:   []int{1, 2, 3, 4, 15, 20, 28},
		Hour:  []int{9, 10, 13, 17, 18, 21, 23},
		/*Minute: []int{10, 13, 17, 20, 28, 33, 37, 42, 46, 51, 56},*/
		Second: []int{23, 33, 45, 57},
	})

	m := Manager{
		Monitor: &monitor,
	}

	message1 := NewTextMessage("hello world", nil, nil)
	action1 := NewRebot("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx", message1)

	message2 := NewTextMessage("hello world", []string{"@all"}, nil)
	action2 := NewRebot("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx", message2)

	message3 := NewMarkdownMessage(`#hello world
**body**
> haha
` + "`code`")
	action3 := NewRebot("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx", message3)
	m.Actions = []Action{action1, action2, action3}

	m.Start()
}
