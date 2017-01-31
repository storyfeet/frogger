package play

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type ClimberSystem struct {
	Base, Step float32
	pres       float32
	obs        []SpaceEntity
}

func NewClimberSystem(base, step float32) *ClimberSystem {
	return &ClimberSystem{
		Base: base,
		Step: step,
		pres: base,
	}
}

func (cs *ClimberSystem) AddByInterface(sf SpaceFace) {
	cs.obs = append(cs.obs, SpaceEntity{sf.GetBasicEntity(), sf.GetSpaceComponent()})
}

func (cs *ClimberSystem) Remove(e ecs.BasicEntity) {
	cs.obs = RemoveSpaceEntity(cs.obs, e.ID())
}

func (cs *ClimberSystem) Update(d float32) {
	bar := cs.Base - (cs.Step * 2)
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
	if cs.Base < cs.pres {
		mDist := (70 + cs.pres - cs.Base) * d
		cs.pres -= mDist
		engo.Mailbox.Dispatch(common.CameraMessage{
			Axis:        common.YAxis,
			Value:       -mDist,
			Incremental: true,
		})
	}

}
