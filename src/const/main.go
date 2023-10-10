package main

import "fmt"

func main() {
	const (
		//常量， 定义的时候就指定的值，不能修改， 常量尽量全部大写
		PI  float32 = 3.1415926 //显式定义
		PI2 float32 = 3.12323234
	)

	const (
		UNKNOWN = 0
		FEMALE
		MALE
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
	*/
}
