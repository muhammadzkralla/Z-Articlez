package main

import (
	"fmt"
	"runtime"
	"sync"
)

func GetSomeBitches(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		fmt.Println(i)
		runtime.Gosched()
	}
}
