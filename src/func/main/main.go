package main

import (
	"errors"
	"fmt"
)

func exchangeNum(num1, num2 *int) {
	t := *num1
	*num1 = *num2
	*num2 = t
}

func main() {
	//go函数支持普通函数、匿名函数、闭包

	/*
			函数名首字母不能是数字
			首字母大写该函数可以被本包文件和其它包文件使用（类似public），
			首字母小写只能被本包文件使用，其它包文件不能使用（类似private）
		函数参数传递的时候，值传递，引用传递，go语言中全部是值传递
	*/

	//num1 := 10
	//num2 := 20
	//fmt.Printf("交换前的两个数：num1 = %v， num2 = %v \n", num1, num2) // 交换前的两个数：num1 = 10， num2 = 20
	//exchangeNum(&num1, &num2)
	//fmt.Printf("交换前的两个数：num1 = %v， num2 = %v \n", num1, num2) // 交换前的两个数：num1 = 20， num2 = 10

	//1. 内存分析：值传递
	/**
	这个exchangeNum函数的功能：将形参num1和num2交换跟外界的main函数中的num1和num2毫不相关

	栈（基本数据类型）：[exchangeNum函数栈帧[num1, num2，t] ,main函数栈帧[num1，num2]] 函数执行后，程序会销毁这个函数对应开辟的栈帧（栈空间）
	堆（复杂数据类型）：
	代码区：
	*/

	/**
	2. go中函数不支持重载（函数名相同，参数列表不同）
	3. go中支持可变参数（如果你希望函数带有可变数量的参数）
	4. 基本数类型和数组默认都是值传递，即进行值拷贝。在函数内修改，不会影响到原来的值。
	5. 以值传递方式的数据类型，如果希望在函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，从效果来看类似引用传递
	*/

	// 6. 在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
	// 为了简化数据类型定义，GO支持自定义数据类型，基本语法：type 自定义数据类型名 数据类型
	// 可以理解为：相当于起了一个别名，例如 type myInt int ---> 这时myInt就等价int 来使用了
	// 例如：type mySum func(int, int) int ---> 这时mySum就等价一个函数类型 func(int, int) int

	e := exchangeNum
	fmt.Printf("e的类型%T， exchangeNum的类型是%T", e, exchangeNum) //e的类型func(*int, *int)， exchangeNum的类型是func(*int, *int)

	// 7. 函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用（把函数本身当做一种数据类型），支持对函数的返回值命名

}

// 3. 不定参数列表，go中支持可变参数（如果你希望函数带有可变数量的参数）
// 定义了一个int 类型的 items 可变参数
func add1(items ...int) (sum int, err error) {
	// 函数内部处理可变参数的时候，将可变参数当做切片来处理
	for _, val := range items {
		sum += val
	}
	return
}

func add(a int, b int) (sum int, err error) { //如果这里定义好变量名称的话 return 语句后面可以不写

	sum = a + b
	//return sum, err
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
