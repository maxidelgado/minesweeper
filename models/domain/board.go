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
