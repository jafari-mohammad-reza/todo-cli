package utils

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func GetInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal("failed to read input")
	}
	text := scanner.Text()
	if len(text) == 0 {
		log.Fatal("empty task title not allowed")
	}
	return text, nil
}
