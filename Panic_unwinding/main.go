package main

import "fmt"

func panicExample() {
	panic("something wrong")

	var a []int
	fmt.Println(a[0]) // index out of range

	var m map[string]int
	fmt.Println(m["x"]) // 不會 panic（回傳 0）

}

func main() {
	panicExample()
}
