package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/term"
)

const WIDTH, HEIGHT = 10, 21
const (
	BLACK = iota
	WHITE
	RED
	GREEN
	BLUE
	PURPLE
	ORANGE
	BROWN
	YELLOW
)

var (
	inputQueue = make(chan byte, 1)
)

func RawMode() func() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	return func() {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}
}

func GetInput() {
	inputBuffer := make([]byte, 1)
	for {
		os.Stdin.Read(inputBuffer)
		select {
		case inputQueue <- inputBuffer[0]:
		default:
		}
	}
}

func HandleInput(G *Grid, reset func()) {
	select {
	case input := <-inputQueue:
		switch input {
		case 'a', 'A':
			G.Move(-1)
		case 'd', 'D':
			G.Move(1)
		case 's', 'S':
			for range 20 {
				G.MoveDown()
			}
		case 'r', 'R':
			G.Rotate()
		case 'q', 'Q':
			reset()
			Draw(G)
			println("You Quit!\nYour Score:", G.Score, "\nPress Ctrl + C to Exit")
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, syscall.SIGINT)
			<-sig
			os.Exit(0)
		default:
		}
	default:
	}
}

func main() {
	reset := RawMode()

	go GetInput()

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	grid := CreateGrid()
	for !grid.GameOver {
		if grid.ActiveBlock == nil {
			block := CreateBlock([]string{"T", "L", "O", "I", "J", "S"}[rand.Intn(4)], 3, 0, rand.Intn(BROWN)+1)
			grid.Place(block)
		}
		HandleInput(grid, reset)
		Draw(grid)
		grid.MoveDown()
		println("Score: ", grid.Score, "\nPress A, S, D for Movement, R for Rotation, Q to Quit")
		<-ticker.C
	}

	reset()
	Draw(grid)
	println("Game Over!\nYour Score:", grid.Score, "\nPress Ctrl + C to Exit")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
}
