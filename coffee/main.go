package main

import (
	"os"
	"sync"
)

func main() {
	imagePath := os.Args[1]
	img := OpenImage(imagePath)
	ansiImage := ConvertImage(img)

	var wg sync.WaitGroup
	wg.Add(1)

	go Render(&wg, ansiImage)
	go Launch(&wg)

	wg.Wait()
}
