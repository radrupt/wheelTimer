package main

import (
	"github.com/radrupt/wheelTimer"
	"sync"
)

func startLateMessage() {
	// 启动延迟消息服务
	w := wheelTimer.New(1,256) // 创建定时任务
	Start(w) // 开启消费者
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go startLateMessage()
	wg.Wait()
}