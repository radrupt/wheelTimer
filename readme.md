## 轮盘延迟消息

#### 安装
    go get github.com/radrupt/wheelTimer

#### 使用
    //实现接口 Task
    type Task interface {
      Run()
      UpdateRemainingRounds(i int)
      GetRemainingRounds() int
    }

    //创建需要执行的消息结构
    package main
    import (
      "github.com/radrupt/wheelTimer"
      "fmt"
    )
    type Activity struct {
      W *wheelTimer.WheelTimer //必须
      remainingRounds int //必须

      Name string //自定义数据
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