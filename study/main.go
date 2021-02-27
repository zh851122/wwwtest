package main

import "fmt"

func init() {
	var i, j, k = 1, 2, 3
	fmt.Println("init hello world")
	fmt.Println(i, j, k)

}

func main() {
	//未使用变量,不允许声明
	fmt.Println("hello world")
	//
}
