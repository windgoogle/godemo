package main

import "fmt"
import "time"

/*
 多路复用，将多个通道的数据汇聚到一起读取
*/

// 每隔一会生产一个数
func multi_send(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一的目标通道
func collect(source chan int, target chan int) {
	for v := range source {
		target <- v
	}
}

// 从目标通道消费数据
func multi_recv(ch chan int) {
	for v := range ch {
		fmt.Printf("receive %d\n", v)
	}
}

// 多路复用 select
func multi_recv0(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("recv %d from ch1\n", v)
		case v := <-ch2:
			fmt.Printf("recv %d from ch2\n", v)
		}
	}
}

func main() {
	//无缓冲通道
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go multi_send(ch1, time.Second)
	go multi_send(ch2, 2*time.Second)
	go collect(ch1, ch3)
	go collect(ch2, ch3)
	//multi_recv(ch3)
	multi_recv0(ch1, ch2)

}
