package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	nums            = 10
	randomNumsCount = 100
)

func main() {
	numChan := make(chan int, nums)
	squareChan := make(chan int, nums)
	var wg sync.WaitGroup

	wg.Add(1)
	go createSlice(nums, numChan, &wg)

	wg.Add(1)
	go calculateSquares(numChan, squareChan, &wg)

	wg.Wait()
	for value := range squareChan {
		fmt.Println(value)
	}
}

func createSlice(count int, dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		randomNumberInRange := rand.Intn(randomNumsCount)
		dataChan <- randomNumberInRange
	}
	close(dataChan)
}

func calculateSquares(inputChan chan int, outputChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range inputChan {
		squareNum(num, outputChan)
	}
	close(outputChan)
}

func squareNum(num int, dataChan chan int) {
	dataChan <- num * num
}
