package main

import (
	"fmt"
	"log"
	"time"
)

type Monitor struct {
	config Config
	events []chan time.Time
}

func (m *Monitor) AddEvent(event chan time.Time) {
	m.events = append(m.events, event)
}

func (m *Monitor) Start() {
	quilt := make(chan bool, 1)
	for {
		t, err := NextTime(m.config)
		if err != nil {
			log.Println("find err:", err)
			return
		}

		log.Println("next:", t)

		timer := time.NewTimer(time.Until(t))

		select {
		case <-timer.C:
			now := time.Now()
			for _, e := range m.events {
				e <- now
			}
		case <-quilt:
			log.Println("quilt")
			return
		}
	}
}

type Manager struct {
	monitor *Monitor
	/*action  *Action*/
}

func (m *Manager) Start() {
	event := make(chan time.Time, 100)
	m.monitor.AddEvent(event)

	go m.monitor.Start()

	for t := range event {
		log.Println("done ", t)
	}
}

type Config struct {
	Year   []int
	Month  []int
	Day    []int
	Hour   []int
	Minute []int
	Second []int
}

func NextTime(c Config) (time.Time, error) {
	now := time.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	if len(c.Year) == 0 {
		c.Year = append(c.Year, year)
	}

	if len(c.Month) == 0 {
		c.Month = append(c.Month, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	}

	if len(c.Day) == 0 {
		for i := 0; i < 31; i++ {
			c.Day = append(c.Day, i+1)
		}
	}

	if len(c.Hour) == 0 {
		for i := 0; i < 24; i++ {
			c.Hour = append(c.Hour, i)
		}
	}

	if len(c.Minute) == 0 {
		for i := 0; i < 60; i++ {
			c.Minute = append(c.Minute, i)
		}
	}

	if len(c.Second) == 0 {
		for i := 0; i < 60; i++ {
			c.Second = append(c.Second, i)
		}
	}

	yearTag := c.Year[0]
	monthTag := c.Month[0]
	dayTag := c.Day[0]
	hourTag := c.Hour[0]
	minuteTag := c.Minute[0]
	secondTag := c.Second[0]

	for _, i := range c.Year {
		yearTag = i
		if year < i {
			//log.Println("year:", year, i)
			tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

			if tagTime.After(now) {
				//log.Println("after:", tagTime)
				return tagTime, nil
			}
		}
		for _, j := range c.Month {
			monthTag = j
			if month < j {
				//log.Println("month:", month, j)
				tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

				if tagTime.After(now) {
					//log.Println("after:", tagTime)
					return tagTime, nil
				}
			}

			for _, k := range c.Day {
				dayTag = k
				if day < k {
					//log.Println("day:", day, k)
					tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

					if tagTime.After(now) {
						//log.Println("after:", tagTime)
						return tagTime, nil
					}
				}

				for _, l := range c.Hour {
					hourTag = l
					if hour < l {
						//log.Println("hour:", hour, l)
						tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

						if tagTime.After(now) {
							//log.Println("after:", tagTime)
							return tagTime, nil
						}
					}

					for _, m := range c.Minute {
						minuteTag = m
						if minute < m {
							//log.Println("minute:", minute, m)
							tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

							if tagTime.After(now) {
								//log.Println("after:", tagTime)
								return tagTime, nil
							}
						}

						for _, n := range c.Second {
							secondTag = n
							if second < n {
								//log.Println("second:", second, n)
								tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

								if tagTime.After(now) {
									//log.Println("after:", tagTime)
									return tagTime, nil
								}
							}

							tagTime := time.Date(yearTag, time.Month(monthTag), dayTag, hourTag, minuteTag, secondTag, 0, time.Local)

							if tagTime.After(now) {
								//log.Println("after:", tagTime)
								return tagTime, nil
							}
						}
						secondTag = c.Second[0]
					}
					minuteTag = c.Minute[0]
				}
				hourTag = c.Hour[0]
			}
			dayTag = c.Day[0]
		}
		monthTag = c.Month[0]
	}

	return now, fmt.Errorf("xxxx")
}

func main() {
	log.Println("go-alarm")

	monitor := Monitor{
		config: Config{
			Year:  []int{2020, 2021, 2022},
			Month: []int{2, 5, 7, 8, 11},
			Day:   []int{1, 2, 3, 15, 20, 28},
			Hour:  []int{10, 13, 17, 18, 21, 23},
			/*Minute: []int{10, 13, 17, 20, 28, 33, 37, 42, 46, 51, 56},*/
			Second: []int{23, 33, 45, 57},
		},
	}

	m := Manager{
		monitor: &monitor,
	}

	m.Start()
}
