package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type Frog struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func NewFrog() *Frog {
	res := Frog{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 100, Height: 100}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}

	return &res
}

type FrogMoveSystem struct {
	f *Frog
}

func NewFrogMoveSystem(f *Frog) *FrogMoveSystem {
	return &FrogMoveSystem{f}
}

func (fms *FrogMoveSystem) Update(d float32) {
	pos := &fms.f.SpaceComponent.Position
	if engo.Input.Button("left").JustPressed() {
		pos.X -= 10
	}

	if engo.Input.Button("right").JustPressed() {
		pos.X += 10
	}
	if engo.Input.Button("up").JustPressed() {
		pos.Y -= 10
	}
	if engo.Input.Button("down").JustPressed() {
		pos.Y += 10
	}
	pos.X += d * 20
}

func (*FrogMoveSystem) Remove(e ecs.BasicEntity) {
}
