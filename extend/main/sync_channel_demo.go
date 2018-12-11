package main

import "fmt"
import "time"
import "sync"

/*
     向一个已经关闭的通道执行写操作会抛出异常，这意味着我们在写通道时一定要确保通道没有被关闭。
     那如何确保呢？Go 语言并不存在一个内置函数可以判断出通道是否已经被关闭。即使存在这样一个函数，当你判断时通道没有关闭，
     并不意味着当你往通道里写数据时它就一定没有被关闭，并发环境下，它是可能被其它协程随时关闭的。
     确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。
	 这个方法确实可以解决单写多读的场景，可要是遇上了多写单读的场合该怎么办呢？
     任意一个读写通道的协程都不可以随意关闭通道，否则会导致其它写通道协程抛出异常。
     这时候就必须让其它不相干的协程来干这件事，这个协程需要等待所有的写通道协程都结束运行后才能关闭通道。
     那其它协程要如何才能知道所有的写通道已经结束运行了呢？
      这个就需要使用到内置 sync 包提供的 WaitGroup 对象，它使用计数来等待指定事件完成。


*/

func sync_send(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 计数值减一
	i := 0
	for i < 4 {
		i++
		ch <- i
	}
}

func sync_recv(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	var ch = make(chan int, 4)
	var wg = new(sync.WaitGroup)
	wg.Add(2)            // 增加计数值
	go sync_send(ch, wg) // 写
	go sync_send(ch, wg) // 写
	go sync_recv(ch)
	// Wait() 阻塞等待所有的写通道协程结束
	// 待计数值变成零，Wait() 才会返回
	wg.Wait()
	//关闭通道
	close(ch)
	time.Sleep(time.Second)
}
