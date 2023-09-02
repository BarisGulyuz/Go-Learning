package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//goroutine()

	//channels()

	channels2()
}

func goroutine() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("User Created")

	go func() {
		defer wg.Done()
		fmt.Println("Email sent")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("SMS sent")
	}()

	wg.Wait()
}

func channels() {
	channel := make(chan string)

	go func() {
		channel <- "Mesaj Geldi"
	}()

	message := <-channel
	fmt.Println(message)

	channel2 := make(chan bool, 3)

	go func() {
		channel2 <- true

		time.Sleep(1 * time.Second)

		channel2 <- false
	}()

	fmt.Println(<-channel2, <-channel2)
}

func channels2() {

	var wg sync.WaitGroup

	ch := make(chan int, 2)

	go func() {
		defer wg.Done()
		ch <- 1
		time.Sleep(1 * time.Second)
		ch <- 2
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Println("Gelen veri:", data)
		}
	}()

	wg.Add(2)
	wg.Wait()

}
