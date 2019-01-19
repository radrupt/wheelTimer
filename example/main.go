package main

import (
	"github.com/radrupt/wheelTimer"
	"sync"
)

func startLateMessage() {
	// 启动延迟消息服务
	w := wheelTimer.New(1,256) // 创建定时任务
	
	// 添加定时任务
	newActivity := Activity{}
	w.NewTimeOut(newActivity.Create(w, "A"), 5)
	w.NewTimeOut(newActivity.Create(w, "B"), 10)
	w.NewTimeOut(newActivity.Create(w, "C"), 15)
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go startLateMessage()
	wg.Wait()
}