package main

import (
	"fmt"
	"github.com/SummerCedrus/ServerKit/hotplugin"
	"github.com/SummerCedrus/ServerKit/misc"
	."github.com/SummerCedrus/ServerKit/netkit"

)

func main(){
	misc.InitLog("run", "server")
	//hotplugin.Run()
	mgr, err := NewServer("127.0.0.1:8080")

	if nil != err{
		fmt.Errorf("Create New Server Error [%s]", err.Error())
	}
	hotplugin.Call("testplugin","Hello")
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

