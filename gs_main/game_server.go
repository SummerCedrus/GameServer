package main

import (
	."GameServer/stage"
	"fmt"
	"github.com/SummerCedrus/ServerKit/misc"
	. "github.com/SummerCedrus/ServerKit/netkit"
	"time"
)

func main(){
	misc.InitLog("run", "server")
	//hotplugin.Run()
	_, err := NewServer("127.0.0.1:8080")

	if nil != err{
		fmt.Errorf("Create New Server Error [%s]", err.Error())
	}
	//hotplugin.Call("testplugin","Hello")
	//mainWork(mgr)
	InitStage()
	for user_id:= int32(1000);user_id<1005;user_id++{
		r := UserPoolMgr.NewRoutine()
		UserMap[user_id] = r.Rid
	}

	for user_id:= int32(1000);user_id<1005;user_id++{
		rid := UserMap[user_id]
		fmt.Println("rid ",rid)
		r := UserPoolMgr.GetRoutine(rid)
		r.RecvQueue <- new(Message)
		time.Sleep(2*time.Second)
		r.CloseChan <- true
		time.Sleep(2*time.Second)
	}
	UserPoolMgr.CloseChan <- true
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

