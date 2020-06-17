package stage

import "GameServer/core/routines"

var (
	UserPoolMgr *routines.RoutinePool

	UserMap map[int32]int32 //userID和routineID映射
 )

func InitStage()  {
	UserPoolMgr = routines.NewRoutinePool()

	go UserPoolMgr.Run()

	UserMap = make(map[int32]int32, 0)
}
