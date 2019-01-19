package wheelTimer

type Task interface {
	Run()
	UpdateRemainingRounds(i int)
	GetRemainingRounds() int
}


//使用

// type XX struct {
// 	remainingRounds int
// }

// func (n *Notifier) Run(){
// 	// todo
// }

// func (n *Notifier) UpdateRemainingRounds(i int){
// 	// todo
// }

// func (n *Notifier) GetRemainingRounds() int {
// 	// todo
// }