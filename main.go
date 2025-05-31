package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/basemax/brain-fuck-runner/bf"
)

func main() {
	fmt.Println("Brainfuck REPL (type `exit` or Ctrl+C to quit)")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		code = strings.TrimSpace(code)
		if code == "exit" {
			break
		}

		interpreter := bf.New(code)
		output, err := interpreter.Execute()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if output != "" {
			fmt.Printf("Output: %s\n", output)
		}
	}
}
