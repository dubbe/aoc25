package main

import (
	"fmt"
)

type Coord struct {
	Row int
	Col int
}

type Grid struct {
	Data   []int
	Width  int
	Height int
}

func NewGrid(width, height int) *Grid {
	size := width * height
	data := make([]int, size)

	return &Grid{
		Data:   data,
		Width:  width,
		Height: height,
	}
}

func (g *Grid) calculateIndex(row, col int) int {
	return (row * g.Width) + col
}

func (g *Grid) isIndexValid(row, col int) bool {
	return row >= 0 && row < g.Height && col >= 0 && col < g.Width
}

func (g *Grid) Get(row, col int) (int, bool) {
	if !g.isIndexValid(row, col) {
		return 0, false
	}

	index := g.calculateIndex(row, col)
	return g.Data[index], true
}

func (g *Grid) Set(row, col, value int) bool {
	if !g.isIndexValid(row, col) {
		return false
	}

	index := g.calculateIndex(row, col)
	g.Data[index] = value
	return true
}

func (g *Grid) GetAdjacent(row, col int) []Coord {
	neighbors := []Coord{}

	offsets := []struct{ dr, dc int }{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, offset := range offsets {
		newRow := row + offset.dr
		newCol := col + offset.dc

		if g.isIndexValid(newRow, newCol) {
			neighborValue, _ := g.Get(newRow, newCol)

			if neighborValue > 0 {
				neighbors = append(neighbors, Coord{Row: newRow, Col: newCol})
			}
		}
	}

	return neighbors
}

func (g *Grid) Iterate() <-chan Coord {
	ch := make(chan Coord)

	go func() {
		for r := 0; r < g.Height; r++ {
			for c := 0; c < g.Width; c++ {
				index := g.calculateIndex(r, c)

				if g.Data[index] > 0 {
					ch <- Coord{Row: r, Col: c}
				}
			}
		}
		close(ch)
	}()

	return ch
}

func (g *Grid) Print() {
	for r := 0; r < g.Height; r++ {
		for c := 0; c < g.Width; c++ {
			index := g.calculateIndex(r, c)
			value := g.Data[index]

			v := "."
			if value == 1 {
				v = "@"
			} else if value == 2 {
				v = "x"
			}
			fmt.Printf("%s", v)
		}
		fmt.Println("")
	}
}
