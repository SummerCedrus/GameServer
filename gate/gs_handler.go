package gate

import (
	."GameServer/protocol"
	"fmt"
	. "github.com/SummerCedrus/ServerKit/netkit"
)

func MsgHander(msg *Message){
	switch msg.Cmd {
	case LOGIN_ACK:
		fmt.Println(msg.Msg.String())
	}
}