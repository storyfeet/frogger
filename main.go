package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type MainScene struct{}

func (*MainScene) Type() string { return "MainScene" }
func (*MainScene) Preload()     {}

func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	engo.Input.RegisterButton("left", engo.ArrowLeft)
	engo.Input.RegisterButton("right", engo.ArrowRight)
	engo.Input.RegisterButton("up", engo.ArrowUp)
	engo.Input.RegisterButton("down", engo.ArrowDown)

	fg := NewFrog()

	rs := &common.RenderSystem{}
	fms := NewFrogMoveSystem(fg)

	w.AddSystem(rs)
	w.AddSystem(fms)

	rs.Add(&fg.BasicEntity, &fg.RenderComponent, &fg.SpaceComponent)
}

func main() {
	opts := engo.RunOptions{
		Title:  "Frogger",
		Width:  600,
		Height: 400,
	}
	engo.Run(opts, &MainScene{})
}
