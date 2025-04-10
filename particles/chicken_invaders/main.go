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
var inputState InputState
var screen [][]rune
var mu sync.Mutex

type Jett struct {
	x, y int
	char rune
}

type InputState struct {
	up    bool
	down  bool
	left  bool
	right bool
}

func createJett() Jett {
	return Jett{
		x:    w / 2,
		y:    h - 2,
		char: '█',
	}
}

func (j *Jett) updateJettPosition() {
	mu.Lock()

	dx, dy := 0, 0

	if inputState.up && !inputState.down {
		dy = -1
	} else if inputState.down && !inputState.up {
		dy = 1
	}

	if inputState.left && !inputState.right {
		dx = -1
	} else if inputState.right && !inputState.left {
		dx = 1
	}

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

		ev := termbox.PollEvent()
		if ev.Type != termbox.EventKey {
			continue
		}

		mu.Lock()

		switch ev.Key {
		case termbox.KeyArrowUp:
			inputState.up = (ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowUp)
		case termbox.KeyArrowDown:
			inputState.down = (ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowDown)
		case termbox.KeyArrowLeft:
			inputState.left = (ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowLeft)
		case termbox.KeyArrowRight:
			inputState.right = (ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowRight)
		case termbox.KeyEsc:
			mu.Unlock()
			log.Fatal("Quitting..")
			return
		}

		if ev.Key == termbox.KeyArrowUp && ev.Ch == 0 {
			inputState.up = false
		}
		if ev.Key == termbox.KeyArrowDown && ev.Ch == 0 {
			inputState.down = false
		}
		if ev.Key == termbox.KeyArrowLeft && ev.Ch == 0 {
			inputState.left = false
		}
		if ev.Key == termbox.KeyArrowRight && ev.Ch == 0 {
			inputState.right = false
		}

		mu.Unlock()

		jett.updateJettPosition()
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

	// Set input mode to receive keyboard input
	termbox.SetInputMode(termbox.InputEsc)

	jett = createJett()

	go render(&wg)
	go handleInput(&wg)

	wg.Wait()
}
