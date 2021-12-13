package main

import (
	"log"
	"tetris/core"

	"github.com/hajimehoshi/ebiten/v2"
)

const blockRow int = 20
const blockCol int = 10

const blockHeight int = 20
const blockWidth int = 20

const startX int = 3
const startY int = 0

const gap int = 1

func startGame() {
	game := core.NewGame(blockHeight, blockWidth, blockCol, blockRow, gap, startX, startY)
	ebiten.SetWindowSize((gap*blockCol)+(blockCol*blockWidth), (gap*blockRow)+(blockRow*blockHeight))
	ebiten.SetWindowTitle("Tetris")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

func main() {
	startGame()
}
