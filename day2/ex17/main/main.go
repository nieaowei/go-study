/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/4 5:26 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func reverse(str string) (result string) {
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result += fmt.Sprintf("%c", str[strlen-i-1])
	}
	return
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	for i := 0; i < len(tmp); i++ {
		result = append(result, tmp[len(tmp)-i-1])
	}
	return string(result)
}

func main() {
	var str1 = "hello"
	var str2 = "world"

	str3 := str1 + " " + str2

	str3 = fmt.Sprintf("%s %s", str1, str2)

	fmt.Println(str3)

	fmt.Printf("len:%v\n", len(str3))

	str4 := str3[3:]
	fmt.Println(str3[1:])

	fmt.Println(str4)

	fmt.Println(reverse(str3))

	fmt.Println(reverse1(str3))
}
