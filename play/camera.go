package play

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type ClimberSystem struct {
	Base, Step, Init float32
	Pres             float32
	obs              []SpaceEntity
}

func NewClimberSystem(base, step float32) *ClimberSystem {
	return &ClimberSystem{
		Init: base,
		Base: base,
		Step: step,
		Pres: base,
	}
}

func (cs *ClimberSystem) AddByInterface(sf SpaceFace) {
	cs.obs = append(cs.obs, SpaceEntity{sf.GetBasicEntity(), sf.GetSpaceComponent()})
}

func (cs *ClimberSystem) Remove(e ecs.BasicEntity) {
	cs.obs = RemoveSpaceEntity(cs.obs, e.ID())
}

func (cs *ClimberSystem) New(w *ecs.World) {
	engo.Mailbox.Listen("ResetMessage", func(message engo.Message) {
		fmt.Println("ResetReceived")
		engo.Mailbox.Dispatch(common.CameraMessage{
			Axis:        common.YAxis,
			Value:       cs.Init - cs.Base,
			Incremental: true,
		})
		cs.Base = cs.Init
		cs.Pres = cs.Init

	})
}

func (cs *ClimberSystem) Update(d float32) {
	bar := cs.Base - (cs.Step * 3)
	highest := bar
	for _, f := range cs.obs {
		if f.Position.Y > highest {
			highest = f.Position.Y
		}
	}

	if highest == bar {
		cs.Base = cs.Base - cs.Step
		engo.Mailbox.Dispatch(RowMessage{1})
		common.CameraBounds.Min.Y -= 50

	}
	if cs.Base < cs.Pres {
		mDist := (70 + cs.Pres - cs.Base) * d
		cs.Pres -= mDist
		engo.Mailbox.Dispatch(common.CameraMessage{
			Axis:        common.YAxis,
			Value:       -mDist,
			Incremental: true,
		})
	}

}
