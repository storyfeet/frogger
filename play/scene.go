package play

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	Render     *common.RenderSystem
	FrogMove   *FrogMoveSystem
	CarSpawn   *CarSpawnSystem
	ObMove     *ObMoveSystem
	CollSys    *common.CollisionSystem
	CrashSys   *CrashSystem
	BoundsSys  *BoundsDeathSystem
	ClimberSys *ClimberSystem
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

	engo.Input.RegisterButton("2left", engo.A)
	//engo.Input.RegisterButton("left", engo.ArrowLeft)
	engo.Input.RegisterButton("2right", engo.D)
	engo.Input.RegisterButton("2up", engo.W)
	engo.Input.RegisterButton("2down", engo.S)

	var sList SysList

	fg1 := NewFrog(engo.Point{200, 350}, FrogCommands(0))
	fg2 := NewFrog(engo.Point{400, 350}, FrogCommands(1))

	sList.Render = &common.RenderSystem{}
	sList.FrogMove = &FrogMoveSystem{}
	sList.ObMove = &ObMoveSystem{}
	sList.CarSpawn = NewCarSpawnSystem(1, &sList, BasicRowFactory(350))
	sList.CollSys = &common.CollisionSystem{}
	sList.CrashSys = &CrashSystem{}
	sList.BoundsSys = &BoundsDeathSystem{rect: engo.AABB{engo.Point{-5, -200}, engo.Point{610, 410}}, w: w}
	sList.ClimberSys = NewClimberSystem(400, 50)

	sList.FrogMove.Add(fg1)
	sList.Render.AddByInterface(fg1)
	sList.CollSys.AddByInterface(fg1)
	sList.CrashSys.Add(fg1)
	sList.ClimberSys.AddByInterface(fg1)
	sList.Render.AddByInterface(fg2)
	sList.CollSys.AddByInterface(fg2)
	sList.FrogMove.Add(fg2)
	sList.CrashSys.Add(fg2)
	sList.ClimberSys.AddByInterface(fg2)

	w.AddSystem(sList.Render)
	w.AddSystem(sList.CollSys)
	w.AddSystem(sList.FrogMove)
	w.AddSystem(sList.ObMove)
	w.AddSystem(sList.CarSpawn)
	w.AddSystem(sList.CrashSys)
	w.AddSystem(sList.BoundsSys)
	w.AddSystem(sList.ClimberSys)

	sList.CarSpawn.Fill()
}
