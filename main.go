package main

import "fmt"

var logo = [5]string{
	"",
	"\033[38;5;198m▀ \033[38;5;213m▀ \033[38;5;141m▀ \033[38;5;117m▀\033[38;5;255m  █▀▀▃ █ ▃▀▀▃ █▃▃▃ █ █ █",
	"\033[38;5;198m▀ \033[38;5;255m█ \033[38;5;141m▀ \033[38;5;255m█\033[38;5;255m  █▃▃▀ █ █  █ █    █ ▀▃▀",
	"\033[38;5;255m█ \033[38;5;255m█ \033[38;5;255m█ \033[38;5;255m█\033[38;5;255m  █    █ ▀▃▃▀ ▀▃▃▃ █ ▃▀ ",
	"",
}

func PrintLogo(row, col int) {
	var output string
	for i, line := range logo {
		output += fmt.Sprintf("\033[%d;%dH%s", row+i, col, line)
	}
	fmt.Println(output)
}

func main() {
	PrintLogo(5, 5)
}
