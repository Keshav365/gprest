package models

import (
	"fmt"
)

type State struct {
	Missionaries int
	Cannibals    int
	Boat         int
}

type Nodec struct {
	State    State
	Parent   *Nodec
	MoveDesc string
}

var possibleMoves = []State{
	{1, 0, 1}, {2, 0, 1}, {0, 1, 1}, {0, 2, 1}, {1, 1, 1},
}

func isValidState(s State) bool {
	if s.Missionaries < 0 || s.Cannibals < 0 || s.Missionaries > 3 || s.Cannibals > 3 {
		return false
	}
	if (s.Missionaries > 0 && s.Missionaries < s.Cannibals) ||
		(3-s.Missionaries > 0 && 3-s.Missionaries < 3-s.Cannibals) {
		return false
	}
	return true
}

func BFSCM() {
	start := State{3, 3, 1}
	goal := State{0, 0, 0}
	queue := []*Nodec{{State: start, Parent: nil, MoveDesc: "Start"}}
	visited := make(map[State]bool)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.State == goal {
			printSolution(current)
			return
		}

		for _, move := range possibleMoves {
			newState := State{
				Missionaries: current.State.Missionaries - move.Missionaries*current.State.Boat,
				Cannibals:    current.State.Cannibals - move.Cannibals*current.State.Boat,
				Boat:         1 - current.State.Boat,
			}
			if isValidState(newState) && !visited[newState] {
				visited[newState] = true
				queue = append(queue, &Nodec{State: newState, Parent: current, MoveDesc: fmt.Sprintf("Move %dM %dC", move.Missionaries, move.Cannibals)})
			}
		}
	}
}

func printSolution(node *Nodec) {
	path := []string{}
	for node != nil {
		path = append([]string{node.MoveDesc}, path...)
		node = node.Parent
	}
	for _, step := range path {
		fmt.Println(step)
	}
}


