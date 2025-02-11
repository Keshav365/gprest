package models

import (
	"fmt"
	"strconv"
)

type Graph struct {
	gAdj map[string][]map[string]float64
}

func NewGraph() Graph {
	return Graph{gAdj: make(map[string][]map[string]float64)}
}

func (g *Graph) AddVertex(n string) {
	g.gAdj[n] = append(g.gAdj[n])
}

func (g *Graph) AddEdge(src, dest string, w float64) {

	if _, exists := g.gAdj[src]; !exists {
		g.AddVertex(src)
	}

	if _, exists := g.gAdj[dest]; !exists {
		g.AddVertex(dest)
	}

	edge := map[string]float64{dest: w}
	g.gAdj[src] = append(g.gAdj[src], edge)
}

func (g *Graph) GetVertex(n string) []map[string]float64 {
	return g.gAdj[n]
}
func (g *Graph) GetVertices() []string {

	var keyList []string
	for k := range g.gAdj {
		keyList = append(keyList, k)
	}
	return keyList
}
func PrintGraph(objMat [][]float64) {
	topper := len(objMat)
	fmt.Println()
	fmt.Printf(" _V")
	for i := 0; i < topper; i++ {
		fmt.Printf("|_%v_", i)
	}
	fmt.Println()
	for i, j := range objMat {
		fmt.Printf(" _%v", i)
		for a, _ := range j {
			if a == i {
				fmt.Printf("| 0 ")
			} else {
				if j[a] == -1 {
					fmt.Printf("|___")
				} else if j[a]/10 >= 1 {
					fmt.Printf("|%v ", j[a])
				} else if j[a]/10 == float64(0) {
					fmt.Printf("|-1 ")
				} else if j[a]/10 < 1 {
					fmt.Printf("| %v ", j[a])
					// fmt.Printf("| %v ", j[a])
				}
			}
			fmt.Printf("")
		}
		fmt.Println()
	}
	fmt.Println()
}

func CnvtMat2Edge(objMat [][]float64) Graph {

	g := NewGraph()
	// fmt.Println(objMat)
	for i, j := range objMat {
		for a, b := range j {
			// fmt.Printf("%v %v ", a, i)
			if j[a] == -1 || j[a] == 0 {
			} else {
				g.AddEdge(strconv.Itoa(i), strconv.Itoa(a), b)
			}
		}
		// fmt.Println()
	}
	return g
}
