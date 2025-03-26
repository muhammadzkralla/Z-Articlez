package main

import (
	"fmt"
	"math"
	"time"
)

const (
	w = 150
	h = 70
)

var sun Body
var planets []Body
var screen [][]rune

type Body struct {
	mass   float64
	char   rune
	x, y   float64
	vx, vy float64
}

func CreateBody(x, y, mass float64, char rune, vx, vy float64) Body {
	return Body{
		mass: mass,
		char: char,
		x:    x,
		y:    y,
		vx:   vx,
		vy:   vy,
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
	for i := range planets {
		Fx := CalculateFx(planets[i], sun)
		Fy := CalculateFy(planets[i], sun)

		Dvx := Fx / planets[i].mass
		Dvy := Fy / planets[i].mass

		planets[i].vx += Dvx
		planets[i].vy += Dvy
		planets[i].Update()

		if int(planets[i].x) >= 0 && int(planets[i].x) < w && int(planets[i].y) >= 0 && int(planets[i].y) < h {
			screen[int(planets[i].y)][int(planets[i].x)] = planets[i].char
		}
	}

	if int(sun.x) >= 0 && int(sun.x) < w && int(sun.y) >= 0 && int(sun.y) < h {
		screen[int(sun.y)][int(sun.x)] = sun.char
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

		time.Sleep(120 * time.Millisecond)
	}
}

func Launch() {
	sun = CreateBody(w/2, h/2, 100000, '@', 0, 0)

	planets = []Body{
		CreateBody(w/2+10, h/2, 3, 'M', 0, 2.5),
		CreateBody(w/2+15, h/2, 6, 'V', 0, 2.2),
		CreateBody(w/2+20, h/2, 10, 'E', 0, 2.0),
		CreateBody(w/2+25, h/2, 7, 'M', 0, 1.8),
		CreateBody(w/2+35, h/2, 200, 'J', 0, 1.4),
		CreateBody(w/2+45, h/2, 95, 'S', 0, 1.1),
		CreateBody(w/2+55, h/2, 40, 'U', 0, 0.9),
		CreateBody(w/2+65, h/2, 30, 'N', 0, 0.7),
	}
}

func main() {
	Launch()
	Render()
}
