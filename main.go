package main

import (
	"fmt"
	"os"

	"github.com/kevincobain2000/re/pkg"
)

var version = "dev"

const README_PATH = "README.md"

func main() {

	tag := ""
	path := README_PATH
	for idx, arg := range os.Args {
		// check if arg is a URL
		if pkg.NewURLHandler().IsRemotePath(arg) {
			path = arg
			break
		}
		if arg == "v" || arg == "version" || arg == "-v" || arg == "--version" {
			printVersion()
			os.Exit(0)
		}
		if arg == "clear" {
			commands := pkg.NewReadmeHandler(path, tag).Codelines()
			pkg.NewPromptHandler(commands).ClearColoredPrompts()
			fmt.Println("Cleared command history")
			os.Exit(0)
		}
		if idx == 1 {
			tag = arg
		}
	}

	commands := pkg.NewReadmeHandler(path, tag).Codelines()

	prompts := pkg.NewPromptHandler(commands).GetColoredPrompts()
	prompts = pkg.NewPromptHandler(prompts).MultiSelect()
	err := pkg.NewPromptHandler(prompts).Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer pkg.DB().Close()
}

func printVersion() {
	fmt.Println(version)
}
