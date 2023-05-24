package pkg

import (
	"testing"

	"github.com/fatih/color"
	"gotest.tools/assert"
)

func TestStripOffTermColors(t *testing.T) {
	tests := []struct {
		prompts []string
		want    []string
	}{

		{
			prompts: []string{
				color.GreenString("echo test"),
				color.RedString("echo test"),
				"echo test",
			},
			want: []string{
				"echo test",
				"echo test",
				"echo test",
			},
		},
	}
	for _, test := range tests {
		t.Run("colors reset test", func(t *testing.T) {
			got := NewPromptHandler(test.prompts).GetColoredPrompts()
			for i := range got {
				assert.Equal(t, test.want[i], got[i])
			}

		})
	}
}

func GetColoredPrompts(t *testing.T) {
	tests := []struct {
		prompts []string
		want    []string
	}{
		{
			prompts: []string{
				"echo test 1",
				"echo test 2",
				"echo test 3",
			},
			want: []string{
				color.GreenString("echo test 1"),
				color.RedString("echo test 2"),
				"echo test",
			},
		},
	}
	storage := NewStorageHandler()
	storage.Set("echo test 1", SUCCESS_PROMPT)
	storage.Set("echo test 2", ERRORED_PROMPT)

	for _, test := range tests {
		t.Run("colors reset test", func(t *testing.T) {
			got := NewPromptHandler(test.prompts).GetColoredPrompts()
			for i := range got {
				assert.Equal(t, test.want[i], got[i])
			}
		})
	}
}
