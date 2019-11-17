/*******************************************************
 *  File :   main.go
 *  Author      :   nieaowei
 *  Date        :   2019/11/5 8:48 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func urlHandle(url string) (res string) {
	re := strings.HasPrefix(url, "http://")
	if !re {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func pathHandle(path string) (res string) {
	re := strings.HasSuffix(path, "/")
	if !re {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}

func main() {
	//var (
	//	url string
	//	path string
	//)
	//fmt.Scanf("%s",&url)
	//fmt.Scanf("%s",&path)
	//fmt.Println(urlHandle(url))
	//fmt.Println(pathHandle(path))

	str := "absdasdaabasDsdabdd"

	fmt.Println(strings.Replace(str, "ab", "cc", -1))

	fmt.Println(strings.Count(str, "dd"))

	fmt.Println(strings.Repeat(str, 2))

	fmt.Println(strings.ToLower(str))

	//fmt.Println(strings.ToUpper(str))

	fmt.Println(strings.ToTitle(str))

	fmt.Println(strings.TrimSuffix("sasd d adsa", "a"))

	fmt.Println(strings.Fields("asd sad asd aswq"))

	fmt.Println(strings.Split("asddqdqwdadasdd", "d"))

	fmt.Println(len(strings.Split("asddqdqwdadasdd", "d")))

	fmt.Println(len(strings.Split("asddqdqwdadasdd", "d")))

	str2 := strings.Split("asddqweqdasdqweqsda", "d")

	fmt.Println(strings.Join(str2, ","))

	fmt.Println(strconv.Atoi("123333"))

	fmt.Println(strconv.Itoa(12334433))

}
