package play

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
)

type CarFactory func(float32) (Drivable, bool)

func BasicCarFactory(pos, vel engo.Point, wait, n, r, lt float32) CarFactory {
	var since float32 = 0
	return func(d float32) (Drivable, bool) {
		since += d

		if since < wait {
			return nil, false
		}

		if since-wait+n+rand.Float32()*r > lt {
			since = 0
			return NewCar(pos, vel), true
		}
		return nil, false

	}
}

type CarSpawnSystem struct {
	sys  *SysList
	rows []CarFactory
}

func NewCarSpawnSystem(level int, sysList *SysList) *CarSpawnSystem {
	rows := []CarFactory{}
	for i := 0; i < 6; i++ {
		m2 := float32(i % 2)
		v := float32(50 - (3 * i))
		rows = append(rows, BasicCarFactory(
			engo.Point{-100 + m2*700, float32((i + 1) * 50)},
			engo.Point{(1 - m2*2) * v, 0},
			4, 0, 1000, 1000,
		))
	}

	return &CarSpawnSystem{
		sys:  sysList,
		rows: rows,
	}
}

//Remove Spawner has no need to remove stuff for now
//Somewhere along the line, I think the factories will become components to be added, isn't that fun.
func (*CarSpawnSystem) Remove(e ecs.BasicEntity) {}

//Update cycle through factories and see if they have a car to makek
func (css *CarSpawnSystem) Update(d float32) {

	for _, v := range css.rows {
		c, ok := v(d)
		if ok {
			css.sys.Render.AddByInterface(c)
			css.sys.ObMove.AddByInterface(c)
			css.sys.CollSys.AddByInterface(c)
			css.sys.BoundsSys.AddByInterface(c)
		}
	}

}
