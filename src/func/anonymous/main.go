package anonymous

import "fmt"

/*
*
1. go 支持匿名函数，如果我们某个函数只是希望使用一次，可以考虑使用匿名函数
2. 匿名函数使用方式
  - 1. 在定义匿名函数是就直接调用，这种方式匿名函数只能调用一次
  - 2. 将匿名函数赋给一个变量（该变量就是函数变量了）， 再通过该变量来调用匿名函数

3. 如何让一个匿名函数，可以在整个程序中有效呢？将匿名函数给一个全局变量就可以了
*/
func main() {
	sum := (func(num1, num2 int) int {
		return num1 + num2
	})(10, 20)
	fmt.Println(sum)

	sum1 := func(num1, num2 int) int {
		return num1 + num2
	}
	fmt.Println(sum1(10, 20))
}
