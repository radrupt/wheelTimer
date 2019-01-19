## 轮盘延迟消息

#### 安装
    go get github.com/radrupt/wheelTimer

#### 使用
##### 需实现接口 Task
    type Task interface {
      Run()
      UpdateRemainingRounds(i int)
      GetRemainingRounds() int
    }

##### 创建需要执行的消息结构 example/Activity.go
    package main
    import (
      "fmt"
    )

    type Activity struct {
      remainingRounds int //必须
      Name string //自定义数据
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

##### 开启
    func startLateMessage() {
      w := wheelTimer.New(1,256) / 创建定时任务
      // 添加并启动定时任务
      w.NewTimeOut(&Activity{Name: "A"}, 5)
      w.NewTimeOut(&Activity{Name: "B"}, 10)
      w.NewTimeOut(&Activity{Name: "C"}, 15)
    }