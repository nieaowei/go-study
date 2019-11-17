/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/17 2:14 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	var age0 [5]int = [5]int{1, 2, 3}
	var age1 = [5]int{1, 2, 3}
	var age2 = [...]int{1, 2, 3}
	var age3 = [...]string{1: "123312", 4: "3213"}
	fmt.Println(age0)
	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(age3)

	var array = [...][3]int{{4, 5, 6}, {1, 2, 3}}
	fmt.Println(array)
	fmt.Println(len(array))
	for key, value := range array {
		for key1, value2 := range value {
			fmt.Printf("%d %d %d \n", key, key1, value2)
		}
		fmt.Println()
	}
}
