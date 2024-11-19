package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error to connect to server:", err)
		return
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Print("Enter text: ")
		text := input.Text()
		switch text {
			case "":
				continue
			case "/clear":
				fmt.Print("\033[H\033[2J")
			case "/quit":
				fmt.Println("Quitting...")
			default:
				fmt.Fprintln(conn, text)
		}
		fmt.Println("<Connected to server>")
		fmt.Println("Enter text: ")
	}
}
