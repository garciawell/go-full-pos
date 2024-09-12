package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 1
	}()

	for {
		select {
		case ms1 := <-c1: //rabit mq
			println("c1 received", ms1)
		case ms2 := <-c2: //rabit kafka
			println("c2 received", ms2)

		case <-time.After(time.Second * 23):
			println("timeout")

		default:
			println("default")
		}
	}

}
