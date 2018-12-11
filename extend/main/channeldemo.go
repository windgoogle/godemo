package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  通道在箭头左边就是写通道
  通道在箭头右边就是读通道
  通道只有一种创建方法，make（）函数
  // 缓冲型通道
	var bufferedChannel = make(chan int, 1024)
	// 非缓冲型通道
	var unbufferedChannel = make(chan int)
    非缓冲型通道必须确保有协程正在尝试读取当前通道，否则写操作就会阻塞直到有其它协程来从通道中读东西。
    非缓冲型通道总是处于既满又空的状态。与之对应的有限定大小的通道就是缓冲型通道。
    在 Go 语言里不存在无界通道，每个通道都是有限定最大容量的。
*/
func main() {
	//chanbase()
	//recv_and_send()
	closechan()
}

func chanbase() {
	var ch chan int = make(chan int, 4)
	for i := 0; i < cap(ch); i++ {
		ch <- i // 写通道
	}
	for len(ch) > 0 {
		var value int = <-ch // 读通道
		fmt.Println(value)
	}
}

func send(ch chan int) {
	for {
		var value = rand.Intn(100)
		ch <- value
		fmt.Printf("send %d\n", value)
	}
}

func recv(ch chan int) {
	for {
		value := <-ch
		fmt.Printf("recv %d\n", value)
		time.Sleep(time.Second)
	}
}

func recv_and_send() {
	var ch = make(chan int, 1)
	// 子协程循环读
	go recv(ch)
	// 主协程循环写
	send(ch)
}

func closechan() {
	var ch = make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)

	value := <-ch
	fmt.Println(value)
	value = <-ch
	fmt.Println(value) //并不会马上关闭

	//value = <- ch       // 读一个已关闭的通道，返回通道类型的零值
	//fmt.Println(value)

	// for range 遍历通道 ,没有上面调用close()，会出错 fatal error: all goroutines are asleep - deadlock!
	for value := range ch {
		fmt.Println(value)
	}

	//ch <- 3    // panic: send on closed channel
}

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
