package main

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth  = 800
	winHeight = 600
	fps       = 60
)

func drawCircle(renderer *sdl.Renderer, centerX, centerY, radius int32, r, g, b uint8) {
	renderer.SetDrawColor(r, g, b, 255)

	for w := int32(0); w < radius*2; w++ {
		for h := int32(0); h < radius*2; h++ {
			dx := radius - w
			dy := radius - h
			if (dx*dx + dy*dy) <= (radius * radius) {
				renderer.DrawPoint(centerX+dx, centerY+dy)
			}
		}
	}
}

func InitSdl() {
	fmt.Println("Initializing SDL...")
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("SDL Init failed: %v", err)
	}
	defer sdl.Quit()

	fmt.Println("Creating window...")
	window, err := sdl.CreateWindow("Solar System",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Window creation failed: %v", err)
	}
	defer window.Destroy()

	fmt.Println("Creating renderer...")
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		log.Fatalf("Renderer creation failed: %v", err)
	}
	defer renderer.Destroy()

	fmt.Println("Rendering solar system...")

	running := true
	for running {
		startTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// Get updated window size
		currentWidth, currentHeight := window.GetSize()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		drawCircle(renderer, currentWidth/2, currentHeight/2, 50, 255, 215, 0)

		drawCircle(renderer, currentWidth/2-150, currentHeight/2-100, 20, 0, 0, 255)

		drawCircle(renderer, currentWidth/2+200, currentHeight/2+80, 30, 255, 0, 0)

		renderer.Present()

		elapsed := time.Since(startTime)
		frameTime := time.Second / fps
		if elapsed < frameTime {
			sdl.Delay(uint32((frameTime - elapsed).Milliseconds()))
		}
	}
}
