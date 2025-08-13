package main

import (
	"fmt"
	"sync"
)

var (
	array = []int{2, 4, 6, 8, 10}
	wg    sync.WaitGroup
)

func main() {
	wg.Add(len(array))
	for _, n := range array {
		go func(x int) {
			fmt.Println(x)
			defer wg.Done()
		}(n)
	}
	wg.Wait()
}
