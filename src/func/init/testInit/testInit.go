package testInit

import "fmt"

var Age int
var Name string

func init() {
	Age = 20
	Name = "我是谁"
	fmt.Println("testInit包的init函数")
}
