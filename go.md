# go

## 1. go中的变量

--- 

```go
package main

import "fmt"

// 全局变量和局部变量
var (
	name = "bobby"
	age  = 18
	ok   bool
)
var ssan = 2

func main() {
	//go是静态语言，静态语言和动态语言相比变量差异很大
	//1. 变量必须先定义后使用 2. 变量必须有类型 3. 类型定下来后不能改变
	//定义变量的方式
	//var name int
	//name = 1

	//var age = 1
	age := 1
	var age2 int

	//go语言中变量定义了不使用是不行的
	fmt.Println(age, age2)

	//2. 多变量定义
	var user1, user2, user3 = "bobby1", 1, "bobby3"
	fmt.Println(user1, user2, user3)

	/*
		注意：
			变量必须先定义才能使用
			go语言是静态语言，要求变量的类型和赋值类型一致
			变量名不能冲突
			简洁变量定义不能用于全局变量
			变量是有零值,也就是默认值
			定义了变量一定要使用，否则会报错
	*/
}
```

## 2. go中的常量

--- 

```go
package main

import "fmt"

func main() {
	const (
		//常量， 定义的时候就指定的值，不能修改， 常量尽量全部大写
		PI  float32 = 3.1415926 //显式定义
		PI2 float32 = 3.12323234
	)

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		x int = 16
		y
		s = "abc"
		z
		m
	)
	fmt.Println(x, y, s, z, m)

	/*
		常量类型只可以定义bool、数值(整数、浮点数和复数) 和 字符串
		不曾使用的常量， 没有强制使用的要求
		显示指定类型的时候，必须确保常量左右值类型一致
	    在列表定义中，如果不赋值则采用前一个常量的值
	*/
}

```

## 3. iota

---

```go
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
	//0 1 haha haha 5 6
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
	// 0 1 haha haha 100 5 6
	fmt.Println(ERR1, ERR2, ERR3, ERR4)
}

```

## 4. go中的匿名变量与作用域

--- 

```go
package main

var _ int

func a(int, bool) (int, bool) {
	return 0, false
}

func main() {
	_, ok := a(0, false)
	// 此处只需要使用 ok 变量 匿名变量 相当于占位一样 不使用 不会报错
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

```
## 5.go中的基本数据类型
- `bool`
- 数值类型
  - 整数
  - 浮点数
  - 复数
  - `byte` 字节
  - `rune` 类型
- 字符和 `string`

### 5.1 bool类型
布尔型的值只可以是常量 `true` 或者 `false` 。一个简单的例子: `var b bool = true`
### 5.2 数值型
#### 5.2.1 整数型
可以简单讲解一下二进制和位数的关系,以及 `int` 和 `uint` 的关系
- `int` 类型是根据操作系统来定义的，64位的操作系统就是64 32位的操作系统就是32
- `int8` 有符号8位整型(-128到127)长度:8bit
- `int16` 有符号16位整型(-32768到32767)
- `int32` 有符号32位整型(-2147483648到2147483647)
- `int64` 有符号64位整型(-9223372036854775808到92233720336854775807
- `uint8` 无符号8位整型(0到255)8位都用于表示数值:
- `uint16` 无符号16位整型(0到65535)
- `uint32` 无符号32位整型(0到4294967295)
- `uint64` 无符号64位整型(0到18446744073709551615)
#### 5.2.2 浮点型
- `float32` 32位浮点型数
- `float64` 64位浮点型数
#### 5.2.3 其他
- `byte` 等于 `uint8`
- `rune` 等于 `int32`
 