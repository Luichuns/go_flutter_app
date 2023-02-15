package main

import (
	"fmt"
	"strconv"
	"sync"
)

// m 声明1个键值对 键string 值int
var m = make(map[string]int)

// 获取 m中键的值
func get(key string) int {
	return m[key]
}

// 添加m 中的键值
func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{} //声明等待组
	for i := 0; i < 10; i++ {
		wg.Add(1)        //每个循环添加1次同步
		go func(n int) { //异步匿名函数
			key := strconv.Itoa(n)
			set(key, n) //添加键值
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i) //把循环的i值 作为异步匿名函数的n参数--`
	}
	wg.Wait()
}
