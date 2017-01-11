package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/coderconvoy/frogger/types"
	"math/rand"
)

type SysList struct {
	Render   *common.RenderSystem
	FrogMove *FrogMoveSystem
	CarSpawn *CarSpawnSystem
	ObMove   *ObMoveSystem
}

var sysList SysList

type FrogMoveSystem struct {
	f *types.GameOb
}

func NewFrogMoveSystem(f *types.GameOb) *FrogMoveSystem {
	return &FrogMoveSystem{f}
}

func (fms *FrogMoveSystem) Update(d float32) {
	pos := &fms.f.SpaceComponent.Position
	if engo.Input.Button("left").JustPressed() {
		pos.X -= 25
	}

	if engo.Input.Button("right").JustPressed() {
		pos.X += 25
	}
	if engo.Input.Button("up").JustPressed() {
		pos.Y -= 25
	}
	if engo.Input.Button("down").JustPressed() {
		pos.Y += 25
	}
}

func (*FrogMoveSystem) Remove(e ecs.BasicEntity) {
}

type movable struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*types.VelocityComponent
}
type ObMoveSystem struct {
	obs []movable
}

func (oms *ObMoveSystem) Add(a *ecs.BasicEntity, b *common.SpaceComponent, c *types.VelocityComponent) {
	oms.obs = append(oms.obs, movable{a, b, c})
}
func (oms *ObMoveSystem) Remove(e ecs.BasicEntity) {
	dp := -1
	for i, v := range oms.obs {
		if v.BasicEntity.ID() == e.ID() {
			dp = i
			break
		}
	}
	if dp >= 0 {
		oms.obs = append(oms.obs[:dp], oms.obs[dp:]...)
	}
}
func (oms *ObMoveSystem) Update(d float32) {
	for _, v := range oms.obs {
		pos := &v.SpaceComponent.Position
		pos.X += 10 * d
	}
}

type CarSpawnSystem struct {
	sys   *SysList
	since float32
	level int
}

func NewCarSpawnSystem(level int, sysList *SysList) *CarSpawnSystem {
	return &CarSpawnSystem{
		sys:   sysList,
		since: 0,
		level: level,
	}
}

func (*CarSpawnSystem) Remove(e ecs.BasicEntity) {}
func (css *CarSpawnSystem) Update(d float32) {
	css.since += d
	if rand.Float32()*50 < css.since*float32(css.level+3) {
		c := types.NewCar(engo.Point{rand.Float32() * 500, rand.Float32() * 300},
			engo.Point{10, 0})
		css.sys.Render.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
		css.sys.ObMove.Add(&c.BasicEntity, &c.SpaceComponent, &c.VelocityComponent)
		css.since = 0
	}

}
