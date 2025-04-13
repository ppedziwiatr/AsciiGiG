package main

import (
	"github.com/ppedziwiatr/ascii-gig/pkg/game"
	"github.com/ppedziwiatr/ascii-gig/pkg/ui"
)

func main() {
	ui.DisplayTitleScreen()

	g := game.NewGame()
	g.Run()
}
