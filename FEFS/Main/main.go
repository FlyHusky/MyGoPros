package main

import (
	"fmt"
	"github.com/goinaction/code/FEFS/TanCeQi"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//全局变量声明

var (
	TCQ_TOTAL_COUNT int           = 100 //系统中探测器的总数量，此值要从数据库读取。
	Tcqs            []TanCeQi.Tcq       //探测器对象 切片， 一开始是个空的。
)

//探测器通讯状态类型 for 通道
type tcqCommChain struct {
	tcqId  int  //探测器的id
	commOk bool //true=表示通讯正常  false=通讯失败
}

//定义等待线程结束 wg
var wg sync.WaitGroup

//init 函数会在 main函数调用前执行，用于初始化
func init() {
	//1-数据库读取信息
	//2-初始化Tcqs切片
	tcqs_init()

	//3- 随机函数的种子

	rand.Seed(time.Now().UnixNano())
}

// 探测器对象 切片初始化
func tcqs_init() {
	name := "ALE"
	for i := 1; i <= TCQ_TOTAL_COUNT; i++ {
		s := strconv.Itoa(i)
		s = s + name
		tcq, err := TanCeQi.NewTcq(s, i, i)
		if err != nil {
			fmt.Printf("创建Tcq对象出错，信息:%s", err)
			os.Exit(1)
		}
		Tcqs = append(Tcqs, tcq) //调用append方法，会动态扩容Tcqs切片
	}
}

func main() {
	fmt.Println("FEFS 电气火灾监控系统已经启动！！！！")

	for _, tcq := range Tcqs {
		fmt.Println(tcq)
	}

	//定义一个探测器通讯状态的通道 缓冲10个
	commWorkChain := make(chan tcqCommChain, 10)

	wg.Add(1)
	go commWorkCenter(commWorkChain)
	//模拟串口通讯，每100ms，完成了一次和终端探测器的通讯任务。

	commLoopId := 1 //通讯循环标从 1开始 到 TCQ_TOTAL_COUNT
	fmt.Println("模拟1000次通讯，约100秒后结束！")
	//模拟1000次循环和终端通讯，月100后，程序自动退出
	for tt := 1; tt < 1000; tt++ {
		//采用随机函数，来模拟通讯是否正常，还是失败。
		if commLoopId > TCQ_TOTAL_COUNT {
			commLoopId = 1
		}

		rnd := rand.Intn(100)
		commsta := false
		if rnd > 50 { //如果随机数大于50，则通讯正常
			commsta = true
		}

		commchain := tcqCommChain{commLoopId, commsta}

		commWorkChain <- commchain
		//fmt.Printf("%d ",tt)
		commLoopId = commLoopId + 1
		time.Sleep(100 * time.Millisecond) //延时100ms

	}

	close(commWorkChain)

	wg.Wait()

	fmt.Println("FEFS电气火灾监控系统关闭！")

}

//通讯事务处理中心，
//接收一个缓冲的通道，通道内包含一个探测器的通讯信息。
func commWorkCenter(comminfo chan tcqCommChain) {

	defer wg.Done()

	//死循环
	for {
		comm, ok := <-comminfo
		if !ok { //通道已经空了，并且已经关闭了
			fmt.Printf("commWorkCenter shutting Down\n")
			return
		}

		//如果通讯正常
		if comm.commOk {
			//通讯从故障恢复到正常
			if Tcqs[comm.tcqId-1].TcqCommSta >= Tcqs[comm.tcqId-1].TcqCommMax {
				fmt.Printf("\n %s 通讯恢复正常！ \n", Tcqs[comm.tcqId-1].GetTcqName())
			}

			Tcqs[comm.tcqId-1].TcqCommSta = 0 //标记通讯正常

		} else {

			if Tcqs[comm.tcqId-1].TcqCommSta >= Tcqs[comm.tcqId-1].TcqCommMax {
				continue
			}

			Tcqs[comm.tcqId-1].TcqCommSta = Tcqs[comm.tcqId-1].TcqCommSta + 1

			if Tcqs[comm.tcqId-1].TcqCommSta >= Tcqs[comm.tcqId-1].TcqCommMax {
				Tcqs[comm.tcqId-1].TcqCommSta = Tcqs[comm.tcqId-1].TcqCommMax + 1
				fmt.Printf("\n %s 通讯故障！ \n", Tcqs[comm.tcqId-1].GetTcqName())
			}
		}
	}

}
