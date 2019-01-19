package wheelTimer

import (
)
type Bucket struct {
	Tasks map[int64]Task
}

// 添加任务到bucket
func (b *Bucket) Add(task *Task) {

}

// 处理当前bucket中到期的任务
func (b *Bucket) Expire() {
	for key, task := range b.Tasks {
		remainingRounds := task.GetRemainingRounds()
		if remainingRounds > 0 {
			task.UpdateRemainingRounds(remainingRounds-1)
		}else {
			delete(b.Tasks, key)
			go task.Run()

		}
	}
}

// 清除bucket
func (b *Bucket) Remove() {
}