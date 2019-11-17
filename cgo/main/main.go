/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/12 8:49 下午
 *  Notes       :
 *******************************************************/
package main

// #include "main.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(1, 10))
}
