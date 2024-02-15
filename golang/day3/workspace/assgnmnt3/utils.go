package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Spaces(count int) string {
	return strings.Repeat(" ", count)
}

func Chars(char byte, count int) string {
	return strings.Repeat(string(char), count)
}

func Input(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
