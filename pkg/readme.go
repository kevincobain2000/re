package pkg

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/carlmjohnson/requests"
)

type ReadmeHandler struct {
	ReadmePath  string
	Tag         string
	urlHandler  *URLHandler
	lineHandler *LineHandler
}

func NewReadmeHandler(readmePath string, tag string) *ReadmeHandler {
	return &ReadmeHandler{
		ReadmePath:  readmePath,
		Tag:         tag,
		urlHandler:  NewURLHandler(),
		lineHandler: NewLineHandler(),
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

func (h *ReadmeHandler) readRemote() []byte {
	// readme contents from URL
	var contents string

	if h.urlHandler.isFullReadmeURL(h.ReadmePath) {
		u := h.urlHandler.githubBlobURLToRawReadmeURL(h.ReadmePath)
		ctx := context.Background()
		err := requests.
			URL(u).
			ToString(&contents).
			Fetch(ctx)
		if err != nil {
			return []byte("")
		}
		return []byte(contents)
	}

	for idx, branch := range h.urlHandler.defaultBranches {
		u := h.urlHandler.githubBranchedURLToRawReadmeURL(h.ReadmePath, branch)
		ctx := context.Background()
		err := requests.
			URL(u).
			ToString(&contents).
			Fetch(ctx)
		if err == nil {
			break
		}
		// check if last
		if idx == len(h.urlHandler.defaultBranches)-1 {
			fmt.Println("Unable to fetch README.md from github.com")
			return []byte("")
		}
	}
	return []byte(contents)
}

func (h *ReadmeHandler) parseCodeBlocks() [][]string {
	var readmeContents []byte
	if h.urlHandler.IsRemotePath(h.ReadmePath) {
		readmeContents = h.readRemote()
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
		block := match[1]

		// split by \n except if line ends with \
		lines := h.lineHandler.splitString(block)
		if len(lines) == 0 {
			continue
		}

		lineFirst := lines[0]
		lineFirst = h.lineHandler.trim(lineFirst)
		tag := h.lineHandler.extractTag(lineFirst)
		lang := h.lineHandler.extractLang(lineFirst)
		if h.Tag == "" {
			if !h.lineHandler.isLanguage(lang) {
				continue // out of entire code block to next
			}
		}

		if h.Tag != "" {
			if h.Tag != tag && h.Tag != lang {
				continue
			}
		}

		// after 1st line
		for _, line := range lines[1:] {
			// trim spaces
			line = h.lineHandler.trim(line)

			// check if line has a length
			if len(line) == 0 {
				continue // out of this line to next
			}
			if h.lineHandler.isComment(line) {
				continue
			}
			line := h.lineHandler.removePrompt(line)
			if !h.lineHandler.startsWithLowercaseEnglishAlphabet(line) {
				continue
			}

			codeLines = append(codeLines, line)
		}
	}
	return codeLines
}
