package main

import "fmt"

func main() {
	/**
	type关键字
	1.定义结构体
	2.定义接口
	3.定义类型别名
	4.类型定义
	5.类型判断
	*/

	var a byte = 'A'

	fmt.Printf("%T", a)

	//	结构体

	type Person struct {
		name    string
		age     uint8
		address string
		height  float32
	}

	p1 := Person{"fofo", 23, "北京", 1.80} // 这种初始化方式 必须全部赋值
	//p2 := Person{
	//	name:    "bobo",
	//	age:     19,
	//	address: "北京",
	//	height:  1.89,
	//} // 这种方式更灵活

	//初始化 与 赋值
	var persons []Person
	persons = append(persons, p1)
	persons = append(persons, Person{
		name: "bobby3",
	})
	//persons2 := []Person{
	//	{name: "bobby1", age: 18, address: "wfwef", height: 1.80},
	//	{
	//		age: 19,
	//	},
	//}

	//匿名结构体  取值
	var p Person
	p.age = 20
	fmt.Println(p.height)
	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		province: "北京市",
		city:     "通州区",
		address:  "xxx"}
	fmt.Println(address.city)

	//s := Student{
	//	p: Person{
	//		name: "bobby", age: 18,
	//	},
	//	score: 95.6,
	//}

	s := Student{
		Person1: Person1{name: "ssss", age: 233},
		score:   95.6,
		name:    "fwefwefw",
	}
	s.printss()

	s2 := Student{
		Person1{name: "ssss", age: 23},
		95.6,
		"fofofofo", //里面和外面同时 有一样的字段时  外面的优先级高
	}
	fmt.Println(s.age)
	fmt.Println(s2.age)
	fmt.Println(s2.name) // fofofof

}

type Person1 struct {
	name string
	age  uint8
}
type Student struct {
	//第一种嵌套方式
	//p     Person
	Person1 //匿名写法
	score   float32
	name    string
}

// 接收器的形态有两种
// 指针形态是引用传递  非指针形态是值传递
func (p *Student) printss() {
	//有可能该函数中想要改p的值,就是person对象很大,数据较大的时候考虑使用指针方式

	p.age = 19
	fmt.Printf("name:%s, age:%d", p.name, p.age)
}
