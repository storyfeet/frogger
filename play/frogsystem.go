package play

import (
	"fmt"
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

func NewFrog(loc engo.Point, commands []KeyCommand) *Frog {
	res := Frog{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 50, Height: 50}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}
	res.DeathComponent = DeathComponent{Origin: loc}
	res.CollisionComponent = common.CollisionComponent{Solid: false, Main: true, Extra: engo.Point{-3, -3}}
	res.SetZIndex(2.5)
	res.JumpComponent.Commands = commands
	res.Reset()

	return &res
}

func (fg *Frog) Reset() {
	fg.SpaceComponent.Position = fg.Origin
	fg.JumpComponent.Target = fg.Origin
	fg.JumpComponent.Next = engo.Point{0, 0}
	fg.RenderComponent.Color = color.Black
	fg.DeadTime = 0
}

var sysList SysList

type FrogMoveSystem struct {
	frogs []*Frog
}

func FrogCommands(n int) []KeyCommand {
	if n == 0 {
		return []KeyCommand{
			{"left", engo.Point{-40, 0}, engo.ArrowLeft},
			{"right", engo.Point{40, 0}, engo.ArrowRight},
			{"up", engo.Point{0, -50}, engo.ArrowUp},
			{"down", engo.Point{0, 50}, engo.ArrowDown},
		}
	}
	if n == 1 {
		return []KeyCommand{
			{"2left", engo.Point{-40, 0}, engo.A},
			{"2right", engo.Point{40, 0}, engo.D},
			{"2up", engo.Point{0, -50}, engo.W},
			{"2down", engo.Point{0, 50}, engo.S},
		}
	}
	return []KeyCommand{
		{"3left", engo.Point{-40, 0}, engo.J},
		{"3right", engo.Point{40, 0}, engo.L},
		{"3up", engo.Point{0, -50}, engo.I},
		{"3down", engo.Point{0, 50}, engo.K},
	}

}

func (fms *FrogMoveSystem) Add(fg *Frog) {
	fms.frogs = append(fms.frogs, fg)
}

func (fms *FrogMoveSystem) Update(d float32) {

	for _, f := range fms.frogs {

		jc := &f.JumpComponent
		pos := &f.SpaceComponent.Position

		if f.DeathComponent.DeadTime == 0 {
			kp := false
			var rel engo.Point
			for _, v := range jc.Commands {
				if engo.Input.Button(v.KName).JustPressed() {
					rel = v.dir
					kp = true
				}
			}

			if kp {
				jc.Next = rel
			}

			if *pos == jc.Target {
				(&jc.Target).Add(jc.Next)
				jc.Next = engo.Point{0, 0}

			} else {
			}
		}

		pos.X = approach(jc.Target.X, pos.X, d*200)
		pos.Y = approach(jc.Target.Y, pos.Y, d*200)
	}
}

func (*FrogMoveSystem) Remove(e ecs.BasicEntity) {
}

type CrashSystem struct {
	obs []*Frog
}

func (cs *CrashSystem) Add(f *Frog) {
	cs.obs = append(cs.obs, f)
}

func (cs *CrashSystem) New(w *ecs.World) {
	fmt.Println("New CrashSystem")
	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		cm, ok := message.(common.CollisionMessage)
		if ok {
			a := cm.Entity.ID()
			b := cm.To.ID()

			for _, f := range cs.obs {
				if f.ID() == b {
					return
				}
			}

			for _, f := range cs.obs {
				if f.ID() == a {
					if f.DeadTime == 0 {
						f.DeadTime = 0.0001
					}
				}
			}
		}
	})
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

	doReset := false
	for _, v := range cs.obs {
		if v.DeathComponent.DeadTime > 0 {
			v.DeathComponent.DeadTime += d
			v.RenderComponent.Color = color.RGBA{255, 0, 0, 255}
		}
		if v.DeadTime > 2 {
			doReset = true
		}
	}

	if doReset {
		engo.Mailbox.Dispatch(ResetMessage{Score: false})

		for i, v := range cs.obs {
			if v.DeadTime == 0 {
				fmt.Println("Score senting")
				engo.Mailbox.Dispatch(ScoreMessage{PNum: i, Inc: 1})
			}
			v.Reset()
		}
	}
}
