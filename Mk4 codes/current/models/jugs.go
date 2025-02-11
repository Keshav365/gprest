package models

import "fmt"

var capA int
var capB int
var goal int

type jugs struct {
	jugA int
	jugB int
	goal int
}

func makeJugs(a, b, c int) jugs {
	capA = a
	capB = b
	goal = c
	return jugs{0, 0, c}
}

func (j *jugs) filljugA() {
	j.jugA = capA
}
func (j *jugs) filljugB() {
	j.jugB = capB
}
func (j *jugs) emptyA() {
	j.jugA = 0
}
func (j *jugs) emptyB() {
	j.jugB = 0
}

func (j *jugs) movetoA() {
	if (j.jugA + j.jugB) >= capA {
		// fmt.Println("Jug A Now full, Jug B remains: ", j.jugB-(capA-j.jugA), "Gallons")
		j.jugB = j.jugB - (capA - j.jugA)
		j.jugA = capA
	} else {
		j.jugA = j.jugB + j.jugB
		j.jugB = 0
	}
}

func (j *jugs) movetoB() {
	if (j.jugA + j.jugB) >= capB {
		// fmt.Println("Jug B Now full, Jug A remains: ", j.jugA-(capB-j.jugB), "Gallons")
		j.jugA = j.jugA - (capB - j.jugB)
		j.jugB = capB
	} else {
		j.jugB = j.jugB + j.jugA
		j.jugA = 0
	}
}

func JugsProb(A, B, goal int) {
	var BigJug int
	var smallJug int

	if A > B {
		BigJug = A
		smallJug = B
	} else {
		smallJug = A
		BigJug = B
	}

	if goal%smallJug != 0 {
		if BigJug%smallJug == 0 {
			fmt.Printf("Error1: Big Jug is multiple of small Jug && Goal is not multiple of small Jug \n")
			return
		}
	}
	if goal > BigJug {
		fmt.Println("The remainder is more than the size of largest jug")
		return
	}

	j1 := makeJugs(smallJug, BigJug, goal)
	fmt.Println("JA:", j1.jugA, "JB:", j1.jugB)
	i := 100
	for j1.jugB != goal {
		if j1.jugB == capB {
			j1.emptyB()
		}
		if j1.jugA == 0 {
			j1.filljugA()
		}
		fmt.Println("JA:", j1.jugA, "JB:", j1.jugB)
		j1.movetoB()
		fmt.Println("JA:", j1.jugA, "JB:", j1.jugB)
		i--
	}
}
