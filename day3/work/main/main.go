/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/4 9:04 下午
 *  Notes       :	Three exercises.
*******************************************************/
package main

import "fmt"

//IsPrime is to determine whether num is a prime number.
func IsPrime(num int) (res bool) {
	// invalid
	if !(num > 1) {
		return
	}
	// valid
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return
		}
	}
	return true
}

//IsDaffodil3n is to determine whether nun is a daffodli number.
func IsDaffodil3n(num int) (res bool) {
	// invalid
	if num < 100 && num > 999 {
		return
	}
	//valid
	x0 := num % 100 % 10
	x1 := (num % 100) / 10
	x2 := num / 100

	x0 *= x0 * x0
	x1 *= x1 * x1
	x2 *= x2 * x2

	if num != (x0 + x1 + x2) {
		return
	}
	return true
}

//Factorial is used to find the factorial of num.
func Factorial(num int) (res int) {
	// invalid
	if num < 0 {
		return
	}
	// valid
	for res = 1; num != 0; num-- {
		res *= num
	}
	return
}

func SumN(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += Factorial(i)
	}
	return
}

//Ex1 is exercise 1.
func Ex1() {
	fmt.Println("---------------------------")
	for i := 101; i < 200; i++ {
		if IsPrime(i) {
			fmt.Println(i, ":", IsPrime(i))
		}
	}
}

//Ex2 is exercise 2.
func Ex2() {
	fmt.Println("---------------------------")
	for i := 100; i <= 999; i++ {
		if IsDaffodil3n(i) {
			fmt.Println(i, ":", IsDaffodil3n(i))
		}
	}
}

//Ex3 is exercise 3.
func Ex3() {
	fmt.Println("---------------------------")
	for i := 1; i < 10; i++ {
		fmt.Println(i, ":", SumN(i))
	}
}

func main() {
	Ex1()
	Ex2()
	Ex3()
}
