package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Prompt will ask the user a question and returns the user's trimmed answer.
// If the user input is empty or is all spaces, then it returns the alt string.
func Prompt(question string, alt string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(text)
	if len(answer) == 0 {
		return alt
	}
	return answer
}
