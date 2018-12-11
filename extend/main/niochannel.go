package main

import "fmt"
import "time"

/*
非阻塞读写依靠select 的default分支的例子
直观的结果是，下面程序运行的例子数字读取的不是连续的，很多数据被丢掉了，去掉default,则变成连续的
*/
func nio_send(ch1 chan int, ch2 chan int) {
	i := 0
	for {
		i++
		select {
		case ch1 <- i:
			fmt.Printf("send ch1 %d\n", i)
		case ch2 <- i:
			fmt.Printf("send ch2 %d\n", i)
			//default: //这个很关键，区分是否NIO
		}
	}
}

func nio_recv(ch chan int, gap time.Duration, name string) {
	for v := range ch {
		fmt.Printf("receive %s %d\n", name, v)
		time.Sleep(gap)
	}
}

func main() {
	// 无缓冲通道
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	// 两个消费者的休眠时间不一样，名称不一样
	go nio_recv(ch1, time.Second, "ch1")
	go nio_recv(ch2, 2*time.Second, "ch2")
	nio_send(ch1, ch2)
}

/**
channel 内部本质是个循环数组
type hchan struct {
  qcount uint  // 通道有效元素个数
  dataqsize uint   // 通道容量，循环数组总长度
  buf unsafe.Pointer // 数组地址
  elemsize uint16 // 内部元素的大小
  closed uint32 // 是否已关闭 0或者1
  elemtype *_type // 内部元素类型信息
  sendx uint // 循环数组的写偏移量
  recvx uint // 循环数组的读偏移量
  recvq waitq // 阻塞在读操作上的协程队列
  sendq waitq // 阻塞在写操作上的协程队列

  lock mutex // 全局锁
}
*/
