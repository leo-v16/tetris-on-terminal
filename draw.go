package main

import "strings"

const (
	RESET_COLOR  = "\033[0m"
	BLACK_COLOR  = "\033[48;2;0;0;0m"       // 0
	WHITE_COLOR  = "\033[48;2;255;255;255m" // 1
	RED_COLOR    = "\033[48;2;255;0;0m"     // 2
	GREEN_COLOR  = "\033[48;2;0;255;0m"     // 3
	BLUE_COLOR   = "\033[48;2;0;0;255m"     // 4
	YELLOW_COLOR = "\033[48;2;255;255;0m"   // 5
)

func Draw(G *Grid) {
	var builder strings.Builder
	builder.WriteString("\033[H\033[2J" + strings.Repeat("🔳", WIDTH+2) + "\n")
	for row := range HEIGHT {
		builder.WriteString(RESET_COLOR + "🔳")
		for col := range WIDTH {
			pixelValue := G.Layout[row][col]
			if G.ActiveBlock != nil &&
				row >= G.ActiveBlock.Y && row <= G.ActiveBlock.Y+2 &&
				col >= G.ActiveBlock.X && col <= G.ActiveBlock.X+2 &&
				G.ActiveBlock.Layout[row-G.ActiveBlock.Y][col-G.ActiveBlock.X] > 0 {
				pixelValue += G.ActiveBlock.Layout[row-G.ActiveBlock.Y][col-G.ActiveBlock.X]
			}
			switch pixelValue {
			case BLACK:
				builder.WriteString(BLACK_COLOR + "  ")
			case WHITE:
				// builder.WriteString("⬜")
				builder.WriteString("🟨")
			case RED:
				builder.WriteString("🟥")
			case GREEN:
				builder.WriteString("🟩")
			case BLUE:
				builder.WriteString("🟦")
			case PURPLE:
				builder.WriteString("🟪")
			case ORANGE:
				builder.WriteString("🟧")
			case BROWN:
				builder.WriteString("🟫")
			case YELLOW:
				builder.WriteString("🟨")
			default:
			}
			// builder.WriteString("  ")
		}
		builder.WriteString(RESET_COLOR + "🔳\n")
	}
	builder.WriteString(strings.Repeat("\033[0m🔳", WIDTH+2) + RESET_COLOR)
	finalOuput := builder.String()
	println(finalOuput)
}
