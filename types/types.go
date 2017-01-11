package types

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
	"math/rand"
)

type VelocityComponent struct {
	vel engo.Point
}

type GameOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

type MovingOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	VelocityComponent
	common.RenderComponent
}

func NewFrog() *GameOb {
	res := GameOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 50, Height: 50}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}
	res.SetZIndex(4.5)

	return &res
}

func NewCar(loc, vel engo.Point) *MovingOb {
	res := MovingOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Position: loc, Width: 100, Height: 50}
	res.VelocityComponent = VelocityComponent{vel}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{uint8(rand.Intn(255)), 0, 255, 255},
	}
	res.SetZIndex(3.5)
	return &res
}
