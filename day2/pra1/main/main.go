/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :
 *******************************************************/
package main

import (
	"fmt"
	"os"
)

func main() {
	goos := os.Getenv("GOOS")
	fmt.Println("goos,", goos)
	path := os.Getenv("PATH")
	fmt.Println("path", path)
}
