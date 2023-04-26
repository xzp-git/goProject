package main

import (
	"fmt"
	"strings"
)

func main() {
	//长度计算
	//name := "fooo中文"
	//fmt.Println(len(name)) // 10  len 返回的是字符串的字节数 一个中文 占 三个字节
	//
	////切片
	//bytes := []byte(name)
	//
	//fmt.Println(bytes)
	//fmt.Println(len(bytes))
	//
	////字符长度, 如果 字符串中有中文，需要转换成 rune 类型
	//
	//runes := []rune(name)
	//
	//fmt.Println(runes)
	//fmt.Println(len(runes)) // 6
	//
	////单引号 和 双引号  单引号包裹的是byte类型或者rune类型  双引号包裹的才是字符串类型
	//age := "年"
	//var c byte = 'c'
	//d := '年' // rune 类型
	//
	//fmt.Println(age)
	//fmt.Println(c)
	//fmt.Println(d)

	//转义符

	//name := "张\"\r\n三\"" // \r \n  回车换行
	//
	//fmt.Println(name)
	//
	//name1 := `张"三"111`
	//
	//fmt.Println(name1) //张"三"111

	//格式化输出

	//name := "foo"
	//age := 18
	//address := "北京 "
	////fmt.Println("用户名：" + name + ",  年龄：" + strconv.Itoa(age) + ",  地址：" + address) //拼凑 字符串 很难维护
	////
	////fmt.Printf("用户名：%s, 年龄：%d, 地址：%s", name, age, address) //这个 很常用 但是 性能没有 上面的好
	////
	////useMsg := fmt.Sprintf("\r\nSprintf \r\n用户名：%s, 年龄：%d, 地址：%s", name, age, address) // 生成字符串  Printf 只管打印
	////
	////fmt.Println(useMsg)
	//
	////通过string的buildr进行字符串拼接   高性能
	//
	//var builder strings.Builder
	//builder.WriteString("用户名:")
	//builder.WriteString(name)
	//builder.WriteString(",年龄:")
	//builder.WriteString(strconv.Itoa(age))
	//builder.WriteString(",地址:")
	//builder.WriteString(address)
	//
	//re := builder.String()
	//
	//fmt.Println(re)

	//字符串比较
	//a := "a"
	//b := "a"
	//
	//fmt.Println(a == b)
	//fmt.Println(a != b)
	//
	////字符串的 大小比较 从左往右 根据 ascall码 比较 相同就比较下一个的ascall 码
	//
	//c := "baello"
	//d := "bello"
	//
	//fmt.Println(c > d)

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
