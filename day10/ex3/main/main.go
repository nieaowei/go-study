/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/22 2:00 下午
 *  Notes       :	Template exceise.
 *******************************************************/
package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/Users/nieaowei/Library/Mobile Documents/com~apple~CloudDocs/项目代码/go-study/day10/ex3/main/index.html")
	//Name := "nieaowei"
	//person := Person{
	//	Name: "nieaowei",
	//	Age:  10,
	//}
	map1 := make(map[string]string)
	map1["age"] = "11"
	map1["name"] = "123"
	t.Execute(w, map1)
}

func main() {
	http.HandleFunc("/index", index)

	http.ListenAndServe("0.0.0.0:10000", nil)

}
