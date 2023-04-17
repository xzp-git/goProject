package main

import (
	"fmt"
	"strconv"
)

/*




 */

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

}
