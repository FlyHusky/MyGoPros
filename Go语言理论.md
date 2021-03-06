有关go语言的一些理论知识笔记
===
### Go语言是一种静态语言。编译器在编译时需要知道程序里每个值的类型。由此编译器可以知道要分配多少内存给值(值的规模)和内存表示什么。
### 有些类型的内部表示与编译代码的机器体系结构有关，比如go中的int，在64位系统上占用8个字节，而在32位系统上占用4个字节。

### 5.1 Go语言允许用户声明自定义的类型。声明后的类型与内置类型的运作方式类似。
### 两种方式声明用户定义的类型。1：用struct关键字来定义结构体类型。
### 2:基于已有的类型，定义一个类型，用type。比如time包内有个定义 type Duration int64 
### Duration 本质是一个int64类型，但是Go不会认定Duration和int64是同一种类型。 
var dur Duration
dur = int64(1000)  //编译器会报错 cannot use int64(1000) (type int64) as type Duration <br>
//在Go中，两种不同类型的值即使互相兼容，也不能互相赋值，编译器不会对不同类型的值做隐式转换。<br>
有关struct的演示在180813中。

### 5.2 方法。 能给用户定义的类型添加新的行为(Go有别与其它语言的一个特点吧)。 方法实际上也是函数，值是在声明时，在关键字<br>
###            func和方法名之间增加了一个类型对象。<br>
相关演示在180813中<br>

### 5.3 接口。Go语言的接口功能，书本中95页。<br>
go 语言中可以定义接口，如下<br>
	type Writer interface {<br>
	    Write(p []byte) (n int, err error)<br>
	}<br>

### 如果其他类型实现了接口中定义的方法，那么这个类型就实现了这个接口<br>
### 接口，可以作为一种类型，当做函数的参数值<br>
如 func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {<br>
实现了接口的类型值，可以做为参数代入 如上面的函数。 <br>
相关演示在180814中<br>

### 5.4 大小写字母。 若变量或函数的名的首字母是大写-包外可见。小写则为私有<br>



