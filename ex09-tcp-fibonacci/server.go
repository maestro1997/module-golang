package main

import (
	"net"
	"fmt"
	"bufio"
	"strconv"
	"time"
	"math/big"
)

func fibonacci(n int) *big.Int {
	a0 := big.NewInt(0)
	a1 := big.NewInt(1)
	for i := 0; i < n; i++ {
		a0, a1 = a1, a0.Add(a0, a1)
	}
	return a0
}

func main() {
	var input string
	var output *big.Int
	cache := make(map[string]*big.Int)
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8412")
	conn, _ := ln.Accept()
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		input = string(message)
		input = input[:len(input) - 1]	
		elem, ok := cache[input]
		t1 := time.Now()
		if ok {
			output = elem
		} else {
			n, _ := strconv.Atoi(input)
			output = fibonacci(n)
		}
		time := time.Since(t1)

		conn.Write([]byte(time.String() + " " +output.String() + "\n"))
		if !ok {
			cache[input] = output
		}
	}
}
