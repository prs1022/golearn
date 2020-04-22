package main

import "fmt"

func main() {
	sumer := adder()
	for i := 1; i <= 100; i++ {
		sumer(i)
	}
	fmt.Println(sumer(0))

	summ := func(a, b int) int { return a + b }(3, 5)
	fmt.Println(summ)

}

//闭包
func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
