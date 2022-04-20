package initialize

import "fmt"

var waitMonkeyJobCh chan interface{}

const (
	count = 2
)

// 启动goroutine 获取ch里面任务
func RegisterMonkeyWorker() {

	// 读取mysql里面jobs,fasong dao ch.

	for i := 0; i < count; i++ {
		go func() {
			for {
				job := <-waitMonkeyJobCh
				fmt.Println(job)
			}

		}()
	}
}
