package example

import (
	"github.com/radrupt/wheelTimer"
)
type listenRedisTask struct {
	w *wheelTimer.WheelTimer
}
//开启定时任务
func startTimerTask(
	w *wheelTimer.WheelTimer,
) {
	newActivity := Activity{}
	d, _ := newActivity.Create(w)
	w.NewTimeOut(d, 5)
	w.NewTimeOut(d, 10)
	w.NewTimeOut(d, 15)
}
func New(
	w *wheelTimer.WheelTimer,
) error {
	startTimerTask(w)
	return nil
}