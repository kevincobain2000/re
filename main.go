package main

import (
	"github.com/mikhae1/execmd"

	"github.com/kevincobain2000/re/pkg"
	"github.com/manifoldco/promptui"
)

const README_PATH = "README.md"

func main() {
	commands := pkg.NewReadmeHandler(README_PATH).Codelines()

	prompt := promptui.Select{
		Label: "Choose command [ctrl+c to exit]:",
		Items: commands,
		Size:  25,
	}

	_, command, err := prompt.Run()

	if err != nil {
		return
	}
	// execute command
	cmd := execmd.NewCmd()
	cmd.PrefixStderr = ""
	cmd.PrefixStdout = ""
	_, err = cmd.Run(command)
	if err != nil {
		return
	}
}
