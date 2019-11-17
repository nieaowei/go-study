/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/13 9:00 下午
 *  Notes       :
 *******************************************************/
package main

import "fmt"

func T1() {
	fmt.Println("----------------T1-----------------")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d  ", i, j, i*j)

		}
		fmt.Println()
	}
	//for i := 1; i <= 9; i++ {
	//	for j := i; j <= 9; j++ {
	//		fmt.Printf("%d x %d = %d \n", i, j, i*j)
	//	}
	//	fmt.Println()
	//}
}

func IsWan(n int64) (res bool) {
	resNum := int64(0)
	if n == 0 {
		return
	}
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			resNum += i
		}
	}
	return n == resNum
}

func T2() {
	fmt.Println("----------------T2-----------------")
	for i := int64(0); i < 1000; i++ {
		if IsWan(i) {
			fmt.Printf("[%d] is wan.\r\n", i)
		}
	}
}

func IsBack(str string) (res bool) {
	resStr := ""
	for _, value := range str {
		resStr += string(value)
	}
	return resStr == str
}

func T3() {
	fmt.Println("----------------T3-----------------")
	str := "中国中"
	fmt.Printf("[%s] is [%v]\r\n", str, IsBack(str))
}

func Total(str string) {
	ziNum, space, number, other, chinese := 0, 0, 0, 0, 0

	for _, value := range str {
		//fmt.Println(n)
		switch true {
		case (value > 'a' && value < 'z') || (value > 'A' && value < 'Z'):
			ziNum++
		case value == ' ':
			space++
		case value > '0' && value < '9':
			number++
		case len(string(value)) == 3:
			chinese++
		default:
			other++
		}
	}
	fmt.Println("ziNum:", ziNum)
	fmt.Println("space:", space)
	fmt.Println("number:", number)
	fmt.Println("chinese:", chinese)
	fmt.Println("other:", other)
	fmt.Println("all:", ziNum+space+number+chinese+other)
}

func T4() {
	fmt.Println("----------------T4----------------")
	str := "daodwhlwd ojlasjfl/;''123';adl apeifepo doh 到时间到的da da']=0%$#$%^&*ekalf"
	Total(str)
}

func T5() {
	fmt.Println("-----------------T5----------------")

}

func main() {
	T1()
	T2()
	T3()
	T4()
	T5()
}
