package play

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
)

type CarSpawnSystem struct {
	sys  *SysList
	rows []func(dt float32) (Drivable, bool)
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

		c := NewCar(engo.Point{x, float32(row * 50)},
			engo.Point{speed, 0})
		css.sys.Render.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
		css.sys.ObMove.Add(&c.BasicEntity, &c.SpaceComponent, &c.VelocityComponent)
		css.sys.CollSys.AddByInterface(c)
		css.sys.BoundsSys.AddByInterface(c)
		css.since = 0
	}

}
