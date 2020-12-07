package main

import (
	."GameServer/stage"
	"fmt"
	"github.com/SummerCedrus/ServerKit/misc"
	. "github.com/SummerCedrus/ServerKit/netkit"
	."GameServer/protocol"
	"github.com/SummerCedrus/ServerKit/protocol"
	"github.com/SummerCedrus/ServerKit/timer"
)

var gsTimer *timer.Timer
func main(){
	misc.InitLog("run", "gs1")
	//hotplugin.Run()
	mgr, err := NewServer("127.0.0.1:4001", protocol.ReflectMessage)

	if nil != err{
		fmt.Errorf("Create New Server Error [%s]", err.Error())
	}
	//hotplugin.Call("testplugin","Hello")
	InitStage()

	//for user_id:= int32(1000);user_id<1005;user_id++{
	//	r := UserPoolMgr.NewRoutine()
	//	UserMap[user_id] = r.Rid
	//}
	//
	//for user_id:= int32(1000);user_id<1005;user_id++{
	//	rid := UserMap[user_id]
	//	fmt.Println("rid ",rid)
	//	r := UserPoolMgr.GetRoutine(rid)
	//	r.RecvQueue <- new(Message)
	//	time.Sleep(2*time.Second)
	//	r.CloseChan <- true
	//	time.Sleep(2*time.Second)
	//}
	//UserPoolMgr.CloseChan <- true
	ConnectGate()
	gsTimer = timer.NewTimer()
	gsTimer.Run()
	mainWork(mgr)
}

func msgHandle(mgr *ConnectMgr) (err interface{}){
	for {
		select {
		case msg, ok := <-mgr.MsgChan:
			if ok {
				if nil != err {
					continue
				}
				fmt.Println(msg.Cmd)
				fmt.Println(msg.Msg.String())
			}
		case _, ok := <-mgr.ConnectChan:
			if ok {
				continue
			}
		case cls, ok := <-mgr.CloseChan:
			if ok && cls {
				ShowDown()
				return
			}

			return 0
		}
	}
}

func mainWork(mgr *ConnectMgr){
	for retryCnt := 0; retryCnt < MAX_CONNECT_RETRY_TIME;{
		err := msgHandle(mgr)
		if nil == err{
			break
		}else{
			retryCnt ++
		}
	}

	ShowDown()
}

func ShowDown(){
	fmt.Println("Server ShutDown!")
}

func ConnectGate(){
	sender := NewClient("127.0.0.1:8080", "")
	info := GsConnectInfo{
		GsId:1,
	}
	sender.Send(&Message{
		Cmd:GS_CONNECT_GATE_REQ,
		Msg:&info,
	})
}

