//本文演示了Go语言的接口功能，对于书本中95页。
//go 语言中可以定义接口，如下
	//type Writer interface {
	//	Write(p []byte) (n int, err error)
	//}

//如果其他类型实现了接口中定义的方法，那么这个类型就实现了这个接口

//接口，可以作为一种类型，当做函数的参数值，
//如 func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {

//实现了接口的类型值，可以做为参数代入 如上面的函数。v

package main

import(
	"fmt"
	"bytes"	//bytes包提供了操作[]byte 常用函数。
	"io"
	"os"
)

func main(){

	var buf  bytes.Buffer
	
	//将字符串写入buffer
	buf.Write([]byte("Hello"))

	//使用Fprintf将字符串拼接到buffer
	fmt.Fprintf(&buf,"Wordl!");     //这里用的取地址符号&buf, 因为方法的接收是指针类型，

	//下面是fprintf 函数的定义。第一个参数是io.Writer
	/*
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
		p := newPrinter()
		p.doPrintf(format, a)
		n, err = w.Write(p.buf)  //这里用了接口write的方法。
		p.free()
		return
	}*/

	//一个bytes.Buffer 的值，为何可以传入Fprintf(io.Writer )里面呢？？？？

	/* //在 io.go 中找到了 Writer ,是个接口。	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	//在buffer.go 找到了 Write 函数，即实现了io.Writer接口。因此buffer可以做为参数赋值给 Fprintf(w io.Writer, 
	func (b *Buffer) Write(p []byte) (n int, err error) {
		b.lastRead = opInvalid
		m, ok := b.tryGrowByReslice(len(p))
		if !ok {
			m = b.grow(len(p))
		}
		return copy(b.buf[m:], p), nil
	}

	*/
	
	//将buffer的内容写到stdout
	io.Copy(os.Stdout,&buf)

}