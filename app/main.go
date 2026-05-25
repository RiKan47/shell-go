package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
// var _ = fmt.Print

func main() {
REPL:
	for {
		fmt.Print("$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}
		// cmd = strings.TrimSpace(cmd)
		args := strings.Fields(input)
		cmd := args[0]

		switch cmd {
		case "echo":
			for i := 1; i < len(args)-1; i++ {
				fmt.Print(args[i] + " ")
			}
			fmt.Println(args[len(args)-1])
		case "exit":
			break REPL
		default:
			fmt.Println(cmd + ": command not found")
		}
		// if cmd == "echo" {
		// 	// fmt.Println(cmd[5:])
		// 	for i := 1; i < len(args)-1; i++ {
		// 		fmt.Print(args[i] + " ")
		// 	}
		// 	fmt.Println(args[len(args)-1])
		// } else if cmd == "exit" {
		// 	break
		// } else {
		// 	fmt.Println(cmd + ": command not found")
		// }
	}
}
