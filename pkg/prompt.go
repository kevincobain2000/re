package pkg

import (
	"fmt"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/mikhae1/execmd"
)

const (
	SUCCESS_PROMPT = "success"
	ERRORED_PROMPT = "errored"
)

type PromptHandler struct {
	prompts        []string
	storageHandler *StorageHandler
}

func NewPromptHandler(prompts []string) *PromptHandler {
	return &PromptHandler{
		prompts:        prompts,
		storageHandler: NewStorageHandler(),
	}
}

func (h *PromptHandler) run(prompt string) error {
	// execute prompt
	cmd := execmd.NewCmd()
	cmd.PrefixStderr = ""
	cmd.PrefixStdout = ""
	_, err := cmd.Run(prompt)
	return err
}

func (h *PromptHandler) Execute() error {
	for _, prompt := range h.prompts {
		prompt := h.stripOffTermColors(prompt)
		if err := h.run(prompt); err != nil {
			h.storageHandler.Set(prompt, ERRORED_PROMPT)
			return err
		}
		h.storageHandler.Set(prompt, SUCCESS_PROMPT)
	}
	return nil
}

func (h *PromptHandler) MultiSelect() []string {
	prompts := []string{}
	prompt := &survey.MultiSelect{
		Message:  "Select:",
		Options:  h.prompts,
		PageSize: 20,
	}
	err := survey.AskOne(prompt, &prompts)
	if err != nil {
		fmt.Print(err.Error())
		return []string{}
	}
	return prompts
}

// stripOffTermColors patterns like example:
// Remove \x1b[32m from beginning of string and \x1b[0m from end of string
func (h *PromptHandler) stripOffTermColors(str string) string {
	// Remove pattern from the beginning of the string
	pattern := `\x1b\[\d+m`
	re := regexp.MustCompile(pattern)
	str = re.ReplaceAllString(str, "")

	// Remove pattern from the end of the string
	pattern = `\x1b\[\dm$`
	re = regexp.MustCompile(pattern)
	str = re.ReplaceAllString(str, "")
	return str
}

func (h *PromptHandler) GetColoredPrompts() []string {
	prompts := []string{}
	for _, prompt := range h.prompts {
		v, err := h.storageHandler.Get(prompt)
		if err != nil {
			prompts = append(prompts, prompt)
			continue
		}
		if v == SUCCESS_PROMPT {
			prompts = append(prompts, color.GreenString(prompt))
		}
		if v == ERRORED_PROMPT {
			prompts = append(prompts, color.RedString(prompt))
		}
	}
	return prompts
}

func (h *PromptHandler) ClearColoredPrompts() {
	for _, prompt := range h.prompts {
		h.storageHandler.Delete(prompt)
	}
}
