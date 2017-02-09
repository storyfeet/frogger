package main

import (
	"flag"

	"engo.io/engo"
	"github.com/coderconvoy/frogger/play"
)

func main() {
	np := flag.Int("np", 1, "np: Number of players")
	flag.Parse()
	opts := engo.RunOptions{
		Title:         "Frogger",
		ScaleOnResize: true,
		Width:         600,
		Height:        400,
	}
	engo.Run(opts, &play.MainScene{*np})
}
