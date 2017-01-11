package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

type Car struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}
