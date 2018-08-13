package main

import(
	"os"
	"strings"
	"runtime"
	"path"
	log "github.com/sirupsen/logrus"
	"fmt"
)

//logrus 中 Hook 接口
//type Hook interface { 
//    Levels() []Level 
//    Fire(*Entry) error
//}

//定义了一个空的结构体，for logrus
type ContextHook struct{}

// Levels returns the whole levels.
func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

//在用golang的Logrus库的时候，用Json格式记录日志，有时候希望加一些固定的Fields,比如在日志收集中心中，Log记录自己的程序名，
//又不想每次Log日志的时候，每处都手动添加到logrus.Fields中。
//这时可以使用logrus的Hook来完成这个功能。每此写入日志时拦截，修改logrus.Entry。

// Fire helps logrus record the related file, function and line.
//哎，看不懂哎。 传入参数 *log.Entry 返回error
func (hook ContextHook) Fire(entry *log.Entry) error {
	pc := make([]uintptr, 3, 3) //定义一个长度和容量都为3的指针。
	cnt := runtime.Callers(6, pc)
	fmt.Println(" fire ")

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		fmt.Println("name=",name)
		if !strings.Contains(name, "github.com/Sirupsen/log") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

//init 函数会在main函数执行前调用。
func init(){
	//设置logrus 日志格式
	log.Info("init ")
	//设置log 的格式，  文本格式，完整的时间
	log.SetFormatter(&log.TextFormatter{FullTimestamp:true,DisableColors:true})
log.Info("init 2")
	//判断用户是否设置了BYTOM_DEBUG的路径
	if os.Getenv("BYTOM_DEBUG")!= ""{
		fmt.Println("bytom_DEBUG")
		log.AddHook(ContextHook{})
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	//cmd := cli.PrepareBaseCmd(commands.RootCmd, "TM", os.ExpandEnv(config.DefaultDataDir()))
	//cmd.Execute()
	log.Info("Hello, here is logrus!")
 
	log.Info("just log info")
	Learn_runtime()

 


}


//runtime 是go语言中重要的包，包含了Go运行时的系统交互操作，比如goruntine,debug，重要的就是调度器和GC(垃圾回收)。
//go语言之所以号称高并发，在语言级上支持并发，就是因为go有自己的调度器。很强大的，不深入研究，现在先学习下runtime
//包内的基础知识，以后再深入。
func Learn_runtime(){
		fmt.Printf("本电脑逻辑CPU的数量是:%d \n",runtime.NumCPU())
		//runtime.Gosched()   停止当前的goroutine， 让给其它的goroutine
		fmt.Printf("当前存在的goroutine的数量是:%d \n",runtime.NumGoroutine())

		//获取当前go的版本
		fmt.Printf("Go version is %s \n",runtime.Version())

		//
}