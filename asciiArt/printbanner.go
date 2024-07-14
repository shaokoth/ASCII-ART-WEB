package asciiArt

import (
	"fmt"
	"strings"
)

// Print the banner for a line of text
func PrintLineBanner(line string, bannerMap map[int][]string) (string, error) {
	if line == "" {
		return "\n", nil
	}

	output := make([]string, 8)

	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			return "", fmt.Errorf("character %c not found in banner map", char)
		}

		for i := 0; i < 8; i++ {
			output[i] += banner[i]
		}
	}

	return strings.Join(output, "\n"), nil
}
