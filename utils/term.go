package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

func AskOnTerm(prompt string, suffix string) (string, error) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: prompt,
	})
	if err != nil {
		return suffix, err
	}
	defer rl.Close()
	rl.CaptureExitSignal()
	rl.WriteStdin([]byte(suffix))

	input, err := rl.Readline()
	if err != nil {
		return suffix, err
	}

	if input == "" {
		input = suffix
	}

	return input, nil
}
func GetIntFromTerm(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ":")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the input:", err)
		return GetIntFromTerm(prompt)
	}
	input = strings.TrimSpace(input)
	result, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return GetIntFromTerm(prompt)
	}
	return result
}

func PrintAsciiLogo() {
	fmt.Println(`
█████████████████████████████████████████████████████████████████████████████████████
	
██ ████████  █████  ██████   █████  ██   ██ ██ ███    ███  █████  ███████ ██    ██ ██ 
██    ██    ██   ██ ██   ██ ██   ██ ██  ██  ██ ████  ████ ██   ██ ██      ██    ██ ██ 
██    ██    ███████ ██   ██ ███████ █████   ██ ██ ████ ██ ███████ ███████ ██    ██ ██ 
██    ██    ██   ██ ██   ██ ██   ██ ██  ██  ██ ██  ██  ██ ██   ██      ██ ██    ██    
██    ██    ██   ██ ██████  ██   ██ ██   ██ ██ ██      ██ ██   ██ ███████  ██████  ██

████████████████████████████████████████-DL-█████████████████████████████████████████
 `)
}
