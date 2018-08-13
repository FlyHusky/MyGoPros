//演示GO语言的Hello World 

package main	//package 后面跟包的名字。 如果当前包里面要写main函数，那么包名必须为main

import(			//采用improt 关键字导入程序将会使用到的包。 fmt是go的标准表，有关控制台输入输出的。
	"fmt"		//go语言是一门严禁的语言，不允许程序存在没有使用到的包和变量。呵呵，go目前还有调试器呢！ go写好后，不用调试的。
)

func main(){	//func 关键字 用于定义函数。
	fmt.Println("Hello Welcome back to Golang world!")	//go是简洁的，每行语句结束，不用加分号。
}