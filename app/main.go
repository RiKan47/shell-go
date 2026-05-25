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
		if cmd == "exit" {
			break
		}
		fmt.Println(cmd + ": command not found")
	}
}
