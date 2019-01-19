package main

import (
	"fmt"
	"time"
)
type Activity struct {
	remainingRounds int
	Name string
}

func (t *Activity) Run(){
	fmt.Println("excute activity: ", t.Name)
	fmt.Println(time.Now())
}

func (t *Activity) UpdateRemainingRounds(i int){
	t.remainingRounds = i
}

func (t *Activity) GetRemainingRounds() int {
	return t.remainingRounds
}