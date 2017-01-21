package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"github.com/coderconvoy/frogger/types"
	"image/color"
)

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

	fg := types.NewFrog()
	a := fg.GetBasicEntity()
	fmt.Println(a.ID())

	sList.Render = &common.RenderSystem{}
	sList.FrogMove = NewFrogMoveSystem(fg)
	sList.ObMove = &ObMoveSystem{}
	sList.CarSpawn = NewCarSpawnSystem(1, &sList)
	sList.CollSys = &common.CollisionSystem{}
	sList.CrashSys = &CrashSystem{}

	sList.Render.AddByInterface(fg)
	sList.CollSys.AddByInterface(fg)
	sList.CrashSys.AddByInterface(fg)

	w.AddSystem(sList.Render)
	w.AddSystem(sList.CollSys)
	w.AddSystem(sList.FrogMove)
	w.AddSystem(sList.ObMove)
	w.AddSystem(sList.CarSpawn)
	w.AddSystem(sList.CrashSys)

}

func main() {
	opts := engo.RunOptions{
		Title:  "Frogger",
		Width:  600,
		Height: 400,
	}
	engo.Run(opts, &MainScene{})
}
