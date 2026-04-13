// Functions for taking input
package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetString: gets input as a string
func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error while reading string: %v", err)
	}

	return strings.TrimSpace(input), nil
}
