package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const usage = "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(usage)
		return
	}

	var outputFile, text, bannerType string

	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "--output="):
			outputFile = strings.TrimPrefix(arg, "--output=")
		case strings.HasPrefix(arg, "--"):
			fmt.Println(usage)
			return
		case text == "":
			text = arg
		default:
			bannerType = arg
		}
	}

	if outputFile != "" && !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println(usage)
		return
	}

	if bannerType == "" {
		bannerType = "standard"
	}

	file, err := os.Open("banner/" + bannerType + ".txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
		return
	}

	result := GenerateAscii(text, content)

	if outputFile != "" {
		if err = os.WriteFile(outputFile, []byte(result), 0o644); err != nil {
			fmt.Println("error writing file", err)
		}
	} else {
		fmt.Print(result)
	}
}

func GenerateAscii(text string, content []string) string {
	newText := strings.Split(text, "\\n")
	var result strings.Builder

	for j, val := range newText {
		if val == "" {
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for i := 0; i < len(newText[j]); i++ {
				asciival := newText[j][i]
				start := int(asciival-32)*9 + 1
				result.WriteString(content[start+row])
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
