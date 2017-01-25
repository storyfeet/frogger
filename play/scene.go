package play

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"image/color"
)

type SysList struct {
	Render    *common.RenderSystem
	FrogMove  *FrogMoveSystem
	CarSpawn  *CarSpawnSystem
	ObMove    *ObMoveSystem
	CollSys   *common.CollisionSystem
	CrashSys  *CrashSystem
	BoundsSys *BoundsDeathSystem
}

type MainScene struct{}

func (*MainScene) Type() string { return "MainScene" }
func (*MainScene) Preload()     {}
func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	engo.Input.RegisterButton("left", engo.ArrowLeft)
	//engo.Input.RegisterButton("left", engo.ArrowLeft)
	engo.Input.RegisterButton("right", engo.ArrowRight)
	engo.Input.RegisterButton("up", engo.ArrowUp)
	engo.Input.RegisterButton("down", engo.ArrowDown)

	var sList SysList

	fg := NewFrog(engo.Point{300, 350})
	a := fg.GetBasicEntity()
	fmt.Println(a.ID())

	sList.Render = &common.RenderSystem{}
	sList.FrogMove = NewFrogMoveSystem(fg)
	sList.ObMove = &ObMoveSystem{}
	sList.CarSpawn = NewCarSpawnSystem(1, &sList)
	sList.CollSys = &common.CollisionSystem{}
	sList.CrashSys = &CrashSystem{}
	sList.BoundsSys = &BoundsDeathSystem{rect: engo.AABB{engo.Point{-5, -5}, engo.Point{610, 410}}, w: w}

	sList.Render.AddByInterface(fg)
	sList.CollSys.AddByInterface(fg)
	sList.CrashSys.Add(fg)

	w.AddSystem(sList.Render)
	w.AddSystem(sList.CollSys)
	w.AddSystem(sList.FrogMove)
	w.AddSystem(sList.ObMove)
	w.AddSystem(sList.CarSpawn)
	w.AddSystem(sList.CrashSys)
	w.AddSystem(sList.BoundsSys)

}
