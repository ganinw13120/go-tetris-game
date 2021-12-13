// Package core provides ...
package core

import (
	"image/color"
)

type Space struct {
	Avaiable [][]bool
	Block    [][]color.Color
}

func NewSpace(blockCol int, blockRow int) Space {
	blocks := make([][]color.Color, 0)
	i := 0
	for i < blockRow {
		j := 0
		row := make([]color.Color, 0)
		for j < blockCol {
			row = append(row, defaultColor)
			j++
		}
		blocks = append(blocks, row)
		i++
	}
	avaiable := make([][]bool, blockRow)
	for k := range avaiable {
		avaiable[k] = make([]bool, blockCol)
		for j := range avaiable[k] {
			avaiable[k][j] = false
		}
	}
	return Space{
		Avaiable: avaiable,
		Block:    blocks,
	}
}

func (space *Space) clearAvaiable() {
	for i, v := range space.Avaiable {
		for j := range v {
			space.Avaiable[i][j] = false
		}
	}
}

func (space *Space) clearBlock() {
	for i, v := range space.Block {
		for j := range v {
			space.Block[i][j] = defaultColor
		}
	}
}

func (space *Space) checkAvaiable(row int, col int) bool {
	if row >= len(space.Avaiable) {
		return true
	} else if col >= len(space.Avaiable[0]) {
		return true
	}
	return space.Avaiable[row][col]
}

func (space *Space) updateSpace(pieces []*Piece) {
	space.clearAvaiable()
	space.clearBlock()
	for _, piece := range pieces {
		for row, v := range piece.Block {
			for col, block := range v {
				if block {
					space.Block[piece.PosY+row][piece.PosX+col] = piece.Color
					if !piece.IsMoving {
						space.Avaiable[piece.PosY+row][piece.PosX+col] = true
					}
				}
			}
		}
	}
}
