package main

import "fmt"

type Person struct {
	name string
}

/**

基本数据类型的内存
  var age int = 18
		age -> 18 在内存中为age变量开辟一个空间，这个空中储存的数值位18，这个空间本身也是有地址的
*/

func main() {

	var age = 18
	// &符号 + 变量 就可以获取这个变量内存的地址
	fmt.Println(&age) //0x14000122008

	// 定义一个指针变量
	// var 代表要声明一个变量
	// ptr 指针变量的名字
	// ptr对应的类型是：*int 是一个指针类型 （可以理解为 指向int类型的指针）
	// &age 就是一个地址 是ptr变量的具体的值
	// ptr -> 指向age的地址 0x14000122008
	var ptr *int = &age
	fmt.Println(ptr)                      // 0x14000122008
	fmt.Println("ptr本身这个存储空间的地址为：", &ptr) //  0x140000ae020
	// 总结：指针就是内存地址

	// 想获取ptr这个指针或者这个地址指向的那个数据
	fmt.Printf("ptr指向的数值为：%v \n", *ptr)

	// 总结：最重要的就是两个字符
	// 1. &取内存地址
	// 2. *根据地址取值

	// 指针的细节
	// 1. 可以通过指针改变指向值
	num := 10
	fmt.Println(num)

	var ptr1 = &num
	*ptr1 = 20
	fmt.Println(num)

	// 2. 指针变量接收的一定是地址值

	// vat ptr1 *int = num 这个代码是错的

	// 3. 指针变量的地址不可以不匹配

	// vat ptr1 *float32 = &num 这个代码是错的

	// 4. 基本数据类型（又叫值类型），都有对应的指针类型，形式为 *数据类型，
	//比如 int的 对应的指针就是 *int，float32对应的指针类型就是 *float32。依次类推

	//	指针初始化方式
	//第一种初始化方式
	//personPointer1 := &Person{"fofo"}
	//
	//fmt.Println(personPointer1)

	//	第二种初始化方式
	//var emptyPerson = Person{}
	//
	//personPointer2 := &emptyPerson
	//
	//fmt.Println(personPointer2)

	//	第三种初始化方式

	//var personPointer3 = new(Person)
	//
	//fmt.Println(personPointer3)

	/**
	初始化 map channel slice 初始化推荐使用make方式

	指针初始化推荐使用new 函数， 指针要初始化否则要出现 nil Pointer

	map 必须初始化
	*/

	//通过 swap 交换指针的值
	//
	//a, b := 1, 2
	//
	//fmt.Println(a, b)
	//swap(&a, &b)
	//fmt.Println(a, b)
	//
	//var k []Person
	//
	//if k == nil {
	//	fmt.Println(len(k), k, "k is nil")
	//}

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
