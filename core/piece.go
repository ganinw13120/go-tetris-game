package core

import (
	"image/color"
	"math/rand"
	"tetris/util"
	"time"
)

type Piece struct {
	Type      int
	Block     [][]bool
	Color     color.Color
	PosX      int
	PosY      int
	boundMinX int
	boundMaxX int
	boundMinY int
	boundMaxY int
	IsMoving  bool
}

func NewPiece(startX int, startY int, boundMinX int, boundMaxX int, boundMinY int, boundMaxY int) Piece {
	Types := make([][][]bool, 0)
	Types = append(Types, [][]bool{
		{true, true, true},
		{false, false, true},
		{false, false, false},
	})
	Types = append(Types, [][]bool{
		{true, true, true, true},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	})
	Types = append(Types, [][]bool{
		{true, true},
		{true, true},
	})
	Types = append(Types, [][]bool{
		{true, true, false},
		{false, true, true},
		{false, false, false},
	})
	Types = append(Types, [][]bool{
		{false, true, true},
		{true, true, false},
		{false, false, false},
	})
	blockColor := make([]color.Color, 0)
	blockColor = append(blockColor, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
	blockColor = append(blockColor, color.RGBA{0xFF, 0xAF, 0xAF, 0xFF})
	blockColor = append(blockColor, color.RGBA{0xFF, 0x80, 0x00, 0xFF})
	blockColor = append(blockColor, color.RGBA{0x00, 0x00, 0xFF, 0xFF})
	blockColor = append(blockColor, color.RGBA{0xFF, 0x00, 0xFF, 0xFF})
	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(len(Types))
	return Piece{
		Type:      1,
		Block:     Types[ran],
		Color:     blockColor[ran],
		PosX:      startX,
		PosY:      startY,
		boundMinX: boundMinX,
		boundMaxX: boundMaxX,
		boundMinY: boundMinY,
		boundMaxY: boundMaxY,
		IsMoving:  true,
	}
}

func (piece *Piece) MoveVertical(distance int, space Space) {
	isFloorHit := false
	for i, v := range piece.Block {
		targetY := util.Lerp(piece.PosY+distance, piece.boundMinY, piece.boundMaxY)
		for j, block := range v {
			if block && space.checkAvaiable(i+targetY, j+piece.PosX) {
				isFloorHit = true
			}
		}
	}
	if !isFloorHit {
		piece.PosY = util.Lerp(piece.PosY+distance, piece.boundMinY, piece.boundMaxY)
		if piece.PosY == piece.boundMaxY {
			piece.IsMoving = false
		}
	} else {
		piece.IsMoving = false
	}
}

func (piece *Piece) MoveHorizontal(distance int, space Space) {
	isFloorHit := false
	targetX := util.Lerp(piece.PosX+distance, piece.boundMinX, piece.boundMaxX)
	for i, v := range piece.Block {
		for j, block := range v {
			if block && space.checkAvaiable(i+piece.PosY, j+targetX) {
				isFloorHit = true
			}
		}
	}
	if !isFloorHit {
		piece.PosX = util.Lerp(piece.PosX+distance, piece.boundMinX, piece.boundMaxX)
	}
}

func (piece *Piece) Rotate(space Space) {
	tempBlock := piece.duplicateBlock()
	for x := 0; x < len(tempBlock)/2; x++ {
		for y := x; y < len(tempBlock)-x-1; y++ {
			temp := tempBlock[x][y]
			tempBlock[x][y] = tempBlock[y][len(tempBlock)-1-x]
			tempBlock[y][len(tempBlock)-1-x] = tempBlock[len(tempBlock)-1-x][len(tempBlock)-1-y]
			tempBlock[len(tempBlock)-1-x][len(tempBlock)-1-y] = tempBlock[len(tempBlock)-1-y][x]
			tempBlock[len(tempBlock)-1-y][x] = temp
		}
	}
	isValid := true
	for y, v := range tempBlock {
		for x, block := range v {
			if block && space.checkAvaiable(y+piece.PosY, x+piece.PosX) {
				isValid = false
			}
		}
	}
	if isValid {
		piece.Block = tempBlock
	}
}

func (piece *Piece) duplicateBlock() [][]bool {
	val := make([][]bool, 0)
	for _, v := range piece.Block {
		row := make([]bool, 0)
		for _, block := range v {
			row = append(row, block)
		}
		val = append(val, row)
	}
	return val
}

func (piece *Piece) ForceMoveVertical(distance int) {
	piece.PosY = util.Lerp(piece.PosY+distance, piece.boundMinY, piece.boundMaxY)
}
