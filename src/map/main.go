package main

import (
	"fmt"
)

func main() {
	var courseMap = map[string]string{
		"go":   "go工程师",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}

	// 取值
	fmt.Println(courseMap["grpc"])

	//放值
	courseMap["mysql"] = "mysql的原理"

	fmt.Println(courseMap)

	/**
	var nilMap map[string]string //  只定义 初始值 是 nil
	// 如果一个map 没有初始化 这样赋值会报错
	nilMap["mysql"] = "mysql"
	*/

	//var nilMap = map[string]string{}

	var nilMap = make(map[string]string) //make 是内置函数， 主要用于slice map channel的初始化

	nilMap["mysql"] = "mysql" //需要这样初始化才不会报错
	nilMap["mysql1"] = "mysql1"
	nilMap["mysql2"] = "mysql2"
	nilMap["mysql3"] = "mysql3"

	fmt.Println(nilMap)

	//	遍历 map结构是无序的  每次遍历的顺序是不一样的

	for key, value := range courseMap {
		fmt.Printf("key: %s, value: %s \r\n", key, value)
	}

	//	courseMap 判断是否存在 java属性
	_, ok := courseMap["java"]
	if !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in")
	}

	//	删除元素

	delete(courseMap, "grpc")
	fmt.Println(courseMap)

	//	map不是线程安全的 两个协程 对map操作的话会报错 需要使用
	//var syncMap = sync.Map{}

	//fmt.Println(syncMap)

}
