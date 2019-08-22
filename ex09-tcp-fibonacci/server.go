package main

import (
	"net"
	"fmt"
	"bufio"
	//"strings"
	"strconv"
)

func fibonacci(n int) int {
	a0 := 0
	a1 := 1
	tmp := 0
	for i:=0; i < n; i++ {
		tmp = a0
		a0 = a1
		a1 = a1 + tmp
	}
	return a0
}

func main() {
	var output string
	var fib int
	ln, _ := net.Listen("tcp", ":8104")
	conn, _ := ln.Accept()
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message from client :%s lolol",string(message))
		n, err := strconv.Atoi(string(message))
		switch {
		case err!= nil:
			output = "Bad data"
			conn.Close()
		case n < 0:
			output = "Negative integer Error"
		default:
			fib = fibonacci(n)
			output = "F(" + strconv.Itoa(n) + ") = " + strconv.Itoa(fib)
		}
		conn.Write([]byte(output + "\n"))
	}
}




