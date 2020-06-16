package routines
import (
	"github.com/SummerCedrus/ServerKit/misc"
	."github.com/SummerCedrus/ServerKit/netkit"
)

type RoutinePool struct {
	pool map[int32]*routine
}
var max_rid = int32(1)

func (rp *RoutinePool) NewRoutine() *routine{
	r := routine{
		Rid:max_rid,
		SendQueue:make(chan *Message),
		RecvQueue:make(chan *Message),
		CloseChan:make(chan bool),
	}
	rp.AddRoutine(&r)

	max_rid++

	return &r
}

func (rp *RoutinePool) AddRoutine(r *routine){
	if _,ok := rp.pool[r.Rid];ok{
		misc.Errorf("AddRoutine failed dumplcate Rid %v",r.Rid)
		return
	}

	rp.pool[r.Rid] = r
}

func (rp *RoutinePool) DelRoutine(r *routine){
	if _,ok := rp.pool[r.Rid];ok{
		misc.Errorf("AddRoutine failed dumplcate Rid %v",r.Rid)
		return
	}

	rp.pool[r.Rid] = r
}