/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/5 8:39 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func main() {
	str := "dsadwqnnkasdkhajjf"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%d,", str[i])
	}

}
