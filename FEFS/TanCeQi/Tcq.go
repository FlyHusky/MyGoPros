package TanCeQi

import (
	"errors"
)

//定义探测器类型, 在go 语言中，struct 替代了Class
type Tcq struct {
	tcqName     string  //探测器名称
	tcqKind     int     //探测器类型
	tcqNet      int     //探测器所在通信支路
	tcqILBaojin float32 //探测器 电流报警值
	TcqCommSta  int     //通信状态  0=通信正常  1-tcqCommMax 表示等待  >=tcqCommMax表示通信故
	TcqCommMax  int     //通信连接失败，最大连续次数
	tcqAddr     int     //探测器的通讯地址
	tcqId       int     //探测器的id
}

//设置探测器名称,要求传入参数tcqname 长度不能大于10个字符。
//tcqName 首字母是小的，即对包外是不可见的，因此要实现set和get方法，安全啊。
func (tcq *Tcq) SetTcqName(tcqname string) error {
	if tcqname == "" {
		return errors.New("tcqname can not be empty!")
	}
	if len(tcqname) > 10 {
		return errors.New("tcqname's length over 10!")
	}
	tcq.tcqName = tcqname
	return nil
}

//返回探测器名称，即箱号。
//这个GetTcqName() 是tcq的方法，
//Go语言中方法的声明是不是很特别。
func (tcq *Tcq) GetTcqName() string {
	return tcq.tcqName
}

func (tcq *Tcq) GetTcqId() int {
	return tcq.tcqId
}

//返回一个新探测器对象,只对tcqName 做有效校验。
//tcqName=tcqname,tcqKind=1,tcqNet=1,tcqILBaojin=500.0,tcqCommSta=0,tcqCommMax=5
//tcqAddr=0,tcqId=0此两项值需要再次修改后可用。
func NewTcq(strname string, id int, addr int) (Tcq, error) {
	var tcq Tcq
	if strname == "" {
		return tcq, errors.New("tcqname can not be empty!")
	}
	if len(strname) > 10 {
		return tcq, errors.New("tcqname's length over 10!")
	}

	tcq.tcqName = strname
	tcq.tcqAddr = addr
	tcq.TcqCommMax = 5
	tcq.TcqCommSta = 0
	tcq.tcqId = id
	tcq.tcqILBaojin = 500.0
	tcq.tcqKind = 1
	tcq.tcqNet = 1
	return tcq, nil
}
