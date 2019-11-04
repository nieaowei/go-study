/*******************************************************
 *  ProjectName :go-stduy
 *  Author      :nieaowei
 *  Date        :2019/11/3
 *  Notes       :
 *******************************************************/
package add

import (
	"fmt"
	_ "go_study/day2/ex2/test"
)

var Name string = "xxx"
var Age int = 123

func Test() {
	Name = "hellword"
}

func init() {
	fmt.Println("add init.")
}
