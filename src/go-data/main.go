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
