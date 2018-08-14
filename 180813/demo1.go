//演示GO语言的Hello World 
//1:演示struct自定义类型，同时用方法给类型添加行为。


package main	//package 后面跟包的名字。 如果当前包里面要写main函数，那么包名必须为main

import(			//采用improt 关键字导入程序将会使用到的包。 fmt是go的标准表，有关控制台输入输出的。
	"fmt"		//go语言是一门严禁的语言，不允许程序存在没有使用到的包和变量。呵呵，go目前还有调试器呢！ go写好后，不用调试的。
	"errors"
)
//go语言中第1中定义用户类型的样子，基于已有的内置数据类型。
type salary int64 //定义一个工资类型。基础类型是int64

//说说我的工资多不多，工资大于15k,才算工资高
func(sa salary)saySalary(){
	if sa >=15000 {
		fmt.Println("工资水平处于上游哦！")
	}else if sa>=10000{
		fmt.Println("工资水平处于中游！")
	}else{
		fmt.Println("工资水平处于下游！")
	}
}



//go语言中第2中定义用户类型的方法用 struct
//演示自定义结构体类型。 type 和 struct 是关键字;user 是类型的名称
type user struct{
	name	string
	email 	string
	ext		 int
	privileged bool
}

//用方法给 自定义类型添加新的行为。
//下面演示修改对象的邮箱地址，当然了可以直接用user.email=email，可以不用搞个函数。
func(u user)changeEmail(email string){
	u.email=email
}

//对changeEmail 升级，不要那么简单,对传入的参数做校验。若传入值不合法，则返回error
func(u user)changeEmail2(email string) error { //u user 值接收

	//判断传入的email 长度是否合法
	if len(email) >50 || len(email) <6 {
		return errors.New("the length or the email is over limited !")
	}else{
		u.email = email
	}
	return nil
}

//对changeEmail 升级，不要那么简单,对传入的参数做校验。若传入值不合法，则返回error
func(u *user)changeEmail3(email string) error { //u *user 指针接收

	//判断传入的email 长度是否合法
	if len(email) >50 || len(email) <6 {
		return errors.New("the length or the email is over limited !")
	}else{
		u.email = email
	}
	return nil
}


func main(){	//func 关键字 用于定义函数。
	fmt.Println("Hello Welcome back to Golang world!")	//go是简洁的，每行语句结束，不用加分号。
	Show_Struct()
}

//演示结构体类型的声明和初始化赋值
func Show_Struct(){
	//声明一个lisa对象，并初始化所有字段。
	lisa := user{
		name: "lisa",
		email: "lisa@email.com",
		ext: 123,
		privileged: true,
	}

	//在声明一个willion对象，并初始化所有字段。
	willion :=user{"willion","willion@email.com",101,false}

	fmt.Println(lisa)
	fmt.Println(willion.email)

	//修改lisa 的邮箱 1
	err:=lisa.changeEmail2("lisa@163.com")
	if err==nil {
		fmt.Println("change lisa's email successfully!")
	}else{
		fmt.Println("change lisa's email fail!")
	}
	fmt.Printf("lisa's email is %s \n",lisa.email)  //打印出来还是，lisa@email.com 而不是 lisa@163.com
	//因为changeEmail2的方法 用的是值接收(u user),调用时会使用这个值的一个副本来执行。

	//修改lisa 的邮箱2 
	err1:=lisa.changeEmail3("lisa@qq.com")	//指针接收的方法
	if err1==nil {
		fmt.Println("change lisa's email successfully!")
	}else{
		fmt.Println("change lisa's email fail!")
	}
	fmt.Printf("lisa's email is %s \n",lisa.email)  //打印出来是 lisa@163.com

	var sa salary = 16000
	sa.saySalary()

	var sb salary = 9000
	sb.saySalary()


}