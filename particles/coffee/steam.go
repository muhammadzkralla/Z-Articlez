package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const w = 60
const h = 15

var shades = []rune{'░', '░', '▒', '▒', '▓', '▓', '█', '█'}
var particles []Particle
var mu sync.Mutex
var screen [][]rune

type Particle struct {
	x, y float64
	vy   float64
	life int
	char rune
}

func NewParticle() Particle {
	prop := rand.NormFloat64() * 6
	return Particle{
		x:    33 + prop,
		y:    float64(h - 1),
		vy:   -1,
		life: 3 + rand.Intn(len(shades)-1),
		char: shades[len(shades)-1],
	}
}

func (p *Particle) Update() {
	p.y += p.vy
	p.life--

	if p.life >= 0 && p.life < len(shades) {
		p.char = shades[p.life]
	}
}

func ClearScreen() {
	screen = make([][]rune, h)
	for i := range screen {
		screen[i] = make([]rune, w)
		for j := range screen[i] {
			screen[i][j] = ' '
		}
	}
}

func UpdateParticles() {
	for i := range particles {
		p := &particles[i]
		p.Update()
		if p.life > 0 && int(p.x) >= 0 && int(p.x) < w && int(p.y) >= 0 && int(p.y) < h {
			screen[int(p.y)][int(p.x)] = p.char
		}
	}
}

func Rip() {
	var aliveParticles []Particle
	for _, p := range particles {
		if p.life > 0 {
			aliveParticles = append(aliveParticles, p)
		}
	}

	particles = aliveParticles
}

func Render(wg *sync.WaitGroup, ansiImage []string) {
	defer wg.Done()

	for {
		mu.Lock()

		ClearScreen()
		UpdateParticles()
		Rip()

		mu.Unlock()

		fmt.Print("\033[H\033[2J")

		for _, row := range screen {
			fmt.Println(string(string(row)))
		}

		for _, row := range ansiImage {
			fmt.Println(row)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func Launch(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var newParticles = make([]Particle, 100)
		for i := range newParticles {
			newParticles[i] = NewParticle()
		}

		mu.Lock()
		particles = append(particles, newParticles...)
		mu.Unlock()

		time.Sleep(100 * time.Millisecond)
	}

}
