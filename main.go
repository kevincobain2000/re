package main

import (
	"os"

	"github.com/mikhae1/execmd"

	"github.com/kevincobain2000/re/pkg"
	"github.com/manifoldco/promptui"
)

var version = "dev"

const README_PATH = "README.md"

func main() {
	versionCmd()

	readmePath := README_PATH
	if len(os.Args) > 1 && os.Args[1] != "" {
		readmePath = os.Args[1]
	}
	commands := pkg.NewReadmeHandler(readmePath).Codelines()

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

func versionCmd() {
	// check if -v or version or v in path
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "version" || os.Args[1] == "v") {
		println(version)
		os.Exit(0)
	}
}
