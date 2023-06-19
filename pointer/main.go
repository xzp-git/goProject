package main

import "fmt"

type Person struct {
	name string
}

func main() {
	//	指针需要初始化

	//第一种初始化方式
	personPointer1 := &Person{"fofo"}

	fmt.Println(personPointer1)

	//	第二种初始化方式
	var emptyPerson = Person{}

	personPointer2 := &emptyPerson

	fmt.Println(personPointer2)

	//	第三种初始化方式

	var personPointer3 = new(Person)

	fmt.Println(personPointer3)

	/**
	初始化 map channel slice 初始化推荐使用make方式

	指针初始化推荐使用new 函数， 指针要初始化否则要出现 nil Pointer


	map 必须初始化
	*/

	//通过 swap 交换指针的值

	a, b := 1, 2

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	var k []Person

	if k == nil {
		fmt.Println(len(k), k, "k is nil")
	}

	/**

	不同类型的数据零值不一样
	bool false
	numbers 0
	string ""
	pointer nil
	slice nil
	map  nil
	channel、interface、 function nil
	struct默认值不是nil、默认值是具体字段的默认值

	*/
}

// p Person 结构体中 定义方法 结构体一般用首字母的小写
//func (p Person) Say() {
//
//}

func swap(a, b *int) {
	//取指针a 所指向的值
	t := *a

	*a = *b

	*b = t

}
