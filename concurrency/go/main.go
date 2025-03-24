package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go GetSomeBitches(&wg)
	go GetSomeBitches(&wg)
	wg.Wait()
}
