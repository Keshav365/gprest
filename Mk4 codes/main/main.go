package main

import (
	"fmt"
	"math"
	"strconv"
)


type Graph struct {
	gAdj map[string]map[string]float64
}


func NewGraph() Graph {
	return Graph{gAdj: make(map[string]map[string]float64)}
}


func (g *Graph) AddEdge(src, dest string, w float64) {
	if _, exists := g.gAdj[src]; !exists {
		g.gAdj[src] = make(map[string]float64)
	}
	if _, exists := g.gAdj[dest]; !exists {
		g.gAdj[dest] = make(map[string]float64)
	}
	g.gAdj[src][dest] = w
	g.gAdj[dest][src] = w 
}


func CnvtMat2Edge(objMat [][]float64) Graph {
	g := NewGraph()
	for i := 0; i < len(objMat); i++ {
		for j := 0; j < len(objMat[i]); j++ {
			if objMat[i][j] != -1 && objMat[i][j] != 0 {
				g.AddEdge(strconv.Itoa(i), strconv.Itoa(j), objMat[i][j])
			}
		}
	}
	return g
}


func (g *Graph) TSP(start string) {
	vertices := []string{}
	for v := range g.gAdj {
		vertices = append(vertices, v)
	}
	n := len(vertices)

	if n < 2 {
		fmt.Println("Graph must have at least 2 vertices for TSP.")
		return
	}

	
	minPath := math.Inf(1)
	var bestPath []string

	
	vertexIndex := make(map[string]int)
	indexVertex := make(map[int]string)
	for i, v := range vertices {
		vertexIndex[v] = i
		indexVertex[i] = v
	}

	
	costMatrix := make([][]float64, n)
	for i := range costMatrix {
		costMatrix[i] = make([]float64, n)
		for j := range costMatrix[i] {
			if i == j {
				costMatrix[i][j] = math.Inf(1)
			} else if val, exists := g.gAdj[indexVertex[i]][indexVertex[j]]; exists {
				costMatrix[i][j] = val
			} else {
				costMatrix[i][j] = math.Inf(1)
			}
		}
	}

	
	var tspHelper func(currIndex int, visited []bool, path []int, cost float64)

	tspHelper = func(currIndex int, visited []bool, path []int, cost float64) {
		if len(path) == n {
			if costMatrix[currIndex][vertexIndex[start]] != math.Inf(1) {
				totalCost := cost + costMatrix[currIndex][vertexIndex[start]]
				if totalCost < minPath {
					minPath = totalCost
					bestPath = make([]string, len(path))
					for i, v := range path {
						bestPath[i] = indexVertex[v]
					}
					bestPath = append(bestPath, start)
				}
			}
			return
		}

		for i := 0; i < n; i++ {
			if !visited[i] && costMatrix[currIndex][i] != math.Inf(1) {
				visited[i] = true
				tspHelper(i, visited, append(path, i), cost+costMatrix[currIndex][i])
				visited[i] = false
			}
		}
	}

	
	startIndex := vertexIndex[start]
	visited := make([]bool, n)
	visited[startIndex] = true
	tspHelper(startIndex, visited, []int{startIndex}, 0)

	fmt.Println("Best TSP Path:", bestPath, "with distance:", minPath)
}


func main() {
	objMat := [][]float64{
		{0, 4, -1, -1, -1, -1, -1, 8, -1},
		{4, 0, 8, -1, -1, -1, -1, 11, -1},
		{-1, 8, 0, 7, -1, 4, -1, -1, 2},
		{-1, -1, 7, 0, 9, 14, -1, -1, -1},
		{-1, -1, -1, 9, 0, 10, -1, -1, -1},
		{-1, -1, 4, 14, 10, 0, 2, -1, -1},
		{-1, -1, -1, -1, -1, 2, 0, 1, 6},
		{8, 11, -1, -1, -1, -1, 1, 0, -1},
		{-1, -1, 2, -1, -1, -1, 6, 7, 0},
	}

	g := CnvtMat2Edge(objMat)
	g.TSP("0") 
}
