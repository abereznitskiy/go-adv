package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	NUMS_COUNT := 10
	numChan := make(chan int, NUMS_COUNT)
	squareChan := make(chan int, NUMS_COUNT)
	var wg sync.WaitGroup

	wg.Add(1)
	go createSlice(NUMS_COUNT, numChan, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numChan {
			squareNum(num, squareChan)
		}
		close(squareChan)
	}()

	wg.Wait()
	for value := range squareChan {
		fmt.Println(value)
	}
}

func createSlice(count int, dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		randomNumberInRange := rand.Intn(100)
		dataChan <- randomNumberInRange
	}
	close(dataChan)
}

func squareNum(num int, dataChan chan int) {
	dataChan <- num * num
}
