package node

import (
	"gol/constants"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Node struct {
	Left   *Node
	Top    *Node
	Right  *Node
	Bottom *Node

	WillBeOn bool
	IsOn     bool

	X int
	Y int

	Screen *tcell.Screen
}

func NewNode() *Node {
	return &Node{
		IsOn:     false,
		WillBeOn: false,
	}
}

func (node *Node) TurnOn() {
	(*node.Screen).SetContent(
		node.X,
		node.Y,
		' ',
		[]rune{},
		tcell.StyleDefault.Background(tcell.ColorWhite),
	)

	node.IsOn = true
}

func (node *Node) TurnOff() {
	(*node.Screen).SetContent(
		node.X,
		node.Y,
		' ',
		[]rune{},
		tcell.StyleDefault.Background(tcell.ColorBlack),
	)

	node.IsOn = false
}

func (node *Node) GetNeighbors() []*Node {
	neighbors := []*Node{}

	if node.Left != nil {
		left := node.Left

		neighbors = append(neighbors, left.Top)
		neighbors = append(neighbors, left)
		neighbors = append(neighbors, left.Bottom)
	}

	neighbors = append(neighbors, node.Top)
	neighbors = append(neighbors, node.Bottom)

	if node.Right != nil {
		right := node.Right

		neighbors = append(neighbors, right.Top)
		neighbors = append(neighbors, right)
		neighbors = append(neighbors, right.Bottom)
	}

	filteredNeighbors := []*Node{}

	for _, neighbor := range neighbors {
		if neighbor != nil {
			filteredNeighbors = append(filteredNeighbors, neighbor)
		}
	}

	return filteredNeighbors
}

func (node *Node) GetAliveNeighborCount() int {

	neighbors := node.GetNeighbors()

	aliveCount := 0

	for _, neighbor := range neighbors {
		if neighbor.IsOn {
			aliveCount++
		}
	}

	return aliveCount
}

func (node *Node) WillStayAlive() bool {
	aliveCount := node.GetAliveNeighborCount()

	return aliveCount == 2 || aliveCount == 3
}

func (node *Node) WillComeToLife() bool {
	aliveCount := node.GetAliveNeighborCount()

	return aliveCount == 3
}

func (node *Node) CalculateState() {
	if node.IsOn && node.WillStayAlive() {
		node.WillBeOn = true
		return
	}

	if node.WillComeToLife() {
		node.WillBeOn = true
		return
	}

	if constants.RANDOM_NOISE && rand.Intn(1000) < 1 {
		node.WillBeOn = true
		return
	}

	node.WillBeOn = false
}

func (node *Node) UpdateState() {
	if node.WillBeOn {
		node.TurnOn()
	} else {
		node.TurnOff()
	}
}
