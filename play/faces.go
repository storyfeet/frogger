package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

//engo-faces
type ECSBasicable interface {
	GetBasicEntity() *ecs.BasicEntity
}

type Spaceable interface {
	GetSpaceComponent() *common.SpaceComponent
}
type Renderable interface {
	GetRenderComponent() *common.RenderComponent
}
type Collidable interface {
	GetCollisionComponent() *common.CollisionComponent
}

//froggerfaces
type Velocitable interface {
	GetVelocityComponent() *VelocityComponent
}

//Complete
type SpaceFace interface {
	ECSBasicable
	Spaceable
}

type DriveFace interface {
	ECSBasicable
	Spaceable
	Collidable
	Renderable
	GetVelocityComponent() *VelocityComponent
}

type SpaceEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

func RemoveSpaceEntity(sl []SpaceEntity, id uint64) []SpaceEntity {
	dp := -1
	for i, v := range sl {
		if v.ID() == id {
			dp = i
			break
		}
	}
	if dp >= 0 {
		return append(sl[:dp], sl[dp+1:]...)
	}
	return sl
}
