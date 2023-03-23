package main

import (
	"fmt"
	"sync"
)

func main() {
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	coba := []interface{}{"coba1", "coba2", "coba3"}
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}

	// GOROUTINE Random
	fmt.Println("GOROUTINE Random")
	wg.Add(8)
	for i := 1; i <= 4; i++ {
		go show_random(bisa, i, &wg)
		go show_random(coba, i, &wg)
	}
	wg.Wait()

	// GOROUTINE With Mutex (Not sure)
	fmt.Println("\nGOROUTINE with MUTEX")
	wg.Add(8)
	for i := 1; i <= 4; i++ {
		go show_ordered(bisa, i, &wg, &mtx)
		go show_ordered(coba, i, &wg, &mtx)
	}
	wg.Wait()

}

func show_random(word interface{}, i int, wg *sync.WaitGroup) {
	fmt.Println(word, i)
	wg.Done()
}

func show_ordered(word interface{}, i int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	fmt.Println(word, i)
	mtx.Unlock()
	wg.Done()
}
