package main

// go run -gcflags="all=-m -m" main.go

import "fmt"

// 範例 1：回傳局部變數的指標（最經典）
func retLocalPtr() *int {
	x := 42
	return &x // x 會逃逸：離開函式後還要活著
}

// 範例 2：把指標塞進 interface{}（常造成逃逸）
func toInterface() interface{} {
	x := 123
	return x // x 可能因為被裝箱成 interface 而逃逸（依編譯器決策）
}

func ptrInInterface() interface{} {
	x := 777
	return &x // x 幾乎必逃逸
}

// 範例 3：閉包捕獲外部變數（變數要延長生命週期）
func makeClosure() func() int {
	x := 10
	return func() int {
		x++
		return x
	} // x 通常會逃逸：closure 需要保存它
}

// 範例 4：把局部指標存到全域（生命週期被拉長）
var gp *int

func storeGlobal() {
	x := 99
	gp = &x // x 必逃逸：全域變數可能在程式整個生命週期被用到
}

// 範例 5：slice append 造成底層陣列配置到 heap（常見於回傳 slice）
func buildSlice() []int {
	s := make([]int, 0, 2)
	for i := 0; i < 5; i++ {
		s = append(s, i) // 容量不夠會擴容 -> 可能配置新底層陣列到 heap
	}
	return s // 回傳 slice，底層資料必須在函式外可用
}

// 範例 6：大物件在 stack 放不下，可能被搬到 heap（依平台/版本）
type Big struct {
	buf [1 << 20]byte // 1MB
}

func bigOnStack() Big {
	var b Big
	b.buf[0] = 1
	return b // 可能觸發逃逸/搬移（依編譯器與 ABI）
}

func main() {
	// 1
	p := retLocalPtr()
	fmt.Println(*p)
	// 2
	v := toInterface()
	fmt.Println(v)
	//
	fmt.Println(ptrInInterface())
	// 3
	f := makeClosure()
	fmt.Println(f())
	fmt.Println(f())
	// 範例 4：把局部指標存到全域（生命週期被拉長）
	storeGlobal()
	fmt.Println(*gp)
	// 5
	fmt.Println(buildSlice())
	// 6
	b := bigOnStack()
	fmt.Println(b.buf[0])
}
