/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/23 8:16 下午
 *  Notes       :	Tail -f example.
 *******************************************************/
package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	tailFileName := "./my.log"
	tailFile, err := tail.TailFile(tailFileName, tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	})
	if err != nil {
		fmt.Println("tail file err")
		return
	}
	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tailFile.Lines
		if !ok {
			fmt.Println("tail file close reopen,filename:", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
