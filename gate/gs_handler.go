package gate

import (
	."GameServer/protocol"
	"fmt"
	. "github.com/SummerCedrus/ServerKit/netkit"
	"net"
)

var (
	GsSessionMap map[int32]*net.TCPConn
)
func MsgHander(msg *Message){
	switch msg.Cmd {
	case GS_CONNECT_GATE_REQ:
		fmt.Println(msg.Msg.String())
	default:
		fmt.Errorf("uncatch cmd %v", msg.Cmd)
	}
}