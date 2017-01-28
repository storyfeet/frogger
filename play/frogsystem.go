package play

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Frog struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
	DeathComponent
	JumpComponent
}

func NewFrog(loc engo.Point) *Frog {
	res := Frog{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 50, Height: 50}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}
	res.DeathComponent = DeathComponent{}
	res.CollisionComponent = common.CollisionComponent{Solid: false, Main: true, Extra: engo.Point{-3, -3}}
	res.SetZIndex(2.5)
	res.Reset(loc)

	return &res
}

func (fg *Frog) Reset(pt engo.Point) {
	fg.SpaceComponent.Position = pt
	fg.JumpComponent.Target = pt
	fg.JumpComponent.Next = pt
	fg.RenderComponent.Color = color.Black
}

var sysList SysList

type FrogMoveSystem struct {
	f *Frog
}

func NewFrogMoveSystem(f *Frog) *FrogMoveSystem {
	return &FrogMoveSystem{f}
}

func (fms *FrogMoveSystem) Update(d float32) {
	jc := &fms.f.JumpComponent
	jcT := &jc.Target
	jcN := &jc.Next

	pos := &fms.f.SpaceComponent.Position
	if fms.f.DeathComponent.DeadTime == 0 {
		rel := engo.Point{0, 0}
		if engo.Input.Button("left").JustPressed() {
			rel = engo.Point{-40, 0}
		}
		if engo.Input.Button("right").JustPressed() {
			rel = engo.Point{40, 0}
		}
		if engo.Input.Button("up").JustPressed() {
			rel = engo.Point{0, -50}
		}
		if engo.Input.Button("down").JustPressed() {
			rel = engo.Point{0, 50}
		}

		np := engo.Point{0, 0}
		if rel != np {
			jcN.X = jcT.X + rel.X
			jcN.Y = jcT.Y + rel.Y
		}

		if *pos == *jcT {
			*jcT = *jcN
		}
	}

	if pos.X < jcT.X {
		pos.X += d * 140
		if pos.X > jcT.X {
			pos.X = jcT.X
		}
	}

	if pos.X > jcT.X {
		pos.X -= d * 140
		if pos.X < jcT.X {
			pos.X = jcT.X
		}
	}
	if pos.Y < jcT.Y {
		pos.Y += d * 140
		if pos.Y > jcT.Y {
			pos.Y = jcT.Y
		}
	}

	if pos.Y > jcT.Y {
		pos.Y -= d * 140
		if pos.Y < jcT.Y {
			pos.Y = jcT.Y
		}
	}
}

func (*FrogMoveSystem) Remove(e ecs.BasicEntity) {
}

type CrashEntity struct {
	*ecs.BasicEntity
	*DeathComponent
	*common.CollisionComponent
	*common.RenderComponent
	*common.SpaceComponent
}

type CrashSystem struct {
	obs []*Frog
}

func (cs *CrashSystem) Add(f *Frog) {
	cs.obs = append(cs.obs, f)
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
			v.Reset(engo.Point{300, 350})
			v.DeadTime = 0
		}
	}
}
