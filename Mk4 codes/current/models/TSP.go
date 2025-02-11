package models

import (
	"fmt"
	"math"
)

func (g *Graph) TSP(start string) {
	vertices := g.GetVertices()
	// n := len(vertices)
	minPath := math.Inf(1)
	var bestPath []string

	permute(vertices, 0, func(path []string) {
		if path[0] != start {
			return
		}
		totalDist := 0.0
		valid := true

		for i := 0; i < len(path)-1; i++ {
			edgeFound := false
			for _, neighbor := range g.gAdj[path[i]] {
				if w, exists := neighbor[path[i+1]]; exists {
					totalDist += w
					edgeFound = true
					break
				}
			}
			if !edgeFound {
				valid = false
				break
			}
		}

		if valid && totalDist < minPath {
			minPath = totalDist
			bestPath = append([]string{}, path...)
		}
	})

	fmt.Println("Best TSP Path:", bestPath, "with distance:", minPath)
}

func permute(arr []string, l int, callback func([]string)) {
	if l == len(arr)-1 {
		callback(arr)
		return
	}
	for i := l; i < len(arr); i++ {
		arr[l], arr[i] = arr[i], arr[l]
		permute(arr, l+1, callback)
		arr[l], arr[i] = arr[i], arr[l]
	}
}
