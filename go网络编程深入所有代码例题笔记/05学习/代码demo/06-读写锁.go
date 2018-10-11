package main

import (
	"runtime"
	"math/rand"
	"time"
	"fmt"
	"sync"
)
// 使用读写锁
var rwMutex sync.RWMutex

var counter int = 0			// 模拟公共区

func readGo(idx int)  {		// 读go程
	for {								// 一个go程可以读 无限 次。
		rwMutex.RLock()		// 读模式加  读写锁
		num := counter			// 从公共区中，读取数据
		fmt.Printf("%dth 读 go程，读到：%d\n", idx, num)
		rwMutex.RUnlock()	// 解锁 读写锁

		time.Sleep(time.Millisecond * 500)		// 放大实验现象
	}
}

func writeGo(idx int)  {
	for {									// 一个go程可以写 无限 次。
		// 生产一个随机数
		num := rand.Intn(500)
		rwMutex.Lock()		// 写模式加  读写锁
		counter = num		// 写数据到 公共区中
		fmt.Printf("-----%dth 写 go程，写入：%d\n", idx, counter)
		rwMutex.Unlock()	// 解锁  读写锁

		time.Sleep(time.Millisecond * 200)		// 放大实验现象
	}
}

func main()  {
	// 播种随机数种子。
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<5; i++ {		// 同时创建 N 个 读go程
		go readGo(i+1)
	}
	for i:=0; i<5; i++ {		// 同时创建 N 个 写go程
		go writeGo(i+1)
	}
	for {						// 防止 主 go 程 退出
		runtime.GC()
	}
}
