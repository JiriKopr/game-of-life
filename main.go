package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gol/constants"
	. "gol/node"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Println("Error while creating new screen")
	}

	if err := screen.Init(); err != nil {
		fmt.Println("Error during initialization")
	}

	screen.Clear()
	screen.Show()

	width, height := screen.Size()

	start := NewNode()
	var top *Node

	for x := 0; x < width; x++ {
		var firstInRow *Node
		var previous *Node
		for y := 0; y < height; y++ {
			var current *Node
			if x == 0 && y == 0 {
				current = start
			} else {
				current = NewNode()
			}

			if previous != nil {
				previous.Right = current
			}

			if top != nil {
				top.Bottom = current
			}

			current.Left = previous
			current.Top = top
			current.X = x
			current.Y = y
			current.Screen = &screen

			if rand.Intn(100) < constants.INITIAL_SPAWN_CHANGE {
				current.TurnOn()
			} else {
				current.TurnOff()
			}

			if firstInRow == nil {
				firstInRow = current
			}

			if top != nil {
				top = top.Right
			}

			previous = current
		}

		top = firstInRow
	}

	for {
		for row := start; row != nil; row = row.Bottom {
			for column := row; column != nil; column = column.Right {
				column.CalculateState()
			}
		}

		for row := start; row != nil; row = row.Bottom {
			for column := row; column != nil; column = column.Right {
				column.UpdateState()
			}
		}

		screen.Show()

		time.Sleep(70 * time.Millisecond)

		if !screen.HasPendingEvent() {
			continue
		}

		ev := screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}

}
