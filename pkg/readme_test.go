package pkg

import (
	"testing"

	"gotest.tools/assert"
)

const RELATIVE_README_PATH = "../README.md"

func TestRead(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH).readLocal()
	assert.Assert(t, len(got) > 0)
}

func TestParseCodeBlocks(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH).parseCodeBlocks()
	assert.Assert(t, got != nil)
	assert.Assert(t, len(got) > 0)
}
func TestParseCodelines(t *testing.T) {
	got := NewReadmeHandler(RELATIVE_README_PATH).Codelines()
	assert.Assert(t, len(got) > 0)
}

func TestIsReadmePathURL(t *testing.T) {
	tests := []struct {
		url string
		got bool
	}{
		{
			url: "README.md",
			got: false,
		},
		{
			url: "https://github.com/kevincobain2000/re",
			got: true,
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewReadmeHandler(test.url).isReadmePathURL()
			assert.Equal(t, test.got, got)

		})
	}
}
func TestConvertGithubURL(t *testing.T) {
	tests := []struct {
		url string
		got string
	}{
		{
			url: "https://github.com/kevincobain2000/re",
			got: "https://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
		{
			url: "https://github.com/kevincobain2000/re/",
			got: "https://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
		{
			url: "http://github.com/kevincobain2000/re/",
			got: "http://raw.githubusercontent.com/kevincobain2000/re/master/README.md",
		},
	}
	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			got := NewReadmeHandler(test.url).convertGithubURL()
			assert.Equal(t, test.got, got)

		})
	}
}
