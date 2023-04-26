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

---

- `bool`
- 数值类型
  - 整数
  - 浮点数
  - 复数
  - `byte`
  - `rune` 
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
- `float32` 32位浮点型数 （3.4 * 10 ^ 38）
- `float64` 64位浮点型数（ 1.8 * 10 ^ 308）
#### 5.2.3 其他
- `byte` 等于 `uint8`  主要用于存放字符的ascall码 
- `rune` 等于 `int32`  `byte` 类型表示的字符类型有限, 这个类型也是用来表示字符的. 例如 中文汉字
- 如果处理的字符只是英文字符, 使用 `byte` 就可以, 如果有中文有英文，使用 `rune` 类型

```go
package main

import "fmt"

/*




 */

func main() {
	//var num int8 = 127
	//
	//var f1 float32 = 399.3333

	var c byte

	c = 'a'
	c1 := 97

	fmt.Println(c)        //  97   打印 a 字符的 ASCll 码
	fmt.Printf("c=%c", c) //使用格式化打印 字符串
	fmt.Println(c1)
	fmt.Printf("c1=%c1", c1)
	fmt.Println("")

	var b rune

	b = '木'

	fmt.Println(b)

	var name string = "fooo"

	fmt.Println(name)
}
```

## 6. go中的基本数据类型转换

--- 

 ```go

package main

import (
	"fmt"
	"strconv"
)


func main() {

	// 浮点数
	a := 5.11
	//转换为 整数
	b := int(a)
	fmt.Println(a)
	fmt.Println(b)

	//	go允许在底层结构相同的两个类型之间互转
	//IT类型的底层是int类型

	type IT int

	var c IT = 199

	//将c IT 类型 转换为 int， b现在是 int 类型
	var d int = int(c)
	var e IT = IT(b)

	fmt.Println(e)

	fmt.Println(d)

	//	字符串 转数字
	var str = "12e"
	myInt, error := strconv.Atoi(str)

	if error != nil {
		fmt.Println(error)
		fmt.Println("Atoi error")
	}
	fmt.Println(myInt)

	//数字转字符串

	var num = 998

	mystr := strconv.Itoa(num)

	fmt.Println(mystr)
}

```

## 7.字符串格式化转换

--- 

 ```go
   {
		//字符串转换为float32
		value, error := strconv.ParseFloat("3.11414", 64)

		if error != nil {
			fmt.Println(error)
			fmt.Println("Atoi error")
		}
		fmt.Println(value)
	}

	{
		//字符串转换为int
		//这里 base 参数是指 你的 字符串参数 是 几进制 然后 吧这个 base进制的字符串 转换成bitSize位的十进制数字
		value, error := strconv.ParseInt("11", 2, 64)

		if error != nil {
			fmt.Println(error)
			fmt.Println("Atoi error")
		}
		fmt.Println(value)
	}

	{
		//字符串转换为bool
		//这里 base 参数是指 你的 字符串参数 是 几进制 然后 吧这个 base进制的字符串 转换成bitSize位的十进制数字
		value, error := strconv.ParseBool("0")
		// "0"  -> false   "1" -> true   "true" -> true  否则就是 false
		if error != nil {
			fmt.Println(error)
			fmt.Println("Atoi error")
		}
		fmt.Println(value)
	}

	{
		//基本类型转字符串
		value := strconv.FormatBool(true)
		fmt.Println(value)
	}

	{
		//基本类型转字符串
		value := strconv.FormatFloat(3.141555, 'f', -1, 64)
		fmt.Println(value)
	}

	{
		//基本类型转字符串
		value := strconv.FormatInt(42, 16)
		fmt.Println(value)
	}
```
 
## 8. 特殊运算符

--- 

```go
var A = 1
c := &A // 代表取 A 变量的地址 

var C *int //指针类型
```
![img.png](img/img-1.png)

## 9.字符串

---

### 9.1 rune 和 字符串长度

---

```go
package main

import "fmt"

func main() {
	//长度计算
	name := "fooo中文"
	fmt.Println(len(name)) // 10  len 返回的是字符串的字节数 一个中文 占 三个字节

	//切片
	bytes := []byte(name)

	fmt.Println(bytes)
	fmt.Println(len(bytes))

	//字符长度, 如果 字符串中有中文，需要转换成 rune 类型

	runes := []rune(name)

	fmt.Println(runes)
	fmt.Println(len(runes)) // 6

	//单引号 和 双引号  单引号包裹的是byte类型或者rune类型  双引号包裹的才是字符串类型
	age := "年"
	var c byte = 'c'
	d := '年' // rune 类型

	fmt.Println(age)
	fmt.Println(c)
	fmt.Println(d)
}

```

### 9.2 转义符

---

```go
package main

import "fmt"

func main() {


	//转义符

	name := "张\"\r\n三\"" // \r \n  回车换行

	fmt.Println(name) //张"
	                  //三"  
	name1 := `张"三"111`
	
	fmt.Println(name1)//张"三"

}

```
 
### 9.3 格式化输出

---

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	//格式化输出

	name := "foo"
	age := 18
	address := "北京 "
	fmt.Println("用户名：" + name + ",  年龄：" + strconv.Itoa(age) + ",  地址：" + address) //拼凑 字符串 很难维护

	fmt.Printf("用户名：%s, 年龄：%d, 地址：%s", name, age, address) //这个 很常用 但是 性能没有 上面的好

	useMsg := fmt.Sprintf("\r\nSprintf \r\n用户名：%s, 年龄：%d, 地址：%s", name, age, address) // 生成字符串  Printf 只管打印

	fmt.Println(useMsg)
}

```

![img.png](img/img-2.png)

![img.png](img/img-3.png)

![img.png](img/img-4.png)

![img.png](img/img-5.png)
 
![img.png](img/img-6.png)

### 9.4 高性能的字符串拼接 - strings.builder

---

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {


	name := "foo"
	age := 18
	address := "北京 "

	//通过string的buildr进行字符串拼接   高性能

	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(name)
	builder.WriteString(",年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址:")
	builder.WriteString(address)

	re := builder.String()

	fmt.Println(re)

}
```


### 9.5 字符串的比较

--- 

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {


  //字符串比较
  a := "a"
  b := "a"

  fmt.Println(a == b)
  fmt.Println(a != b)

  //字符串的 大小比较 从左往右 根据 ascall码 比较 相同就比较下一个的ascall 码

  c := "baello"
  d := "bello"

  fmt.Println(c > d)

}
```

### 9.6 字符串常用的方法

```go
package main

import (
	"fmt"
	"strings"
)

func main() {

  name := "fooo-bar"
  //是否包含
  fmt.Println(strings.Contains(name, "bar"))
  //子串出现的次数
  fmt.Println(strings.Count(name, "o"))
  //分割
  fmt.Println(strings.Split(name, "-"))

  //	字符串是否包含前缀  字符串是否包含后缀
  fmt.Println(strings.HasPrefix(name, "fo"))
  fmt.Println(strings.HasSuffix(name, "bar"))

  //查询子串出现的位置
  fmt.Println(strings.Index(name, "-"))

  //替换

  fmt.Println(strings.Replace(name, "fooo", "aaa", -1))

  //大小写 转换
  fmt.Println(strings.ToLower("GGG"))
  fmt.Println(strings.ToUpper("ggg"))

  //去掉字符串 前后的字符
  fmt.Println(strings.Trim("   -wf- w e-  --", " -"))
  
}
```