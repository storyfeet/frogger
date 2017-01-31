package play

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
)

type CarFactory func(float32) (DriveFace, bool)

func (cf CarFactory) Init() []DriveFace {

	res := []DriveFace{}
	var i float32
	for i = 0; i < 25; i += 0.05 {
		c, ok := cf(0.05)

		if ok {
			sc := c.GetSpaceComponent()
			vc := c.GetVelocityComponent()
			sc.Position.X += vc.Vel.X * (25 - i)
			sc.Position.Y += vc.Vel.Y * (25 - i)
			res = append(res, c)
		}
	}
	return res

}

func BasicCarFactory(pos, vel engo.Point, wait, n, r, lt float32) CarFactory {
	var since float32 = 0
	return func(d float32) (DriveFace, bool) {
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
		newFac := BasicCarFactory(
			engo.Point{-100 + m2*700, float32((i + 1) * 50)},
			engo.Point{(1 - m2*2) * v, 0},
			4, 0, 1000, 1000,
		)
		rows = append(rows, newFac)
	}

	return &CarSpawnSystem{
		sys:  sysList,
		rows: rows,
	}
}

func (css *CarSpawnSystem) Fill() {
	for _, fac := range css.rows {
		cars := fac.Init()

		for _, c := range cars {
			css.sys.Render.AddByInterface(c)
			css.sys.ObMove.AddByInterface(c)
			css.sys.CollSys.AddByInterface(c)
			css.sys.BoundsSys.AddByInterface(c)
		}
	}
}

//Remove Spawner has no need to remove stuff for now
//Somewhere along the line, I think the factories will become components to be added, isn't that fun.
func (*CarSpawnSystem) Remove(e ecs.BasicEntity) {}

//Update cycle through factories and see if they have a car to makek
func (css *CarSpawnSystem) Update(d float32) {

	for _, fac := range css.rows {
		c, ok := fac(d)
		if ok {
			css.sys.Render.AddByInterface(c)
			css.sys.ObMove.AddByInterface(c)
			css.sys.CollSys.AddByInterface(c)
			css.sys.BoundsSys.AddByInterface(c)
		}
	}

}
