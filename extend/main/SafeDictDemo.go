package main

import (
	"log"
	"os"
)
import "sync"

var Trace = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

type SafeDict struct {
	data        map[string]int
	*sync.Mutex //外部结构体继承匿名结构体所有方法
	// mutex *sync.Mutex
}

func NewSafeDict(data map[string]int) *SafeDict {
	//return &SafeDict{
	//	data:  data,
	//	&sync.Mutex{},
	//}
	return &SafeDict{data, &sync.Mutex{}}
}

func (d *SafeDict) Len() int {
	//d.mutex.Lock()
	d.Lock() //外部结构体继承匿名结构体所有方法
	//defer d.mutex.Unlock()
	defer d.Unlock()
	return len(d.data)
}

func (d *SafeDict) Put(key string, value int) (int, bool) {
	//d.mutex.Lock()
	d.Lock()
	//defer d.mutex.Unlock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	d.data[key] = value
	return old_value, ok
}

func (d *SafeDict) Get(key string) (int, bool) {
	//d.mutex.Lock()
	d.Lock()
	//defer d.mutex.Unlock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (int, bool) {
	//d.mutex.Lock()
	d.Lock()
	//defer d.mutex.Unlock()
	d.Unlock()
	old_value, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return old_value, ok
}

func write0(d *SafeDict) {
	d.Put("banana", 5)
	Trace.Println("write method")
	//time.Sleep(100000)
}

func read0(d *SafeDict) {
	//fmt.Println(d.Get("banana"))
	Trace.Println(d.Get("pear"))
}

func main() {
	d := NewSafeDict(map[string]int{
		"apple": 2,
		"pear":  3,
	})
	go read0(d)
	write0(d)
}
