package sys

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/coderconvoy/frogger/types"
	"image/color"
	"math/rand"
)

type GameOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.CollisionComponent
	common.RenderComponent
	DeathComponent
}

func NewFrog() *GameOb {
	res := GameOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 50, Height: 50}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}
	res.DeathComponent = DeathComponent{}
	res.CollisionComponent = common.CollisionComponent{Solid: false, Main: true}
	res.SetZIndex(4.5)

	return &res
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
	if fms.f.DeathComponent.DeadTime > 0 {
		return
	}

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
		vel := &v.VelocityComponent.Vel
		pos.X += vel.X * d
		pos.Y += vel.Y * d
	}
}

type CarSpawnSystem struct {
	sys   *SysList
	since float32
	level int
}
type CrashEntity struct {
	*ecs.BasicEntity
	*types.DeathComponent
	*common.CollisionComponent
	*common.RenderComponent
	*common.SpaceComponent
}

type CrashSystem struct {
	obs []CrashEntity
}

func (cs *CrashSystem) Add(be *ecs.BasicEntity, dc *types.DeathComponent, cc *common.CollisionComponent, rc *common.RenderComponent, sc *common.SpaceComponent) {
	cs.obs = append(cs.obs, CrashEntity{be, dc, cc, rc, sc})
}

func (cs *CrashSystem) AddByInterface(ob interface {
	GetBasicEntity() *ecs.BasicEntity
	GetDeathComponent() *types.DeathComponent
	GetCollisionComponent() *common.CollisionComponent
	GetRenderComponent() *common.RenderComponent
	GetSpaceComponent() *common.SpaceComponent
}) {
	cs.Add(ob.GetBasicEntity(), ob.GetDeathComponent(), ob.GetCollisionComponent(), ob.GetRenderComponent(), ob.GetSpaceComponent())
}

func (cs *CrashSystem) Remove(e ecs.BasicEntity) {
	dp := -1
	for i, v := range cs.obs {
		if v.ID() == e.ID() {
			dp = i
			break
		}
	}
	if dp >= 0 {
		cs.obs = append(cs.obs[:dp], cs.obs[dp:]...)
	}
}

func (cs *CrashSystem) Update(d float32) {

	for _, v := range cs.obs {
		if v.CollisionComponent.Collides || v.DeathComponent.DeadTime > 0 {
			v.DeathComponent.DeadTime += d
			v.RenderComponent.Color = color.RGBA{255, 0, 0, 255}
		}
		if v.DeadTime > 2 {
			v.DeadTime = 0
			v.RenderComponent.Color = color.RGBA{0, 0, 0, 255}
			v.SpaceComponent.Position = engo.Point{300, 350}
		}
	}
}
