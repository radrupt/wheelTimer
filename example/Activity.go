package main

import (
	"github.com/radrupt/wheelTimer"
	"fmt"
)
type Activity struct {
	W *wheelTimer.WheelTimer
	remainingRounds int
	Name string
}
func (task *Activity) Create(
	w *wheelTimer.WheelTimer,
	name string,
) (*Activity) {
	return &Activity{
		W: w,
		Name: name,
	}
}

func (t *Activity) Run(){
	fmt.Println("excute activity: ", t.Name)
}

func (t *Activity) UpdateRemainingRounds(i int){
	t.remainingRounds = i
}

func (t *Activity) GetRemainingRounds() int {
	return t.remainingRounds
}