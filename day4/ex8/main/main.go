/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/17 3:19 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	slice := make([]int, 10, 10)
	slice = append(slice, 10)
	fmt.Println(len(slice), cap(slice), slice)

	//array := [...]int{1,2,3,4,5}
	var array [5]int = [...]int{1, 2, 3, 4, 5}
	slice1 := array[1:]
	slice1[1] = 100
	fmt.Println(slice1)

	slice1 = append(slice1, 10)
	slice1 = append(slice1, 10)

	fmt.Println(slice1)

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 1)
	fmt.Println(s1, s2)
	copy(s2, s1)
	fmt.Println(s1, s2)

	str := "啊hello"

	strRune := []rune(str)
	//str = "world"
	strRune[0] = 'd'
	str = string(strRune)
	fmt.Printf("%T\n", str[0])

	fmt.Println(str)

}
