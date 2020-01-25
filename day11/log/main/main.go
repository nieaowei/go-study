/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 9:05 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Cfg().SetFileName("./config.ini")

	//g.Log().SetStack(false)
	g.Log().Debug("debug")
	g.Log().Info("info")
	g.Log().Notice("notice")
	g.Log().Warning("waring")
	g.Log().Error("error")
	g.Log().Fatal("fatal")

}
