package main

type Grid struct {
	Layout      [HEIGHT][WIDTH]int
	ActiveBlock *Block
	GameOver    bool
	Score       int
}

func CreateGrid() *Grid {
	return &Grid{GameOver: false}
}

func (G *Grid) Place(block *Block) {
	G.ActiveBlock = block
}

func (G *Grid) CheckGameOver() {
	for col := range WIDTH {
		if G.Layout[0][col] != 0 {
			G.GameOver = true
			return
		}
	}
}

func (G *Grid) Collapse() {
	collapseRows := []int{}
	for row := range HEIGHT {
		collapse := true
		for col := range WIDTH {
			if G.Layout[row][col] == 0 {
				collapse = false
				break
			}
		}
		if collapse {
			for col := range WIDTH {
				G.Layout[row][col] = YELLOW
				Draw(G)
			}
			collapseRows = append(collapseRows, row)
		}
	}

	for _, row := range collapseRows {
		for base := row; base > 0; base-- {
			for col := range WIDTH {
				G.Layout[base][col] = G.Layout[base-1][col]
			}
		}
	}

	G.Score += len(collapseRows)
}

func (G *Grid) MoveDown() {
	if G.ActiveBlock == nil {
		return
	}

	block := *G.ActiveBlock
	block.Move(0, 1)

	collision := false
	for row := range 3 {
		for col := range 3 {
			y := block.Y + row
			x := block.X + col
			if block.Layout[row][col] > 0 &&
				(y < 0 || y >= HEIGHT ||
					x < 0 || x >= WIDTH ||
					G.Layout[y][x] > 0) {
				collision = true
				goto exit
			}
		}
	}

exit:
	if !collision {
		G.ActiveBlock.Move(0, 1)
	} else {
		for row := range 3 {
			for col := range 3 {
				y := G.ActiveBlock.Y + row
				x := G.ActiveBlock.X + col
				if y >= 0 && y < HEIGHT && x >= 0 && x < WIDTH {
					G.Layout[y][x] += G.ActiveBlock.Layout[row][col]
				}
			}
		}
		G.Collapse()
		G.CheckGameOver()
		G.ActiveBlock = nil
	}
}

func (G *Grid) Move(x int) {
	if G.ActiveBlock == nil {
		return
	}

	block := *G.ActiveBlock
	block.Move(x, 0)

	for row := range 3 {
		for col := range 3 {
			y := block.Y + row
			x := block.X + col
			if block.Layout[row][col] > 0 &&
				(y < 0 || y >= HEIGHT || x < 0 || x >= WIDTH ||
					G.Layout[y][x] > 0) {
				return
			}
		}
	}
	G.ActiveBlock.Move(x, 0)
}

func (G *Grid) Rotate() {
	if G.ActiveBlock == nil {
		return
	}

	block := *G.ActiveBlock
	block.Rotate()

	for row := range 3 {
		for col := range 3 {
			y := block.Y + row
			x := block.X + col
			if block.Layout[row][col] > 0 &&
				(y < 0 || y >= HEIGHT || x < 0 || x >= WIDTH ||
					G.Layout[y][x] > 0) {
				return
			}
		}
	}

	G.ActiveBlock.Rotate()
}
