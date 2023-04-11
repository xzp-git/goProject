package main

import "fmt"

func main() {
	//	iota 特殊常量 可以认为是一个可以北边一起修改的变量
	// 列入 使用 const 声明常量错误码 要在中间 插入一个错误码时，需要将后续的值都改一遍会很麻烦
	//const (
	//	ERR1 = 1
	//	ERR2 = 2
	//	ERR3 = 3
	//	ERR4 = 4
	//)

	//const (
	//	ERR1 = iota
	//	ERR2
	//	ERR3
	//	ERR4
	//)
	// iota 不会规定类型 如果中断了 iota 则需要 显示的赋值 后续会自动递增
	//自增类型默认是 int 类型
	const (
		ERR1 = iota
		ERR2
		ERR21 = "haha"
		ERR22 = "haha"
		ERR23 = 100
		ERR3  = iota
		ERR4
	)
	//每次出现const iota 归零
	const (
		ERRNEW1 = iota
		ERRNEW2
		ERRNEW21 = "haha"
		ERRNEW22 = "haha"
		ERRNEW23 = 100
		ERRNEW3  = iota
		ERRNEW4
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR4)
}
