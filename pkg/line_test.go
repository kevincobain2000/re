package pkg

import (
	"testing"

	"gotest.tools/assert"
)

func TestTrim(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{
			line: " Hello world ",
			want: "Hello world",
		},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			got := NewLineHandler().trim(test.line)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestRemovePrompt(t *testing.T) {
	tests := []struct {
		line string
		want string
	}{
		{
			line: "$ echo hello",
			want: "echo hello",
		},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			got := NewLineHandler().removePrompt(test.line)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStartsWithLowercaseEnglishAlphabet(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{
			line: "Hello world",
			want: false,
		},
		{
			line: "hello world",
			want: true,
		},
		{
			line: "123 hello world",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			got := NewLineHandler().startsWithLowercaseEnglishAlphabet(test.line)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestIsanguage(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{
			line: "sh",
			want: true,
		},
		{
			line: "hello world",
			want: false,
		},
		{
			line: "123 hello world",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			got := NewLineHandler().isLanguage(test.line)
			assert.Equal(t, test.want, got)
		})
	}
}
