package main

import (
	"fmt"
	"goProject/src/func/init/testInit"
)

/**
1. init 函数：初始化函数，可以用来进行一些初始化的操作，每一个源文件都可以包含一个init 函数，该函数会在main函数执行前，被GO运行框架调用。

2. 全局变量定义 init函数，main函数 都有的话 执行流程是怎么样
test 函数执行
init 函数被执行
main 函数被执行

3. 多个源文件都有init函数的时候，如何执行

testInit包的init函数  --- 先执行 引用的包的init函数
test 函数执行  --- 全局变量
init 函数被执行  --- 本包的init函数
main 函数被执行 --- 本包的main函数
Age= 20
Name= 我是谁

*/

var num = test()

func main() {

	/**
	testInit包的init函数
	test 函数执行
	init 函数被执行
	main 函数被执行
	Age= 20
	Name= 我是谁
	*/
	fmt.Println("main 函数被执行")
	fmt.Println("Age=", testInit.Age)
	fmt.Println("Name=", testInit.Name)
}

func init() {
	fmt.Println("init 函数被执行")
}

func test() int {
	fmt.Println("test 函数执行")
	return 10
}
