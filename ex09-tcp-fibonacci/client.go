package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	//"strconv"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8104")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, string(text) + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	}
}
