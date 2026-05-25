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
	for {
		fmt.Print("$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}
		cmd = strings.TrimSpace(cmd)

		// switch cmd {
		// case "echo":
		// 	fmt.Println()
		// 	case "exit"
		// }
		if strings.HasPrefix(cmd, "echo") {
			fmt.Println(cmd[5:])
		} else if cmd == "exit" {
			break
		} else {
			fmt.Println(cmd + ": command not found")
		}
	}
}
