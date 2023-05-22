package pkg

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/carlmjohnson/requests"
)

var languages = []string{
	"sh",
	"bash",
	"zsh",
	"fish",
	"powershell",
}

type ReadmeHandler struct {
	ReadmePath       string
	DefaultURLBranch string
}

func NewReadmeHandler(readmePath string) *ReadmeHandler {
	return &ReadmeHandler{
		ReadmePath:       readmePath,
		DefaultURLBranch: "master",
	}
}

func (h *ReadmeHandler) readLocal() []byte {
	file, err := os.Open(h.ReadmePath)
	if err != nil {
		return []byte("")
	}
	defer func() {
		if err := file.Close(); err != nil {
			// do nothing
			_ = err
		}
	}()

	contents, err := os.ReadFile(h.ReadmePath)
	if err != nil {
		return []byte("")
	}
	return contents
}
func (h *ReadmeHandler) isReadmePathURL() bool {
	return strings.HasPrefix(h.ReadmePath, "http")
}
func (h *ReadmeHandler) convertGithubURL() string {
	url := h.ReadmePath

	// if url doesn't end with / then append /
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	// replace github.com with raw.githubusercontent.com
	url = strings.Replace(url, "github.com", "raw.githubusercontent.com", 1)
	url += h.DefaultURLBranch + "/" + "README.md"
	return url
}

func (h *ReadmeHandler) readURL() []byte {
	var s string

	// possibly check here for /tree/branch/ in url
	// for now check for default branches
	branches := []string{
		"main",
		"master",
		"develop",
	}

	for idx, branch := range branches {
		h.DefaultURLBranch = branch

		url := h.convertGithubURL()
		ctx := context.Background()
		err := requests.
			URL(url).
			ToString(&s).
			Fetch(ctx)
		if err == nil {
			break
		}
		// check if last
		if idx == len(branches)-1 {
			fmt.Println("Unable to fetch README.md from github.com")
			return []byte("")
		}
	}
	return []byte(s)
}

func (h *ReadmeHandler) parseCodeBlocks() [][]string {
	var readmeContents []byte
	if h.isReadmePathURL() {
		readmeContents = h.readURL()
	} else {
		readmeContents = h.readLocal()
	}

	// Extract code blocks using regular expressions
	regex := regexp.MustCompile("(?s)```(.*?)```")
	codeBlocks := regex.FindAllStringSubmatch(string(readmeContents), -1)
	return codeBlocks
}

func (h *ReadmeHandler) Codelines() []string {
	codeLines := make([]string, 0)
	codeBlocks := h.parseCodeBlocks()
	for _, match := range codeBlocks {
		lines := match[1]
		for idx, line := range strings.Split(lines, "\n") {
			// trim spaces
			line = strings.TrimSpace(line)

			// ```sh is the first line of the code block, sh is extracted from it
			// check if this can be interpreted as a language
			if idx == 0 && !h.isLanguage(line) {
				break // out of entire code block
			}
			// check if line has a length
			if len(line) == 0 {
				continue
			}
			if h.isComment(line) {
				continue
			}
			if h.startsWithUpperCase(line) {
				continue
			}
			if h.startsWithNumeric(line) {
				continue
			}
			if !h.startsWithEnglish(line) {
				continue
			}

			line := h.removePrompt(line)

			if idx > 0 {
				codeLines = append(codeLines, line)
			}
		}
	}
	return codeLines
}

func (h *ReadmeHandler) removePrompt(line string) string {
	if strings.HasPrefix(line, "$") || strings.HasPrefix(line, ">") {
		line = line[1:]
		line = strings.TrimSpace(line)
	}
	return line
}

func (h *ReadmeHandler) startsWithUpperCase(line string) bool {
	if len(line) == 0 {
		return false
	}
	return line[0] >= 'A' && line[0] <= 'Z'
}
func (h *ReadmeHandler) startsWithEnglish(line string) bool {
	if len(line) == 0 {
		return false
	}
	return line[0] >= 'a' && line[0] <= 'z'
}

func (h *ReadmeHandler) startsWithNumeric(line string) bool {
	if len(line) == 0 {
		return false
	}
	return line[0] >= '0' && line[0] <= '9'
}

func (h *ReadmeHandler) isLanguage(line string) bool {
	return h.contains(languages, line)
}

func (h *ReadmeHandler) isComment(line string) bool {
	// check if line is a comment
	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
		return true
	}
	return false
}

// check if in array
func (h *ReadmeHandler) contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
