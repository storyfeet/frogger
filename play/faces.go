package play

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type ECSBasicable interface {
	GetBasicEntity() *ecs.BasicEntity
}

type Spaceable interface {
	ECSBasicable
	GetSpaceComponent() *common.SpaceComponent
}
type Renderable interface {
	Spaceable
	GetRenderComponent() *common.RenderComponent
}
type Collidable interface {
	Spaceable
	GetCollisionComponent() *common.CollisionComponent
}

type Drivable interface {
	Collidable
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
