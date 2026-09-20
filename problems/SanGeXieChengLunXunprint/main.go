package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string, 1)
	ch3 <- "a"

	go func() {
		for j := 1; j <= 10; j++ {
			k := <-ch1
			fmt.Print("第" + strconv.Itoa(j) + "次")
			fmt.Println("第2个协程：" + k)
			ch2 <- "c"
		}
		wg.Done()
	}()
	go func() {
		for j := 1; j <= 10; j++ {
			k := <-ch2
			fmt.Print("第" + strconv.Itoa(j) + "次")
			fmt.Println("第3个协程：" + k)
			ch3 <- "a"
		}
		wg.Done()
	}()
	go func() {
		for j := 1; j <= 10; j++ {
			k := <-ch3
			fmt.Print("第" + strconv.Itoa(j) + "次")
			fmt.Println("第1个协程：" + k)
			ch1 <- "b"
		}
		wg.Done()
	}()

	wg.Wait()
}
