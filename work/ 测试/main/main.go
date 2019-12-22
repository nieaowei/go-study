/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/19 10:44 上午
 *  Notes       :
 *******************************************************/
package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		year  int
		month int
		date  int
	)
	n, err := fmt.Scanf("%d %d %d", &year, &month, &date)
	// input fail
	if err != nil {
		return
	}
	// some var input fail
	if n != 3 {
		return
	}
	_, _, _, err = JudgetDate(year, month, date, 1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func JudgetDate(year, month, date, option, mode int) (int, int, int, error) {
	if year < 1900 || year > 2100 {
		year = -1
	}
	//闰年
	flagYear := false
	if year%4 == 0 && year%100 != 0 {
		flagYear = true
	}
	//非闰年
	if month > 12 || month < 1 {
		return year, month, date, errors.New("月份错误")
	}
	if date < 1 {
		return year, month, date, errors.New("月份错误")

	}
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if date > 31 {
			return year, month, date, errors.New("月份错误")
		}
	case 2:
		if flagYear && date > 29 {
			return year, month, date, errors.New("月份错误")
		} else if date > 28 {
			return year, month, date, errors.New("月份错误")
		}
	default:
		if date > 30 {
			return year, month, date, errors.New("月份错误")
		}
	}
	fmt.Println("您输入的日期:", year, month, date)

	if mode == 1 {
		//option := 0
		fmt.Printf("1.前一天\n2.后一天\n3.前一周\n4.后一周\n")
		fmt.Println("输入您的选择:")
		//fmt.Scanf("%d", &option)
		switch option {
		case 1:
			if _, _, _, err := JudgetDate(year, month, date-1, 0, 0); err != nil {
				fmt.Println(err)
				return year, month, date, err
			}
			fmt.Println(year, month, date-1)
		case 2:
			if _, _, _, err := JudgetDate(year, month, date+1, 0, 0); err != nil {
				fmt.Println(err)
				return year, month, date, err
			}
			fmt.Println(year, month, date+1)
		case 3:
			if _, _, _, err := JudgetDate(year, month, date-7, 0, 0); err != nil {
				fmt.Println(err)
				return year, month, date, err
			}
			fmt.Println(year, month, date-7)
		case 4:
			if _, _, _, err := JudgetDate(year, month, date+7, 0, 0); err != nil {
				fmt.Println(err)
				return year, month, date, err
			}
			fmt.Println(year, month, date+7)
		}
	}
	return year, month, date, nil
}
