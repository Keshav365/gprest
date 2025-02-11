package models

import (
	"fmt"
	"slices"
)

type tube struct {
	visited []string
	added   []string
}

func newTube(size int) tube {
	visited := make([]string, 0, size)
	added := make([]string, 0, size)
	q1 := tube{visited, added}
	return q1
}

func printPath(path []string) {
	for _, j := range path {
		if j == path[len(path)-1] {

			fmt.Printf("%v", j)
		} else {
			fmt.Printf("%v->", j)
		}

	}
	fmt.Printf("\n")
}

func enq(adj string, q *tube) {

	if slices.Contains(q.added, adj) || slices.Contains(q.visited, adj) {
	} else {
		q.added = append(q.added, adj)
	}
}

func dq(q *tube) string {
	first := q.added[0]
	// fmt.Println("Visited: ", q.visited)
	q.visited = append(q.visited, q.added[0])

	// fmt.Println("Adj: ", q.added)
	q.added = q.added[1:]
	return first
}

func push(adj string, st *tube) bool {
	check := false
	if slices.Contains(st.added, adj) || slices.Contains(st.visited, adj) {
	} else {
		check = true
		st.added = append(st.added, adj)
	}
	return check
}

func pop(st *tube) string {
	last := st.added[len(st.added)-1]
	// fmt.Println("Visited: ", st.visited)
	st.visited = append(st.visited, last)

	// fmt.Println("Adj: ", st.added)
	st.added = st.added[:len(st.added)-1]
	return last
}

func BFS(src string, g *Graph) []string {
	var bfscost float64
	gQ := newTube(len(g.gAdj))
	enq(src, &gQ)
	len := len(g.gAdj)
	for len > 0 {
		src = dq(&gQ)
		for _, i := range g.gAdj[src] {

			// if i != nil {
			// fmt.Printf("%T: %v %T: %v\n", k, k, i, i)
			for a, b := range i {
				// fmt.Printf("%T: %v %T: %v\n", a, a, b, b)
				enq(a, &gQ)
				// fmt.Println("cost: ", b)
				bfscost = bfscost + b
			}
			// }
		}
		// fmt.Println()
		len--
	}

	printPath(gQ.visited)
	fmt.Println("BFS Cost: ", bfscost)
	fmt.Println()
	return gQ.visited
}

func DFS(src string, g *Graph) []string {
	var dfscost float64
	gSt := newTube(len(g.gAdj))
	push(src, &gSt)
	for len(gSt.added) > 0 {
		src = pop(&gSt)
		for _, i := range g.gAdj[src] {
			for a, b := range i {
				if push(a, &gSt) {
					dfscost += b
				}
			}
		}
	}
	printPath(gSt.visited)
	fmt.Println("DFS Cost:", dfscost)
	return gSt.visited
}
