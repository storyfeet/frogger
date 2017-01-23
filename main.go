package main

import (
	"engo.io/engo"
	"github.com/coderconvoy/frogger/play"
)

func main() {
	opts := engo.RunOptions{
		Title:  "Frogger",
		Width:  600,
		Height: 400,
	}
	engo.Run(opts, &play.MainScene{})
}
