package play

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type SysList struct {
	Render     *common.RenderSystem
	FrogMove   *FrogMoveSystem
	CarSpawn   *CarSpawnSystem
	ObMove     *ObMoveSystem
	CollSys    *common.CollisionSystem
	CrashSys   *CrashSystem
	BoundsSys  *BoundsDeathSystem
	ClimberSys *ClimberSystem
	ScoreSys   *ScoreSystem
}

type MainScene struct{ NPlayers int }

func (*MainScene) Type() string { return "MainScene" }
func (*MainScene) Preload() {
	err := engo.Files.Load("Targa.ttf")
	if err != nil {
		fmt.Println("Could not load font Targa")
	}
}

func (ms *MainScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)

	var sList SysList

	sList.Render = &common.RenderSystem{}
	sList.FrogMove = &FrogMoveSystem{}
	sList.ObMove = &ObMoveSystem{}
	sList.CarSpawn = NewCarSpawnSystem(1, &sList, BasicRowFactory(350))
	sList.CollSys = &common.CollisionSystem{}
	sList.CrashSys = &CrashSystem{}
	sList.BoundsSys = &BoundsDeathSystem{rect: engo.AABB{engo.Point{-5, -200}, engo.Point{610, 410}}, w: w}
	sList.ClimberSys = NewClimberSystem(400, 50)
	sList.ScoreSys = &ScoreSystem{}

	for i := 0; i < ms.NPlayers; i++ {
		fc := FrogCommands(i)
		for _, kc := range fc {
			engo.Input.RegisterButton(kc.KName, kc.key)

		}
		fg1 := NewFrog(engo.Point{200, 350}, fc)
		sList.FrogMove.Add(fg1)
		sList.Render.AddByInterface(fg1)
		sList.CollSys.AddByInterface(fg1)
		sList.CrashSys.Add(fg1)
		sList.ClimberSys.AddByInterface(fg1)
		sc1 := sList.ScoreSys.CreatePlayer()
		sList.Render.AddByInterface(sc1)

	}

	w.AddSystem(sList.Render)
	w.AddSystem(sList.CollSys)
	w.AddSystem(sList.FrogMove)
	w.AddSystem(sList.ObMove)
	w.AddSystem(sList.CarSpawn)
	w.AddSystem(sList.CrashSys)
	w.AddSystem(sList.BoundsSys)
	w.AddSystem(sList.ClimberSys)
	w.AddSystem(sList.ScoreSys)

	sList.CarSpawn.Fill()

}
