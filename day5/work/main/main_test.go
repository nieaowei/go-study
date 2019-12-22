/*******************************************************
 *  File        :   main_test.go.go
 *  Author      :   nieaowei
 *  Date        :   2019/12/21 8:30 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortBubble(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"冒泡排序",
			args{
				[]int{1, 4, 7, 2, 3, 5, 10, 9},
			},
			[]int{10, 9, 7, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortBubble(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("res:%v want:%v", tt.args.array, tt.want)
			}
		})
		fmt.Println(tt.args.array)
	}

}

func TestSortSelect(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"选择排序",
			args{
				[]int{1, 4, 7, 2, 3, 5, 10, 9},
			},
			[]int{1, 2, 3, 4, 5, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortSelect(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("res:%v want:%v", tt.args.array, tt.want)
			}
		})
		fmt.Println(tt.args.array)
	}
}

func TestSortInsert(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"插入排序",
			args{
				[]int{1, 4, 7, 2, 3, 5, 10, 9},
			},
			[]int{1, 2, 3, 4, 5, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortInsert(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("res:%v want:%v", tt.args.array, tt.want)
			}
		})
		fmt.Println(tt.args.array)
	}
}

func TestSortFast(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"快速排序",
			args{
				[]int{7, 4, 1, 2, 3, 5, 10, 9},
			},
			[]int{1, 2, 3, 4, 5, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortFast(tt.args.array)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("res:%v want:%v", tt.args.array, tt.want)
			}
		})
		fmt.Println(tt.args.array)
	}
}
