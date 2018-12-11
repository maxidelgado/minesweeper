package domain

import (
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Cell struct {
	HasMine  bool `json:"hasMine"`
	HasFlag  bool `json:"hasFlag"`
	Revealed bool `json:"revealed"`
	XAxis    int  `json:"xAxis"`
	YAxis    int  `json:"yAxis"`
	Value    int  `json:"value"`
}

type Rows []Cell

type Board struct {
	Rows  int    `json:"rows"`
	Cols  int    `json:"cols"`
	Cells []Rows `json:"cells"`
}


func (b *Board) Setup() {
	rows := make(Rows, b.Rows*b.Cols)

	b.setMines(rows)
	b.setValues()
}

func (b *Board) OpenCell(cell Cell) {
	b.Cells[cell.XAxis][cell.YAxis].Revealed = true

	if cell.Value == 0 {
		adjacentCells := b.getAdjacentCells(cell)
		for _, c := range adjacentCells {
			if c.Value == 0 && !c.Revealed {
				b.OpenCell(c)
			}
		}
	}
}

func (b *Board) FlagCell(cell Cell) {
	b.Cells[cell.XAxis][cell.YAxis].HasFlag = true
}

func (b Board) getAdjacentCells(cell Cell) (cells []Cell) {
	for i := cell.XAxis - 1; i < cell.XAxis+2; i++ {
		if i < 0 || i > b.Rows-1 {
			continue
		}
		for j := cell.YAxis - 1; j < cell.YAxis+2; j++ {
			if j < 0 || j > b.Cols-1 {
				continue
			}
			if i == cell.XAxis && j == cell.YAxis {
				continue
			}
			cell := b.Cells[i][j]
			cells = append(cells, cell)
		}
	}
	return
}

func (b *Board) setMines(rows Rows) {
	i := 0
	maxMines := b.Rows * 2

	for i < maxMines {
		x := rand.Intn(b.Rows * b.Cols)

		if !rows[x].HasMine {
			rows[x].HasMine = true
			rows[x].Value = -1
			i++
		}
	}

	b.Cells = make([]Rows, b.Rows)

	for row := range b.Cells {
		b.Cells[row] = rows[(b.Cols * row):(b.Cols * (row + 1))]
	}

	for i, row := range b.Cells {
		for j, _ := range row {
			b.Cells[i][j].XAxis = i
			b.Cells[i][j].YAxis = j
		}
	}
}

func (b *Board) setValues() {
	for i := 0; i < b.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			adjacentCells := b.getAdjacentCells(b.Cells[i][j])

			for _, c := range adjacentCells {
				if c.HasMine {
					b.Cells[i][j].Value++
				}
			}
		}
	}
}