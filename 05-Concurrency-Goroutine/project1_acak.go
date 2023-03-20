package main

import (
	"fmt"
	"sync"
)

func main() {
	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go printData(data1, data2, i, &wg)
	}

	wg.Wait()
}

func printData(data1 interface{}, data2 interface{}, i int, wg *sync.WaitGroup) {
	fmt.Println(data1, i)
	fmt.Println(data2, i)

	wg.Done()

}
