package pkg

import (
	"testing"

	"gotest.tools/assert"
)

const RELATIVE_README_PATH = "../README.md"

func TestRead(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH, "").readLocal()
	assert.Assert(t, len(got) > 0)
}

func TestParseCodeBlocks(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH, "").parseCodeBlocks()
	assert.Assert(t, got != nil)
	assert.Assert(t, len(got) > 0)
}
func TestParseCodelines(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH, "").Codelines()
	assert.Assert(t, len(got) > 0)
}
