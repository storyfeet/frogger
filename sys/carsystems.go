package sys

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/coderconvoy/frogger/types"
	"image/color"
	"math/rand"
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
		row := rand.Intn(6) + 1
		speed := float32((15 - row) * 5)
		var x float32 = -100
		if row%2 == 0 {
			speed = -speed
			x = 600
		}

		c := types.NewCar(engo.Point{x, float32(row * 50)},
			engo.Point{speed, 0})
		css.sys.Render.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
		css.sys.ObMove.Add(&c.BasicEntity, &c.SpaceComponent, &c.VelocityComponent)
		css.sys.CollSys.AddByInterface(c)
		css.since = 0
	}

}
