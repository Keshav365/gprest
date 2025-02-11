package main

import (
	models "backup/models"
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
		{0, 4, -1, -1, -1, -1, -1, 8, -1},  // A
		{4, 0, 8, -1, -1, -1, -1, 11, -1},  // B
		{-1, 8, 0, 7, -1, 4, -1, -1, 2},    // C
		{-1, -1, 7, 0, 9, 14, -1, -1, -1},  // D
		{-1, -1, -1, 9, 0, 10, -1, -1, -1}, // E
		{-1, -1, 4, 14, 10, 0, 2, -1, -1},  // F
		{-1, -1, -1, -1, -1, 2, 0, 1, 6},   // G
		{8, 11, -1, -1, -1, -1, 1, 0, 7},   // H
		{-1, -1, 2, -1, -1, -1, 6, 7, 0},   // I
	}

	// objMat := [][]float64{
	// 	{0, 4, -1, 5, -1, -1, -1},
	// 	{-1, 0, 8, -1, -1, 9, -1},
	// 	{-1, -1, 0, -1, 1, -1, 1},
	// 	{-1, -1, -1, 0, -1, 2, -1},
	// 	{-1, 1, -1, -1, 0, 1, -1},
	// 	{1, -1, -1, -1, -1, 0, -1},
	// 	{-1, -1, -1, -1, -1, 1, 0},
	// }

	models.PrintGraph(objMat)
	g2 := models.CnvtMat2Edge(objMat)
	// fmt.Println(g2)
	models.BFS("0", &g2)
	models.DFS("0", &g2)

}
