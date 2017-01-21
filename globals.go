package main

import (
	"github.com/coderconvoy/frogger/sys"
)

type SysList struct {
	Render   *common.RenderSystem
	FrogMove *sys.FrogMoveSystem
	CarSpawn *sys.CarSpawnSystem
	ObMove   *sys.ObMoveSystem
	CollSys  *common.CollisionSystem
	CrashSys *sys.CrashSystem
}
