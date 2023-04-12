package main

var _ int

func a(int, bool) (int, bool) {
	return 0, false
}

func main() {
	_, ok := a(0, false)
	if ok {

	}
}

/**
变量的作用域

全局变量 任何地方都可以用到

函数内部定义的变量 在函数内部可以使用，出了函数就不可以使用了
{

} 双花括号内也可以写代码语句，代码块内部的变量 外部不可以使用
*/
