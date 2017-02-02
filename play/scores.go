package play

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type ScoreEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	Score int
}

type ScoreSystem struct {
	scores []*ScoreEntity
}

func (ss *ScoreSystem) CreatePlayer() *ScoreEntity {
	fnt := &common.Font{
		URL:  "Targa.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		fmt.Println("Could not Create preloaded Targa.ttf")
	}

	np := &ScoreEntity{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable: common.Text{
				Font: fnt,
				Text: "score 0",
			},
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{0, 0},
		},
	}

	np.SetShader(common.HUDShader)

	ss.scores = append(ss.scores, np)
	return np
}

func (ss *ScoreSystem) New(w ecs.World) {
	engo.Mailbox.Listen("ScoreMessage", func(m engo.Message) {
	})
}

func (ss *ScoreSystem) Remove(e ecs.BasicEntity) {}

func (ss *ScoreSystem) Update(d float32) {}
