/*******************************************************
 *  File        :   main.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/18 9:17 下午
 *  Notes       :
 *******************************************************/
package main

import (
	"encoding/binary"
	"fmt"
)

func BigEndian() { // 大端序
	// 二进制形式：0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 十六进制表示
	fmt.Printf("%d use big endian: \n", testInt)
	fmt.Println(uint32(testInt))
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt)) //大端序模式
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes) //大端序模式的字节转为int32
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func LittleEndian() { // 小端序
	//二进制形式： 0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 16进制
	fmt.Printf("%d use little endian: \n", testInt)
	fmt.Println(uint32(testInt))
	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt)) //小端序模式
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes) //小端序模式的字节转换
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	BigEndian()
	LittleEndian()
}
