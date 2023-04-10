package utils

import (
	"log"
	"sync"
)

func consumeChan[T any](mychan chan T, procNum int) []T {
	mySlice := make([]T, procNum)
	for i := 0; i < procNum; i++ {
		select {
		case msg := <-mychan:
			mySlice = append(mySlice, msg)
		}
	}
	close(mychan)
	return mySlice
}

func ThreadMap[T any](inputList []T, fn func(T) ResultError) []ResultError {

	wg := &sync.WaitGroup{}

	resErrChan := make(chan ResultError)

	procNum := 0
	for _, input := range inputList {
		wg.Add(1)
		procNum++
		go func(inputArg T) {
			defer wg.Done()
			resErrChan <- fn(inputArg)
		}(input)
	}

	resErrSlice := consumeChan(resErrChan, procNum)
	wg.Wait()

	return resErrSlice
}

func ThisIsAFunc(ss string) ResultError {
	log.Println(ss)
	return ResultError{ss, nil}
}
