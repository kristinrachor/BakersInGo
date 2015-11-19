package main

import (
	"math/rand"
	"fmt"
	"time"
)

//creating customer struct. Each Customer has its own channel so it can get its result back
type Customer struct {
	sleep int
	fib int
	fibResult int
	retChan chan int
}

//serve customer function takes in the servers number (just for printing purposes) and takes in a customer channel to "return" the fib result. It will take a customer out of the line channel and computes its fib. 
func serveCust(num int, line chan Customer) {
	for {
		cust := <-line
		fmt.Printf("Server %d is calculating the fib of %d\n", num, cust.fib)
		cust.retChan <- fibonacci(cust.fib)
	}
}

//the customer function will sleep for given customer's sleep time and then add the customer to the line. The customers fibResult will be taken from the customers channel provided by the server. 
func custFunc(line chan Customer, cust Customer) {
	time.Sleep(time.Duration(cust.sleep) * time.Second)
	
	line <- cust
	cust.fibResult = <-cust.retChan
	fmt.Printf("Fibbonacci of %d is %d!\n", cust.fib, cust.fibResult)
	close(cust.retChan)
}

//standard fib function
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

//channel of customers waiting for help by a server
var line = make(chan Customer)

//main is called when you start the program. Loop to make servers and call go function "serveCust" and loop to make customer struct and then call "custFunc". Sleep is called for 20 seconds before it closes the line channel to insure all customers were helped. 
func main() {
	for i:= 0; i < 4; i++ {
		go serveCust(i, line)
	}

	for i:= 0; i < 25; i++ {
		c := Customer{
			fib: rand.Intn(40) + 10,
			fibResult: 0,
			sleep: rand.Intn(10),
			retChan: make(chan int),
		}
		go custFunc(line, c)
	}

	time.Sleep(time.Duration(20)*time.Second)

	close(line)
}
