package main

type Block struct {
	Layout [3][3]int
	X      int
	Y      int
	Type   string
	Color  int
}

func CreateBlock(blockType string, x, y, color int) *Block {
	typeMap := map[string][3][3]int{
		"T": {
			{0, 1, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
		"L": {
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 1},
		},
		"J": {
			{0, 1, 0},
			{0, 1, 0},
			{1, 1, 0},
		},
		"S": {
			{0, 1, 1},
			{0, 1, 0},
			{1, 0, 1},
		},
		"O": {
			{0, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		},
		"I": {
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
	}

	layout := typeMap[blockType]
	for row := range 3 {
		for col := range 3 {
			layout[row][col] *= color
		}
	}

	return &Block{layout, x, y, blockType, color}
}

func (B *Block) Rotate() {
	rotation := [3][3]int{}
	for row := range 3 {
		for col := range 3 {
			rotation[row][col] = B.Layout[col][row]
		}
	}

	for row := range 3 {
		rotation[row][0], rotation[row][2] = rotation[row][2], rotation[row][0]
	}

	B.Layout = rotation
}

func (B *Block) Move(x, y int) {
	B.X += x
	B.Y += y
}
