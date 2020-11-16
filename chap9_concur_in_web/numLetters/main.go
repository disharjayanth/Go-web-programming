package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers1() {
	defer fmt.Printf("***********\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}

func printLetters1() {
	defer fmt.Printf("***********\n")
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c", i)
	}
}

func printNumbers2(wg *sync.WaitGroup) {
	defer fmt.Printf("***********\n")
	defer wg.Done()
	for i := 0; i < 65; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
}

func printLetters2(wg *sync.WaitGroup) {
	defer fmt.Printf("***********\n")
	defer wg.Done()
	for i := 'A'; i < 'A'+65; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	printNumbers1()
	printLetters1()
	go printLetters2(&wg)
	go printNumbers2(&wg)
	wg.Wait()
}
