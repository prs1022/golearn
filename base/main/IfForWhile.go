package main

import (
	"fmt"
)

func main() {
	//ifWhile()

	switchCase2()
}

func switchCase2() {
	var a int
	for true {
		fmt.Scan(&a)
		switch a {
		case 1, 2:
			fmt.Println("输入的1,2中一个")
		case 3, 4:
			fmt.Println("输入了3或4")
		}
	}
}

func switchCase() {
	var a interface{}
	switch i := a.(type) {
	case nil:
		fmt.Println("a type:%T", i)
	case string:
		fmt.Println("a type:%T", i)
	case interface{}:
		fmt.Println("a type %T", i)
	default:

	}
	switch a {
	case a:
		fmt.Println("2")
		fallthrough
	case false:
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("4")
	}
}

func ifWhile() {
	var a string
	fmt.Println("请输入密码\n[exit]退出")
	fmt.Scan(&a)
	for true {
		if a != "exit" {
			if a != "123" {
				fmt.Println("密码错误")
			} else {
				fmt.Println("居然被你猜到了")
				break
			}
			fmt.Scan(&a)
		} else {
			fmt.Println("是不是玩不起")
			break
		}
	}
}
