/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func modufy(a int) {
	a = 10
	return
}

func modify1(a *int) {
	*a = 10
}

func main() {
	a := 5
	b := make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	modufy(a)
	fmt.Println("a=", a)

	modify1(&a)
	fmt.Println("a=", a)

}
