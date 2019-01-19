package wheelTimer

import (
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
type WheelTimer struct {
	tickDuration int // 每转一格的时间， 单位s
	ticksPerWheel int // 轮盘的长度
	mask         int // 掩码
	WheelBucket  []Bucket
	// 状态
	// "WORKER_STATE_INIT"
	// "WORKER_STATE_STARTED"
	// "WORKER_STATE_SHUTDOWN"
	state        string
	tick       int //当前轮盘转动的格数
	taskid     int64
	preTime int64
}

func New(tickDuration int, ticksPerWheel int) *WheelTimer {
	wheelBucket := createWheel(ticksPerWheel)
	return &WheelTimer {
		tickDuration: tickDuration,
		ticksPerWheel: ticksPerWheel,
		mask: ticksPerWheel - 1,
		WheelBucket: wheelBucket,
		state: "WORKER_STATE_INIT",
		tick: 0,
		taskid: 0,
		preTime: time.Now().UnixNano(),
	}
}

func createWheel(ticksPerWheel int) []Bucket {
	wheelBucket := make([]Bucket, ticksPerWheel)
	for index := 0; index < ticksPerWheel; index++ {
		wheelBucket[index] = Bucket{make(map[int64]Task)}
	}
	return wheelBucket
}

func (w *WheelTimer) Start() {
	if w.state == "WORKER_STATE_STARTED" {
		return
	}
	w.state = "WORKER_STATE_STARTED"
	// 开启goroutine
	// goroutine 里判断State如果为WORKER_STATE_SHUTDOWN,则done
	wg.Add(1)
	go turn(w)
}

func (w *WheelTimer) Stop() {
	w.state = "WORKER_STATE_SHUTDOWN"
}

func (w *WheelTimer) Shutdown() {
	wg.Wait()
} 

// 添加定时任务
// deadline: 单位s
func (w *WheelTimer) NewTimeOut(task Task, timeInterval int) {
	w.Start()
	idx := (timeInterval / w.tickDuration + w.tick) & w.mask
	remainingRounds := (timeInterval / w.tickDuration - w.tick) / w.ticksPerWheel
	if remainingRounds < 0 {
		remainingRounds = 0
	}
	task.UpdateRemainingRounds(remainingRounds)
	w.WheelBucket[idx].Tasks[w.taskid] = task
	atomic.AddInt64(&w.taskid,1)
}

func turn(w *WheelTimer) {
	defer wg.Done()
	var sec int
	for {
		if w.state == "WORKER_STATE_SHUTDOWN" {
			break
		}
		idx := w.tick & w.mask
		w.WheelBucket[idx].Expire()
		nextTime := w.preTime + 1000000000
		now := time.Now()
		w.preTime = nextTime
		time.Sleep(time.Duration(nextTime - now.UnixNano()) * time.Nanosecond)
		sec += 1
		w.tick = (w.tick + 1) & w.mask
	}
}
