package pkg

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/mikhae1/execmd"
)

type PromptHandler struct {
	prompts []string
}

func NewPromptHandler(prompts []string) *PromptHandler {
	return &PromptHandler{
		prompts: prompts,
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
		if err := h.run(prompt); err != nil {
			return err
		}
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
	survey.AskOne(prompt, &prompts)
	return prompts
}
