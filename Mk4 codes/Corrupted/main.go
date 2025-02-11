package main

import (
	models "main/models"
)

func main() {

	g1 := models.NewGraph()
	g1.AddEdge("0", "1", 4)
	g1.AddEdge("0", "3", 5)
	g1.AddEdge("1", "2", 8)
	g1.AddEdge("1", "5", 9)
	g1.AddEdge("2", "4", 1)
	g1.AddEdge("2", "6", 1)
	g1.AddEdge("3", "5", 2)
	g1.AddEdge("4", "1", 1)
	g1.AddEdge("4", "5", 1)
	g1.AddEdge("5", "0", 1)
	g1.AddEdge("6", "4", 1)
	// models.BFS("0", &g1)
	// fmt.Println(models.DFS("0", &g1))

	objMat := [][]float64{
		{0, 4, -1, 5, -1, -1, -1},
		{-1, 0, 8, -1, -1, 9, -1},
		{-1, -1, 0, -1, 1, -1, 1},
		{-1, -1, -1, 0, -1, 2, -1},
		{-1, 1, -1, -1, 0, 1, -1},
		{1, -1, -1, -1, -1, 0, -1},
		{-1, -1, -1, -1, -1, 1, 0},
	}

	models.PrintGraph(objMat)
	g2 := models.CnvtMat2Edge(objMat)
	models.BFS("0", &g2)
	models.DFS("0", &g2)

}
