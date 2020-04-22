package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 我第一个go 程序
func main() {
	fmt.Println("hello world")
	rand.Seed(2)
	fmt.Println("生成个随机数:", rand.Intn(10))

	fmt.Println(math.Pi)

	// go的类型不同于C family ,好处看这篇 https://blog.go-zh.org/gos-declaration-syntax
	var num int = 1
	fmt.Println(num)
	//函数
	x := 1
	y := 2
	r1, r2 := add(x, y)
	fmt.Printf("%d+%d=%d,差的绝对值:%f\n", x, y, r1, r2)
	v1, v2 := "hello", "world"
	fmt.Println(swap(v1, v2))

	//变量,类型可以从值推断，两种声明方式，初始化
	var temp = "in main"
	fmt.Println(temp+"\t"+outTemp+"\t", outTemp2, outTemp3)
	fmt.Println(zeroNum, zeroStr, zeroBool)

	v := 42.001 // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	//常量
	const game = "炉石传说"
	const play = true
	const money = 10.0

	//操作符
	fmt.Println("求余5%3=", 5%3)
	x++
	fmt.Println("自增x++", x)
	x--
	fmt.Println("自减x--", x)
	//关系运算符
	fmt.Println("x==y", x == y)
	fmt.Println("x!=y", x != y)
	//返回变量存储地址
	fmt.Println("x存储地址", &x)
	var point *int = &x
	fmt.Println("指针point所指向的值", *point)

}

var (
	outTemp2 = true
	outTemp3 = 1 << 10
	outTemp  = "out main"
	//零值
	zeroNum  int
	zeroStr  = ""
	zeroBool bool
)

func add(x, y int) (int, float64) {
}

func swap(v1, v2 string) (a string) {
	a = v2 + v1
	return
}
