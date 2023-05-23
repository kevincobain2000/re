package pkg

import (
	"strings"
)

type LineHandler struct {
	languages []string
}

func NewLineHandler() *LineHandler {
	return &LineHandler{
		languages: []string{
			"sh",
			"bash",
			"zsh",
			"fish",
			"powershell",
		},
	}
}

func (h *LineHandler) splitString(str string) []string {
	lines := strings.Split(str, "\n")
	result := make([]string, 0)
	currentLine := ""

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if strings.HasSuffix(trimmedLine, "\\") {
			// Line ends with a backslash, append it to the current line
			currentLine += trimmedLine[:len(trimmedLine)-1] + "\\\n       "
		} else {
			// Line doesn't end with a backslash, add the current line to the result
			if currentLine != "" {
				result = append(result, currentLine+trimmedLine)
				currentLine = ""
			} else {
				result = append(result, trimmedLine)
			}
		}
	}

	return result
}

func (h *LineHandler) trim(line string) string {
	return strings.TrimSpace(line)
}
func (h *LineHandler) removePrompt(line string) string {
	if strings.HasPrefix(line, "$") || strings.HasPrefix(line, ">") {
		line = line[1:]
		line = strings.TrimSpace(line)
	}
	return line
}

func (h *LineHandler) startsWithLowercaseEnglishAlphabet(line string) bool {
	if len(line) == 0 {
		return false
	}
	return line[0] >= 'a' && line[0] <= 'z'
}

func (h *LineHandler) isLanguage(line string) bool {
	return h.contains(h.languages, line)
}

func (h *LineHandler) isComment(line string) bool {
	// check if line is a comment
	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
		return true
	}
	return false
}

// check if in array
func (h *LineHandler) contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// extract tag from code block line
// ```sh re -> re
// ```sh sql -> sql
func (h *LineHandler) extractTag(line string) string {
	// split by space
	parts := strings.Split(line, " ")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

// extract lang from code block line
// ```sh re -> sh
// ```sh sql -> sh
func (h *LineHandler) extractLang(line string) string {
	// split by space
	parts := strings.Split(line, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}
