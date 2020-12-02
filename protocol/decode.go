package protocol

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func ReflectMessage(Cmd uint32) (proto.Message, error) {
	switch Cmd {
	case GS_CONNECT_GATE_REQ:
		return &GsConnectInfo{}, nil
	case GS_CONNECT_GATE_ACK:
		return &NullMsg{}, nil
	default:
		fmt.Printf("Error Cmd [%d]", Cmd)
		return nil, errors.New("Error Cmd")
	}
}
