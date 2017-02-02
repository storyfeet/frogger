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
	fnt    *common.Font
}

func (ss *ScoreSystem) CreatePlayer() *ScoreEntity {
	if ss.fnt == nil {
		ss.fnt = &common.Font{
			URL:  "Targa.ttf",
			FG:   color.Black,
			Size: 64,
		}
	}
	err := ss.fnt.CreatePreloaded()
	if err != nil {
		fmt.Println("Could not Create preloaded Targa.ttf")
	}
	pnum := len(ss.scores)

	np := &ScoreEntity{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable: common.Text{
				Font: ss.fnt,
				Text: fmt.Sprintf("P%d 0", pnum),
			},
		},
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{float32(pnum * 200), 0},
		},
	}

	np.SetZIndex(10)
	np.SetShader(common.HUDShader)

	ss.scores = append(ss.scores, np)
	return np
}

func (ss *ScoreSystem) New(w *ecs.World) {
	fmt.Println("New ScoreSystem")
	engo.Mailbox.Listen("ScoreMessage", func(m engo.Message) {
		fmt.Println("ScoreMessage Recieved :")
		sm, ok := m.(ScoreMessage)
		if !ok {
			return
		}
		for i, v := range ss.scores {
			if i == sm.PNum {
				v.Score += sm.Inc
				v.RenderComponent.Drawable = common.Text{
					Font: ss.fnt,
					Text: fmt.Sprintf("P%d %d", i, v.Score),
				}
			}
		}

	})
}

func (ss *ScoreSystem) Remove(e ecs.BasicEntity) {}

func (ss *ScoreSystem) Update(d float32) {}
