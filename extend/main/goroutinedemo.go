package main

import (
	"fmt"
	"runtime"
)
import "time"

/**
GO 里只有一个主协程，其他是子协程，是平行关系，尽管有些子行程是其他子协程中创建的
子协程异常会导致主协程退出
*/

func main() {
	//base()  基本测试
	//large()  //百万协程测试
	//deadloop()
	numofgoroutines()
}

func base() {
	fmt.Println("run in main goroutine")
	go func() {
		fmt.Println("run in child goroutine")
		go func() {
			fmt.Println("run in grand child goroutine")
			go func() {
				fmt.Println("run in grand grand child goroutine")
				panic("wtf") //子协程异常会导致主协程退出
			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("main goroutine will quit")
}

//百万大协程测试
func large() {
	fmt.Println("run in main goroutine")
	i := 1
	for {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
		if i%10000 == 0 {
			fmt.Printf("%d goroutine started\n", i)
		}
		i++
	}

}

func deadloop() {
	fmt.Println("run in main goroutine")
	n := 3
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println("dead loop goroutine start")
			for {
			} // 死循环
		}()
	}
	for {
		time.Sleep(time.Second)
		fmt.Println("main goroutine running")
	}
}

func numofgoroutines() {
	fmt.Println(runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
	fmt.Println(runtime.NumGoroutine())
}
