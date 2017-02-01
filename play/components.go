package play

import (
	"engo.io/engo"
)

type VelocityComponent struct {
	Vel engo.Point
}

func (vc *VelocityComponent) GetVelocityComponent() *VelocityComponent {
	return vc
}

type DeathComponent struct {
	DeadTime float32
	Origin   engo.Point
}

func (dc *DeathComponent) GetDeathComponent() *DeathComponent {
	return dc
}

type KeyCommand struct {
	k   string
	dir engo.Point
}

type JumpComponent struct {
	Target   engo.Point
	Next     engo.Point
	Commands []KeyCommand
}

func (jc *JumpComponent) GetJumpComponent() *JumpComponent {
	return jc
}
