package main

import (
	models "backup/models"
	"fmt"
)

func main() {



	objMat := [][]float64{
		{0, 4, -1, -1, -1, -1, -1, 8, -1},
		{4, 0, 8, -1, -1, -1, -1, 11, -1},
		{-1, 8, 0, 7, -1, 4, -1, -1, 2},
		{-1, -1, 7, 0, 9, 14, -1, -1, -1},
		{-1, -1, -1, 9, 0, 10, -1, -1, -1},
		{-1, -1, 4, 14, 10, 0, 2, -1, -1},
		{-1, -1, -1, -1, -1, 2, 0, 1, 6},
		{8, 11, -1, -1, -1, -1, 1, 0, 7},
		{-1, -1, 2, -1, -1, -1, 6, 7, 0},
	}

	fmt.Println()
	fmt.Printf("        (1)___8___(2)___7______(3)\n")
	fmt.Printf("       / |         | \\          |\\ \n")
	fmt.Printf("     4/  |         |  \\         | \\9\n")
	fmt.Printf("     /   |         |2  \\        |  \\ \n")
	fmt.Printf("    /    |         |    \\       |   \\ \n")
	fmt.Printf("  (0)    |11      (8)    \\4     |14 (4)\n")
	fmt.Printf("    \\    |         |      \\     |   / \n")
	fmt.Printf("     \\   |         |       \\    |  / \n")
	fmt.Printf("     8\\  |         |        \\   | /10\n")
	fmt.Printf("       \\ |         |         \\  |/ \n")
	fmt.Printf("       (7)___1____(6)___2_____(5)\n")
	fmt.Println()

	fmt.Println("***************************************************")
	fmt.Println("           Printing Adjacency Matrix: ")
	fmt.Println("***************************************************")
	models.PrintGraph(objMat)
	fmt.Println("***************************************************\n")

	g1 := models.CnvtMat2Edge(objMat)
	fmt.Println("***************************************************")
	fmt.Println("      Printing BFS Path for the Given Graph: ")
	fmt.Println("***************************************************")
	models.BFS("0", &g1)
	fmt.Println("***************************************************\n")

	fmt.Println("\n***************************************************")
	fmt.Println("      Printing DFS Path for the Given Graph: ")
	fmt.Println("***************************************************")
	models.DFS("0", &g1)
	fmt.Println("***************************************************\n")


	fmt.Println("\n***************************************************")
	fmt.Println("    Printing Dijkstra Path for the Given Graph: ")
	fmt.Println("***************************************************")
	g1.Dijkstra("0")
	fmt.Println("***************************************************\n")


	fmt.Println("\n***************************************************")
	fmt.Println("            Solving Water Jug Problem: ")
	fmt.Println("***************************************************")
	models.JugsProb(3,5,2)
	fmt.Println("***************************************************\n")

	fmt.Println("\n***************************************************")
	fmt.Println("    Solving Cannabal and Missionaries Problem: ")
	fmt.Println("***************************************************")
	models.BFSCM()
	fmt.Println("***************************************************\n")

	fmt.Println("\n***************************************************")
	fmt.Println("      Printing TSP Path for the Given Graph: ")
	fmt.Println("***************************************************")
	g1.TSP("0")
	fmt.Println("***************************************************\n")
	// WaterJugBFS(3,5,2)

}
