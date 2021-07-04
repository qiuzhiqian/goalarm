# goalarm

这是一个使用golang编写的时间定时处理库。会通过配置来查找下一个需要触发的时间点。

# 使用
```go
import github.com/qiuzhiqian/goalarm

func main() {
    var monitor goalarm.Monitor
    monitor.SetConfig(&goalarm.Config{
		Month:  []int{2, 5, 7, 8, 11},
		Day:    []int{1, 2, 3, 4, 15, 20, 28},
		Hour:   []int{9, 10, 13, 17, 18, 21, 23},
		Second: []int{23, 33, 45, 57},
	})

    event := make(chan time.Time, 100)
	monitor.AddEvent(event)

	go monitor.Start()

	for t := range event {
		log.Println("done ", t)
	}

	monitor.Done()
}
```