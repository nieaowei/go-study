/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/17 3:36 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"sort"
)

func main() {

	slice := []int{2, 5, 1, 4, 7, 10}
	sliceStr := []string{"ads", "dswe", "bsad"}
	sliceDouble := []float64{23.1, 1, 23.021}

	fmt.Printf("%T", slice)
	fmt.Printf("%T\n", sliceStr)

	sort.Ints(slice)
	sort.Strings(sliceStr)
	sort.Float64s(sliceDouble)

	fmt.Println(slice)
	fmt.Println(sliceStr)
	fmt.Println(sliceDouble)
}
