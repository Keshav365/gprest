package models

import (
	"fmt"
	"math"
)

type Node struct {
	vertex   string
	distance float64
}

func (g *Graph) Dijkstra(start string) {
	dist := make(map[string]float64)
	prev := make(map[string]string)
	visited := make(map[string]bool)

	for v := range g.gAdj {
		dist[v] = math.Inf(1)
		prev[v] = ""
	}
	dist[start] = 0

	for len(visited) < len(g.gAdj) {
		minVertex := ""
		minDist := math.Inf(1)
		for v, d := range dist {
			if !visited[v] && d < minDist {
				minDist = d
				minVertex = v
			}
		}
		if minVertex == "" {
			break
		}
		visited[minVertex] = true

		for _, neighbor := range g.gAdj[minVertex] {
			for v, w := range neighbor {
				if !visited[v] && dist[minVertex]+w < dist[v] {
					dist[v] = dist[minVertex] + w
					prev[v] = minVertex
				}
			}
		}
	}

	fmt.Println("Vertex\tDistance\tPath")
	for v, d := range dist {
		path := reconstructPath(prev, v)
		fmt.Printf("%s\t%.2f\t%s\n", v, d, path)
	}
}

func reconstructPath(prev map[string]string, target string) string {
	path := ""
	for target != "" {
		if path == "" {
			path = target
		} else {
			path = target + " -> " + path
		}
		target = prev[target]
	}
	return path
}
