//本文件演示Go语言中的接口功能，同pro1一样，无非是自己写个例子。
// 同时演示了嵌入类型

package main

import (
	"fmt"
)

//定义一个接口
type Notifier interface {
	Notify() //内有一个Notify()函数
}

type User struct {
	name  string
	email string
}

//User 实现 Notifier 接口
func (usr *User) Notify() {
	fmt.Printf("正在发邮件给%s<%s>\n", usr.name, usr.email)
}

type Admin struct {
	name  string
	email string
	phone string
}

func (ad Admin) Notify() {
	fmt.Printf("正在发邮件给%s<%s>\n", ad.name, ad.email)
}

//演示嵌入类型
type Admin2 struct{
	User	//嵌入了结构体类型user， 只用写类型名即可
	phone string	//额外增加了一个phone字段
}


//声明一个方法，接收参数值为notifier接口类型
func SendNotification(n Notifier) { //这个函数也实现了多态，
	fmt.Println("发邮件！")
	n.Notify() //依据 n 值类型的不同，调用了不同的Notify 函数。
}

func main() {
	raj := User{"raj", "raj@163.com"}
	sheldon := Admin{"sheldon", "sheldon@gmail.com", "16868686688"}
	SendNotification(&raj)
	SendNotification(&sheldon)

	//嵌入类型-  内部类型的提升
	penny := Admin2{ User{"penny","penny@qq.com"},"16816881888"}
	penny.User.Notify()
	penny.Notify()	//user是Admin2的内部类型，这里被自动提升到了外部。哈哈。
	SendNotification(&penny) //因为自动提升到了外部，所以Admin2也实现noitifer的接口。哈哈
	//当然了，如果Admin2自己在外部实现notifier接口，这内部user的不会被自动提示为外部，代码部在演示看书113页。
}
