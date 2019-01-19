package main

import (
	"github.com/radrupt/wheelTimer"
	"sync"
)

func startLateMessage() {
	w := wheelTimer.New(1,256) // 创建定时任务
	// 添加并启动定时任务
	w.NewTimeOut(&Activity{Name: "A"}, 5)//5 seconds
	w.NewTimeOut(&Activity{Name: "B"}, 10)//10 seconds
	w.NewTimeOut(&Activity{Name: "C"}, 15)//15 seconds
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	startLateMessage()
	wg.Wait()
}