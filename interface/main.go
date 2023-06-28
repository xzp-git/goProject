package main

import "fmt"

type Duck interface {
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (p *pskDuck) Gaga() {
	fmt.Println("Gaga")
}
func (p *pskDuck) Walk() {
	fmt.Println("Walk")
}
func (p *pskDuck) Swimming() {
	fmt.Println("Swimming")
}

// 鸭子类型强调的是事物的外部行为，而不是内部的结构
func main() {

	//go 语言中处处都是 inteface，到处都是我interface

	var psd Duck = &pskDuck{}
	psd.Gaga()
}
