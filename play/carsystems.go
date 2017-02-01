package play

import (
	"fmt"
	"image/color"
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type MovingOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	VelocityComponent
	common.RenderComponent
	common.CollisionComponent
}

func NewCar(loc, vel engo.Point) *MovingOb {
	res := MovingOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Position: loc, Width: 100, Height: 50}
	res.VelocityComponent = VelocityComponent{vel}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{uint8(rand.Intn(255)), 0, 255, 255},
	}
	res.CollisionComponent = common.CollisionComponent{Solid: true, Main: false}
	res.SetZIndex(3.5)
	return &res
}

//Bounds Death System, for killing cars out of bounds

type BoundsDeathSystem struct {
	rect engo.AABB
	w    *ecs.World
	obs  []SpaceEntity
}

func (bds *BoundsDeathSystem) Add(be *ecs.BasicEntity, sc *common.SpaceComponent) {
	bds.obs = append(bds.obs, SpaceEntity{be, sc})
}

func (bds *BoundsDeathSystem) AddByInterface(ob interface {
	Spaceable
	ECSBasicable
}) {
	be := ob.GetBasicEntity()
	if be == nil {
		fmt.Printf("Log No Entity")
		return
	}
	bds.Add(ob.GetBasicEntity(), ob.GetSpaceComponent())
}

func (bds *BoundsDeathSystem) New(w *ecs.World) {
	engo.Mailbox.Listen("RowMessage", func(message engo.Message) {
		bds.rect.Min.Y -= 50
	})
	engo.Mailbox.Listen("ResetMessage", func(message engo.Message) {
		t := bds.obs
		for _, v := range t {
			bds.w.RemoveEntity(*v.BasicEntity)
		}
	})
}

func (bds *BoundsDeathSystem) Update(d float32) {
	t := bds.obs
	for _, v := range t {
		sc := v.SpaceComponent
		if sc.Position.X+sc.Width < bds.rect.Min.X ||
			sc.Position.X > bds.rect.Max.X ||
			sc.Position.Y+sc.Height < bds.rect.Min.Y ||
			sc.Position.Y > bds.rect.Max.Y {

			bds.w.RemoveEntity(*v.BasicEntity)
		}
	}
}

func (bds *BoundsDeathSystem) Remove(ob ecs.BasicEntity) {

	bds.obs = RemoveSpaceEntity(bds.obs, ob.ID())
}

type movable struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*VelocityComponent
}

type ObMoveSystem struct {
	obs []movable
}

func (oms *ObMoveSystem) Add(a *ecs.BasicEntity, b *common.SpaceComponent, c *VelocityComponent) {
	oms.obs = append(oms.obs, movable{a, b, c})
}

func (obs *ObMoveSystem) AddByInterface(ob interface {
	ECSBasicable
	Spaceable
	Velocitable
}) {
	obs.Add(ob.GetBasicEntity(), ob.GetSpaceComponent(), ob.GetVelocityComponent())
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
		oms.obs = append(oms.obs[:dp], oms.obs[dp+1:]...)
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
