package routines

import (
	"fmt"
	"github.com/SummerCedrus/ServerKit/misc"
	."github.com/SummerCedrus/ServerKit/netkit"
)

const(
	MAX_RETRY_NUM = 20
)
type routine struct {
	Rid		  int32		   //routine id
	SendQueue chan *Message//消息发送队列
	RecvQueue chan *Message//消息接收队列
	CloseChan chan bool			//
	PoolChan  chan int32
}


func (r *routine) run(){
	fmt.Println("run routine ",r.Rid)
	go r.work()
}
func (r *routine) handle() error{
	for{
		select {
		case _,ok := <- r.RecvQueue:
			if ok{
				fmt.Println("recv...")
			}
		case _,ok := <- r.SendQueue:
			if ok{
				fmt.Println("send...")
			}
		case cls,ok := <- r.CloseChan:
			if ok && cls{
				fmt.Println("close...")
				return nil
			}
		}
	}
}
func (r *routine) work(){
	defer r.close()
	for retryCnt := 0;retryCnt<MAX_RETRY_NUM;{
		err := r.handle()
		if nil == err{
			break
		}
		misc.Errorf("routine err rid [%v] err[%v] retrycnt [%v]",r.Rid, err.Error(), retryCnt)
		retryCnt++
	}
}

func (r *routine) close(){
	close(r.SendQueue)
	close(r.RecvQueue)
	close(r.CloseChan)
	r.PoolChan <- r.Rid
	fmt.Println("close routine ",r.Rid)
}