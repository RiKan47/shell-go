package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
// var _ = fmt.Print

// type helper func(string) (bool, error)

func main() {

	cmdDb := map[string]bool{
		// "type": func(s string) (bool, error){
		// 	fmt.Println()
		// }
		"type": true,
		"echo": true,
		"exit": true,
	}

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
		// case "type":
		case "type":
			if _, ok := cmdDb[args[1]]; ok {
				fmt.Println(args[1] + " is a shell builtin")
			} else {
				// dirs := os.Args
				if path, _ := exec.LookPath(args[1]); path != "" {
					fmt.Println(args[1] + " is " + path)
				} else {
					fmt.Println(args[1] + ": not found")
				}
			}
		default:
			if path, _ := exec.LookPath((cmd)); path != "" {
				userCmd := exec.Command(cmd, args[1:]...)
				out, err := userCmd.Output()
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				fmt.Print(string(out))
			} else {
				fmt.Println(cmd + ": command not found")
			}
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
