package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	w = 70
	h = 40
)

var jett Jett
var screen [][]rune
var mu sync.Mutex

type Jett struct {
	x, y int
	char rune
}

func createJett() Jett {
	return Jett{
		x:    w / 2,
		y:    h - 2,
		char: '█',
	}
}

func (j *Jett) updateJettPosition(dx, dy int) {
	mu.Lock()

	if j.x+dx >= 0 && j.y+dy >= 0 && j.x+dx < w && j.y+dy < h {
		j.x += dx
		j.y += dy
	}

	mu.Unlock()
}

func clearScreen() {
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

func updateJett() {
	screen[jett.y][jett.x] = jett.char
}

func render(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mu.Lock()

		clearScreen()
		updateJett()

		mu.Unlock()

		fmt.Print("\033[H\033[2J")

		for _, row := range screen {
			fmt.Println(string(row))
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func handleInput(wg *sync.WaitGroup) {
	defer wg.Done()

	for {

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				jett.updateJettPosition(0, -1)
			case termbox.KeyArrowDown:
				jett.updateJettPosition(0, 1)
			case termbox.KeyArrowLeft:
				jett.updateJettPosition(-1, 0)
			case termbox.KeyArrowRight:
				jett.updateJettPosition(1, 0)
			case termbox.KeyEsc:
				log.Fatal("Quitting..")
				return
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	jett = createJett()

	go render(&wg)
	go handleInput(&wg)

	wg.Wait()
}
