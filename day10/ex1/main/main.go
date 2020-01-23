/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/22 12:54 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("handler hello")
	writer.Write([]byte("hello"))
}

func login(writer http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	writer.Write([]byte(username + ",welcome."))
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	http.ListenAndServe("0.0.0.0:10000", nil)
}
