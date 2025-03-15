package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const w = 150
const h = 40

var rockets []Rocket
var particles []Particle
var mu sync.Mutex
var screen [][]rune

type Particle struct {
	x, y   float64
	vx, vy float64
	life   int
	char   rune
}

type Rocket struct {
	x, y     float64
	vy       float64
	char     rune
	exploded bool
}

func CreateParticle(x float64, y float64) Particle {
	p := ChooseMyParticle()
	return Particle{
		x:    float64(x),
		y:    float64(y),
		vx:   rand.Float64(),
		vy:   rand.Float64(),
		life: 10 + rand.Intn(10),
		char: p,
	}
}

func (p *Particle) update() {
	p.x += p.vx
	p.y += p.vy
	p.vy += 0.1
	p.life--
}

func CreateRocket(x float64) Rocket {
	return Rocket{
		x:        x,
		y:        h - 1,
		vy:       -0.7,
		char:     '|',
		exploded: false,
	}
}

func (r *Rocket) Update() {
	r.y += r.vy
	if r.y <= 5+rand.Float64()*10 {
		r.exploded = true
	}
}

func ColorMyPencils() rune {
	tmp := fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s\033[0m", 255, 0, 0, 255, 0, 255, string('A'))

	return rune(tmp[0])
}

func ChooseMyParticle() rune {
	shapes := "@#$%^&*()_+-=./:;"
	idx := rand.Intn(len(shapes))
	return rune(shapes[idx])
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

func UpdateRockets() {
	for i := range rockets {
		r := &rockets[i]
		r.Update()
		if r.exploded {
			newParticles := make([]Particle, 50)
			for i := range newParticles {
				newParticles[i] = CreateParticle(r.x, r.y)
			}

			particles = append(particles, newParticles...)
		} else {
			if int(r.y) >= 0 && int(r.y) < h && int(r.x) >= 0 && int(r.x) < w {
				screen[int(r.y)][int(r.x)] = r.char
			}
		}
	}
}

func UpdateParticles() {
	for i := range particles {
		p := &particles[i]
		p.update()
		if p.life > 0 && int(p.y) >= 0 && int(p.y) < h && int(p.x) >= 0 && int(p.x) < w {
			screen[int(p.y)][int(p.x)] = p.char
		}
	}
}

func RipRockets() {
	var aliveRockets []Rocket
	for _, r := range rockets {
		if !r.exploded {
			aliveRockets = append(aliveRockets, r)
		}
	}
	rockets = aliveRockets
}

func RipParticles() {
	var aliveParticles []Particle
	for _, p := range particles {
		if p.life > 0 {
			aliveParticles = append(aliveParticles, p)
		}
	}

	particles = aliveParticles
}

func Render(wg *sync.WaitGroup) {
	defer wg.Done()

	for {

		mu.Lock()

		ClearScreen()
		UpdateRockets()
		UpdateParticles()
		RipRockets()
		RipParticles()

		mu.Unlock()

		fmt.Print("\033[H\033[2J")
		for _, row := range screen {
			fmt.Println(string(row))
		}

		time.Sleep(70 * time.Millisecond)
	}
}

func Launch(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mu.Lock()
		x := float64(rand.Intn(w))
		rockets = append(rockets, CreateRocket(x))
		mu.Unlock()

		time.Sleep(time.Duration(1+rand.Intn(4)) * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(6)

	go Launch(&wg)
	go Launch(&wg)
	go Launch(&wg)
	go Launch(&wg)
	go Launch(&wg)

	go Render(&wg)

	wg.Wait()
}
