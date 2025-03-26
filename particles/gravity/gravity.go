package main

import (
	"fmt"
	"math"
	"time"
)

const (
	w = 150
	h = 40
)

var body1, body2 Body
var screen [][]rune

type Body struct {
	mass   float64
	char   rune
	x, y   float64
	vx, vy float64
}

func CreateBody(x, y, mass float64, char rune) Body {
	return Body{
		mass: mass,
		char: char,
		x:    x,
		y:    y,
		vx:   0,
		vy:   0,
	}
}

func (b *Body) Update() {
	b.x += b.vx
	b.y += b.vy
}

func ClearScreen() {
	screen = make([][]rune, h)
	for i := range screen {
		screen[i] = make([]rune, w)
		for j := range screen[i] {
			if i == 0 || i == h-1 {
				screen[i][j] = '─'
			} else if j == 0 || j == w-1 {
				screen[i][j] = '│'
			} else {
				screen[i][j] = ' '
			}
		}
	}

	screen[0][0] = '┌'
	screen[0][w-1] = '┐'
	screen[h-1][0] = '└'
	screen[h-1][w-1] = '┘'
}

func UpdateBodies() {
	Fx := CalculateFx(body1, body2)
	Fy := CalculateFy(body1, body2)

	Dvx1 := Fx / body1.mass
	Dvy1 := Fy / body1.mass

	Dvx2 := Fx / body2.mass
	Dvy2 := Fy / body2.mass

	body1.vx += Dvx1
	body1.vy += Dvy1

	body2.vx -= Dvx2
	body2.vy -= Dvy2

	body1.Update()
	body2.Update()

	if int(body1.x) >= 0 && int(body1.x) < w && int(body1.y) >= 0 && int(body1.y) < h {
		screen[int(body1.y)][int(body1.x)] = body1.char
	}

	if int(body2.x) >= 0 && int(body2.x) < w && int(body2.y) >= 0 && int(body2.y) < h {
		screen[int(body2.y)][int(body2.x)] = body2.char
	}
}

func CalculateDist(body1, body2 Body) float64 {
	x1 := body1.x
	x2 := body2.x
	y1 := body1.y
	y2 := body2.y

	dx2 := (x2 - x1) * (x2 - x1)
	dy2 := (y2 - y1) * (y2 - y1)

	return math.Sqrt(dx2 + dy2)
}

func CalculateForce(body1, body2 Body) float64 {
	G := 1e-3
	F := G * body1.mass * body2.mass
	D := max(CalculateDist(body1, body2), 1)

	return F / (D * D)
}

func CalculateFx(body1, body2 Body) float64 {
	F := CalculateForce(body1, body2)
	D := CalculateDist(body1, body2)
	Dx := body2.x - body1.x

	return F * (Dx / D)
}

func CalculateFy(body1, body2 Body) float64 {
	F := CalculateForce(body1, body2)
	D := CalculateDist(body1, body2)
	Dy := body2.y - body1.y

	return F * (Dy / D)
}

func Render() {
	for {
		ClearScreen()
		UpdateBodies()

		fmt.Print("\033[H\033[2J")
		for _, row := range screen {
			fmt.Println(string(row))
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func Launch() {
	sunMass := 100000
	earthMass := 10
	body1 = CreateBody(w/2, h/2, float64(sunMass), '@')
	body2 = CreateBody((2*w)/3, h/3, float64(earthMass), '*')

	body2.vy = 1.2
	body2.vx = 1.0
}

func main() {
	Launch()
	Render()
}
