package routines

import (
	"fmt"
	"github.com/SummerCedrus/ServerKit/misc"
	. "github.com/SummerCedrus/ServerKit/netkit"
)

const MAX_CHAN_LEN = 20
type RoutinePool struct {
	pool map[int32]*routine
	DelChan chan int32	//删除进入channel的ID
	CloseChan chan bool
}
var max_rid = int32(1)

func NewRoutinePool() *RoutinePool {
	rp := RoutinePool{
		pool: make(map[int32]*routine, 0),
		DelChan:make(chan int32,MAX_CHAN_LEN),
		CloseChan:make(chan bool, 1),
	}

	return &rp
}

func (rp *RoutinePool) Run(){
	defer func() {
		close(rp.DelChan)
		close(rp.CloseChan)
	}()
	for{
		select{
		case rid, ok := <- rp.DelChan:
			if ok {
				rp.delRoutine(rid)
			}

		case cls,ok := <- rp.CloseChan:
			if cls && ok{
				fmt.Println("RoutinePool Close ", cls)
				return
			}
		}
	}
}
func (rp *RoutinePool) NewRoutine() *routine{
	r := routine{
		Rid:max_rid,
		SendQueue:make(chan *Message, MAX_CHAN_LEN),
		RecvQueue:make(chan *Message, MAX_CHAN_LEN),
		CloseChan:make(chan bool, 1),
		PoolChan:rp.DelChan,
	}
	rp.addRoutine(&r)
	r.run()
	max_rid++

	return &r
}

func (rp *RoutinePool) addRoutine(r *routine){
	fmt.Println("addRoutine ",r.Rid)
	if _,ok := rp.pool[r.Rid];ok{
		misc.Errorf("AddRoutine failed dumplcate Rid %v",r.Rid)
		return
	}

	rp.pool[r.Rid] = r
}

func (rp *RoutinePool) delRoutine(rid int32){
	if _,ok := rp.pool[rid];!ok{
		misc.Errorf("DelRoutine failed cant find Rid %v",rid)
		return
	}
	delete(rp.pool, rid)
}

func (rp *RoutinePool) GetRoutine (rid int32) *routine {
	if r,ok := rp.pool[rid];ok{
		return r
	}

	misc.Errorf("GetRoutine failed valid Rid %v",rid)

	return  nil
}