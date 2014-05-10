package main

import "fmt"
import "sync"

func Par(processes ...*func()) {
	var wg sync.WaitGroup
	wg.Add(len(processes))
	for _, v := range processes {
		go func(v func()) {
			defer wg.Done()
			v()
		}(*v)
	}
	wg.Wait()
}

func Producer(c chan int) {
			c <- 2
}

func Consumer(c chan int) {
		v := <-c
		fmt.Println(v)
}


func main() {
//	c := make(chan int)
//  prod := Producer 
//	con := Consumer
//	Par( &prod(c), &con(c)	)
}
