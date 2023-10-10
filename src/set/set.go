package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	// go 语言提供了哪些集合类型的数据结构 数组、 *切片（slice ）、 * map、 list

	//数组   courses1 和 courses2 的类型不一样
	var courses1 [3]string //有3个字符串元素的 数组
	var courses2 [4]string //有4个字符串元素的 数组
	var courses3 []string  // [3]string 和  []string  是两种不一样的类型   [3]string类型 是数组 []string类型是切片

	courses1[0] = "go"
	courses1[1] = "grpc"
	courses1[2] = "gin"

	fmt.Println(courses1)

	fmt.Printf("%T\r\n", courses1) //[3]string
	fmt.Printf("%T\r\n", courses2) //[4]string
	fmt.Printf("%T\r\n", courses3) //[]string

	for _, value := range courses1 {
		fmt.Println(value)
	}

	//数组的初始化
	courses4 := [3]string{"go", "grpc", "gin"}

	fmt.Println(courses4)

	courses5 := [3]string{2: "gin"} //没有初始化的位置 是 "" 空字符串
	for _, value := range courses5 {
		fmt.Println(value == "")
	}
	fmt.Println(courses5)

	courses6 := [...]string{"go", "grpc", "gin"} //初始化是几个  类型长度就是几
	fmt.Printf("%T\r\n", courses6)               //[3]string
	fmt.Println(courses5)

	if courses4 == courses6 { //比较是比较数组的每一项是否相等
		fmt.Println("equal")
	}

	// 多维数组
	var courses7 = [2][2]string{{"go", "1h"}, {"grpc", "2h"}}
	fmt.Println(courses7)

	//切片
	var courses8 []string

	//append的用法   往 courses8 中 添加 元素 必须使用 courses8 接受
	courses8 = append(courses8, "go")
	courses8 = append(courses8, "grpc")

	//切片的初始化  3 种
	//1.从数组直接创建 2.使用 string{} 3. make
	//1.从数组创建切片
	allCourses4 := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	allCourses4Slice := allCourses4[0:1]

	fmt.Println(allCourses4Slice)
	fmt.Printf("%T\r\n", allCourses4Slice)

	//2. 直接创建切片   allCourses4 := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	//3. make函数
	allCourses4SliceForMake := make([]string, 3)
	allCourses4SliceForMake[0] = "c"
	allCourses4SliceForMake[1] = "c"
	allCourses4SliceForMake[2] = "c"
	//allCourses4SliceForMake[3] = "c"  使用make 创建切片 超过初始化长度 回报错
	fmt.Println(allCourses4SliceForMake)

	//tips:

	var allCoursesNoLength []string

	//allCoursesNoLength[0] = "c"  如果slice 没有声明初始化长度 这样赋值 会报错
	//需要使用append方法
	allCoursesNoLength = append(allCoursesNoLength, "c")

	fmt.Println(allCoursesNoLength)

	// 切片的访问 访问单个  访问多个
	allCourses := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	fmt.Println(allCourses[0])
	fmt.Println(allCourses[1:4])
	fmt.Println(allCourses[1:]) // 只有 start  没有end  表示 取到完
	fmt.Println(allCourses[:3]) //只有end  表示 取到 开头
	fmt.Println(allCourses[:])  //没有 start 没有 end  表示 复制了一份

	//切片的数据 添加

	allCoursesForAdd := []string{"go", "grpc"}
	allCoursesForAdd2 := []string{"mysql", "es"}

	//allCoursesForAdd = append(allCoursesForAdd, "gin", "mysql", "es ")

	//方法1 for 循环合并
	//for _, value := range allCoursesForAdd2 {
	//	allCoursesForAdd = append(allCoursesForAdd, value)
	//}

	//方法2 ...  合并
	allCoursesForAdd = append(allCoursesForAdd, allCoursesForAdd2[1:]...)

	fmt.Println(allCoursesForAdd)

	//go的slice在函数参数传递的时候是值传递还是引用传递：值传递、效果又呈现出了引用的效果（不完全是）
	{
		allCourses := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
		printSlice(allCourses)
		//此处 append的数据并没有打印出来
		fmt.Println(allCourses)
	}

	//slice的底层原理
	{
		//courses8 = append(courses8, "go")

		//append 方法 如果此时数组长度不够，会进行扩容，如果 长度小于 1024 会成倍扩容，会创建一个新的数组，将值拷贝出来，所以需要赋值给原来的变量，因为此时slice的内存地址已经改变了
	}

}

// 结构体
type slice struct {
	array unsafe.Pointer //用来存储实际数据的数组指针，指向一块连续的内存
	len   int            //切片中元素的数量
	cap   int            //array数组的长度
}

func printSlice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}
