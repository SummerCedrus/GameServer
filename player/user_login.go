package player

import (
	."GameServer/protocol"
	"fmt"
)

func LoginOK(msg ClientInfo){
	fmt.Printf("%v", msg.String())
}