考核

1.（a）

协程比主程慢，无lock就先unlock了

（b）

```go
package main

import (
   "fmt"
)

func main() {
   var channel = make(chan string)
   var send="下山的路又堵起了"
   var receive string
   go func() {
      channel <- send
   }()
   receive = <- channel
   fmt.Println(receive)

}
```

（c）

```go
package main

import (
   "fmt"
   "sync"
)

func main() {
   begin := make(chan interface{})
   var group sync.WaitGroup
   for i := 0; i < 10; i++ {
      group.Add(1)
      go func(x int) {
         defer group.Done()
         <-begin
         fmt.Println("恭喜不知道叫什么名字队获得ACM校赛二等奖")
      }(i)
   }
   fmt.Println("打印开始")
   close(begin)
   group.Wait()
   fmt.Println("打印完毕")


}
```

（d）

想法：先寻找最精简求素数算法，每个数为其开通一个协程

参考：知乎，CSDNgo高并发原理



2.

交了文件因为不知道为啥括号有点问题，交了c++的转化成go就行了（有事赶时间先溜了）。



3.
