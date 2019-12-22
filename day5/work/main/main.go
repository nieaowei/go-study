/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/22 7:56 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

//SortBubble is 冒泡排序
func SortBubble(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

//SortSelect is 选择排序
func SortSelect(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

//SortInsert is 插入排序
func SortInsert(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] > array[j-1] {
				break
			}
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

func SortFast(array []int) {
	if len(array) == 0 {
		return
	}
	key := array[0]
	i, j := 0, len(array)-1
	for i < j {
		if array[j] < key {
			array[j], array[i] = array[i], array[j]
			for i < j {
				if array[i] > key {
					array[j], array[i] = array[i], array[j]
				} else {
					i++
				}
			}
		} else {
			j--
		}
	}
	SortFast(array[:i])
	SortFast(array[i+1:])
}

//
func main() {
	array := [4]int{3, 2, 13, 24}
	array1 := array[:3]
	//copy(array1,array[:3])
	fmt.Println(array)
	fmt.Println(array1)
	array1[2] = 12
	fmt.Println(array)
	fmt.Println(array1)
}
