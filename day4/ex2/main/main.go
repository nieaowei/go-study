/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/16 7:39 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	s1 := new([]int)

	//s1 := append(s1,1,2)
	fmt.Printf("%T\n", s1)
	*s1 = append(*s1, 1, 2)
	fmt.Println(s1)

	s2 := make([]int, 10)
	s2 = append(s2, 1, 2)
	fmt.Printf("%T\n", s2)
	fmt.Println(s2)
}
