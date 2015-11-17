package main

import (
	"math/rand"
	"fmt"
)

type Customer struct {
	fib int
	fibResult int
}

type Server struct {
	
}

//func serveCust {

//} 

func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}	

var line = make(chan Customer)
var servers = make(chan Server)
var results = make(chan Customer)

func main() {
	for i:= 0; i < 3; i++ {
		go func() {servers <- Server{}}()
		
	}

	for i:= 0; i < 5; i++ {
		c := Customer{fib: rand.Intn(40) + 10, fibResult: 0}
		c.fibResult = fibonacci(c.fib)
		go func() {line <- c}()
	}


	for i:= 0; i < 5; i++ {
		c2 := <-line
	
		fmt.Printf("The result of fib(%d) is %d!\n", c2.fib, c2.fibResult)
	}
	close(line)
	close(servers)
	close(results)
	fmt.Printf("blah");


}

