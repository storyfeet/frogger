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
}

func (dc *DeathComponent) GetDeathComponent() *DeathComponent {
	return dc
}

type JumpComponent struct {
	Target engo.Point
	Next   engo.Point
}

func (jc *JumpComponent) GetJumpComponent() *JumpComponent {
	return jc
}
