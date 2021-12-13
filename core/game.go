package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const roundClock int = 10

type Game struct {
	// blocks      [][]color.Color
	space       Space
	blockHeight int
	blockWidth  int
	blockCol    int
	blockRow    int
	gap         int
	startX      int
	startY      int
	pieces      []*Piece
	fallenPiece *Piece
	clock       int
}

// https://enchantia.com/software/graphapp/doc/tutorial/colours.htm
var defaultColor color.Color = color.RGBA{0x80, 0x80, 0x80, 0xFF}

func NewGame(blockHeight int, blockWidth int, blockCol int, blockRow int, gap int, startX int, startY int) Game {
	space := NewSpace(blockCol, blockRow)
	fallenPiece := NewPiece(startX, startY, 0, blockCol-1, 0, blockRow-1)
	pieces := make([]*Piece, 0)
	pieces = append(pieces, &fallenPiece)
	game := Game{
		space:       space,
		blockHeight: blockHeight,
		blockWidth:  blockWidth,
		blockCol:    blockCol,
		blockRow:    blockRow,
		gap:         gap,
		fallenPiece: &fallenPiece,
		pieces:      pieces,
		startX:      startX,
		startY:      startY,
		clock:       0,
	}
	return game
}

func (g *Game) createNewPiece() {
	fallenPiece := NewPiece(g.startX, g.startY, 0, g.blockCol-1, 0, g.blockRow-1)
	g.pieces = append(g.pieces, &fallenPiece)
	g.fallenPiece = &fallenPiece
}

func (g *Game) onBlockCollide() {
	blocks := make([][]*bool, g.blockRow)
	for k := range blocks {
		blocks[k] = make([]*bool, g.blockCol)
	}
	for p, piece := range g.pieces {
		for i, row := range piece.Block {
			for j, col := range row {
				if col {
					blocks[piece.PosY+i][piece.PosX+j] = &(g.pieces[p].Block[i][j])
				}
			}
		}
	}
	for _, row := range blocks {
		isValidRow := true
		for _, col := range row {
			if col == nil {
				isValidRow = false
				break
			}
		}
		if isValidRow {
			for _, col := range row {
				*col = false
			}
			//Move down
		}
	}
}

func (g *Game) Update() error {
	for _, k := range inpututil.PressedKeys() {
		if k == ebiten.KeyRight && inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.fallenPiece.MoveHorizontal(1, g.space)
		}
		if k == ebiten.KeyLeft && inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.fallenPiece.MoveHorizontal(-1, g.space)
		}
		if k == ebiten.KeySpace && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.fallenPiece.Rotate(g.space)
		}
	}

	g.clock++
	if g.clock >= roundClock {
		g.clock = 0
	} else {
		return nil
	}
	g.fallenPiece.MoveVertical(1, g.space)

	if !g.fallenPiece.IsMoving {
		g.space.updateSpace(g.pieces)
		g.onBlockCollide()
		g.createNewPiece()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.space.updateSpace(g.pieces)
	for row, v := range g.space.Block {
		for col, block := range v {
			imgBlock := ebiten.NewImage(g.blockWidth, g.blockHeight)
			imgBlock.Fill(block)
			opImg := ebiten.DrawImageOptions{}
			opImg.GeoM.Translate(float64((g.gap*col)+(col*g.blockWidth)), float64((g.gap*row)+(row*g.blockWidth)))
			screen.DrawImage(imgBlock, &opImg)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
	return (g.gap * g.blockCol) + (g.blockCol * g.blockWidth), (g.gap * g.blockRow) + (g.blockRow * g.blockHeight)
}
