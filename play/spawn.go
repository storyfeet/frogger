package play

import (
	"math/rand"

	"engo.io/ecs"
	"engo.io/engo"
)

type CarFactory func(float32) (DriveFace, bool)

type RowFactory func(r int) CarFactory

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

func BasicRowFactory(base float32) RowFactory {
	return func(r int) CarFactory {
		m2 := float32(r % 2)
		v := float32(40 + (5 * r))
		return BasicCarFactory(
			engo.Point{-100 + m2*700, base - float32((r+1)*50)},
			engo.Point{(1 - m2*2) * v, 0},
			4, 0, 1000, 1000,
		)

	}
}

type CarSpawnSystem struct {
	sys    *SysList
	rows   []CarFactory
	rowfac RowFactory
	rownum int
}

func NewCarSpawnSystem(level int, sysList *SysList, rowfac RowFactory) *CarSpawnSystem {
	rows := []CarFactory{}
	for i := 0; i < 8; i++ {
		newFac := rowfac(i)
		rows = append(rows, newFac)
	}

	return &CarSpawnSystem{
		sys:    sysList,
		rows:   rows,
		rowfac: rowfac,
		rownum: 7,
	}
}

func (css *CarSpawnSystem) Fill() {
	for _, fac := range css.rows {
		cars := fac.Init()

		for _, c := range cars {
			css.AddCar(c)
		}
	}
}

func (css *CarSpawnSystem) New(w *ecs.World) {
	engo.Mailbox.Listen("RowMessage", func(message engo.Message) {
		css.rownum++
		nFac := css.rowfac(css.rownum)
		for _, c := range nFac.Init() {
			css.AddCar(c)
		}
		css.rows = append(css.rows, nFac)
		css.rows = css.rows[1:]
	})
	engo.Mailbox.Listen("ResetMessage", func(message engo.Message) {
		rows := []CarFactory{}
		for i := 0; i < 8; i++ {
			newFac := css.rowfac(i)
			rows = append(rows, newFac)
		}
		css.rownum = 7
		css.rows = rows
		css.Fill()
	})
}

//Remove Spawner has no need to remove stuff for now
//Somewhere along the line, I think the factories will become components to be added, isn't that fun.
func (*CarSpawnSystem) Remove(e ecs.BasicEntity) {}

//Update cycle through factories and see if they have a car to makek
func (css *CarSpawnSystem) Update(d float32) {

	for _, fac := range css.rows {
		c, ok := fac(d)
		if ok {
			css.AddCar(c)
		}
	}
}

func (css *CarSpawnSystem) AddCar(c DriveFace) {
	css.sys.Render.AddByInterface(c)
	css.sys.ObMove.AddByInterface(c)
	css.sys.CollSys.AddByInterface(c)
	css.sys.BoundsSys.AddByInterface(c)
}
