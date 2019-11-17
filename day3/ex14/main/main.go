/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 8:44 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

//variable parameber
func add(a int, argv ...int) (res int) {
	res = a
	for i := 0; i < len(argv); i++ {
		res += argv[i]
	}
	return
}

func addStr(str string, argv ...string) (res string) {
	res = str
	for i := 0; i < len(argv); i++ {
		res += argv[i]
	}
	return
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 10))
	fmt.Println(addStr("ad", "dsad", "ds", "a"))
	a()

	//匿名函数1
	c := func(a int, b int) int {
		return a + b
	}
	//匿名函数2

	b := func(a int, b int) int {
		return a + b
	}(1, 4)

	fmt.Println(c)
	fmt.Println(c(1, 2))

	fmt.Println(b)
}

func a() {
	i := 0
	defer fmt.Println(i)
	defer fmt.Println("12321")
	i++
}
