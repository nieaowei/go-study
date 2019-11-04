/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func test() {
	var a = 100
	fmt.Println(a)
	for i := 0; i < 100; i++ {
		var b = i * 2
		fmt.Println("b=", b)
	}
}

func test2() {
	var a int8 = 100
	var b int16 = 32
	fmt.Println(a, b)
}

func main() {
	a, b := 1, 2
	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)

	a, b = swap1(a, b)
	fmt.Println("a=", a, "b=", b)

	test()

	test2()

}
