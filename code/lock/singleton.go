package main

import (
	"fmt"
	"sync"
)

type singleton struct {}

var (
	instance *singleton
	once sync.Once
)

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func worker(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fooSingleton := getInstance()
	fmt.Printf("%#v\n", fooSingleton)
}

func main() {
	fmt.Println("Get Singleton in Multi Thread")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go worker(&waitGroup)
	go worker(&waitGroup)

	waitGroup.Wait()
}
