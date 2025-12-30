package main

import (
	"fmt"
	"os"
	"sync"
)

func foo() (x int) {
	defer func() {
		x = x + 1
	}()
	return 10
}

func tool() int {
	x := 10
	defer fmt.Println(x)
	// 修改 defer 前的 x 不會影響到 defer 印出的值
	x = 20
	return x
}

// 延後讀值 閉包會抓變數，不是抓值
func deferRead() {
	x := 10
	defer func() {
		fmt.Println("DeferRead: ", x)
	}() // 傳入當下的 x 值 (10)
	x = 20
}

// 1. 關閉資源 close resource
func deferClosure() error {
	file, err := os.Open("a.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	// 進行檔案操作
	return nil
}

// 2. mutex 解鎖 unlock mutex
func deferUnlock(mutex *sync.Mutex) error {
	mutex.Lock()
	defer mutex.Unlock()
	// 臨界區 critical section
	return nil
}

// 3. recover panic（高階）
func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recovered:", r)
		}
	}()

	panic("boom")
}

func main() {
	result := foo()
	fmt.Println("foo: ", result)  // 輸出 11
	fmt.Println("tool: ", tool()) // 輸出 10
	deferRead()                   // 輸出 20 閉包會抓變數，不是抓值
}
