package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	//go函数支持普通函数、匿名函数、闭包

	/*
			go中函数是"一等公民"
			1. 函数本身是可以当作变量
			2. 匿名函数 闭包
			3. 函数是可以满足接口的

		函数参数传递的时候，值传递，引用传递，go语言中全部是值传递
	*/

	sum, _ := add(1, 10)
	fmt.Println(sum)

	sum1, _ := add1(1, 10, 10)
	fmt.Println(sum1)

	var testMap = map[string]string{
		"mysql": "111",
	}
	test(testMap)

	fmt.Println(testMap) // map[mysql:test]

	//defer  链接数据库、打开文件、开始锁， 无论如何 最后都要记得去关闭数据库、关闭文件、解锁
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock() //defer后面的代码是会放在return之前执行的 多个defer的话， 是和栈一样的概念，先进后执行

	A()

}

// func add(a, b int) int {  参数的类型一样可以简写
//func add(a int, b int) (int, error) {
//	return a + b, nil
//}

func add(a int, b int) (sum int, err error) { //如果这里定义好变量名称的话 return 语句后面可以不写

	sum = a + b
	//return sum, err
	return
}

func test(a map[string]string) {
	a["mysql"] = "test"
	
}

// 不定参数列表
func add1(items ...int) (sum int, err error) {

	for _, val := range items {
		sum += val
	}
	return
}

//go语言错误处理的理念,
//error panic  recover
//一个函数可能出错,trycatch去包住这个函数,
//11开发函数的人需要有一个返回值去告诉调用者是否成功,
////go设计者认为必须要处理这个error,防御编程
//go设计者要求我们必须要处理这个err,代码中大量的会出现iferr!=nil

func A() (int, error) {
	//panic("this is an panic")//panic会导致程序的退出4,平时开发中不要随便用,一般我们在哪里用到:我们一个服务各启动的过程中
	//比如我的服务想要启动,必须有些依赖服务准备好,日志文件存在、mysql能链接通、比如配置文件没有问题,这个时候服务方能启动的时候
	//如果我们的服务启动检查中发现了这些任何一个不满足那就调用panic主动调用
	//但是你的服务一旦启动了,这个时候你的某行代码中不小心心panic那么不好意思你的程序挂了,这是重大事故
	//但是架不住有些地方我们的代码写的不小心会导致被动触力panic   例如 map 不初始化 就赋值
	//recover这个函数能捕获到panic

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var names map[string]string
	names["go"] = "go工程师"
	//fmt.Println("this is a func")
	return 0, errors.New("this is an error")

	/**
	//1.defer需要放在panic之前定义,另外 recover只有在defer调用的函数中才会生效
	//2.recover处理异常后,逻辑并不会恢复到panic的那个点去
	//3.多个defer会形成栈,后定义的defer会先执行

	*/
}
