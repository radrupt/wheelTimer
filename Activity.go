package example

// 签到定时提醒

import (
	"wheelTimer/wheelTimer"
	"time"
	"fmt"
)

type Activity struct {
	W *wheelTimer.WheelTimer
	remainingRounds int
}

func (task *Activity) Create(
	w *wheelTimer.WheelTimer,
) (*Activity, error) {
	return &Activity{
		W: w,
	}, nil
}

func (t *Activity) Run(){
	fmt.Println("excute activity", time.Now())
}

func (t *Activity) UpdateRemainingRounds(i int){
	t.remainingRounds = i
}

func (t *Activity) GetRemainingRounds() int {
	return t.remainingRounds
}