package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kevincobain2000/re/pkg"
)

var version = "dev"

const README_PATH = "README.md"

func main() {
	flagVersion := flag.Bool("v", false, "show version")
	flagTag := flag.String("t", "sh", "tag to only show commands in code block with this tag")
	flag.Parse()

	if *flagVersion {
		printVersion()
	}

	path := README_PATH
	for _, arg := range os.Args {
		// check if arg is a URL
		if pkg.NewURLHandler().IsRemotePath(arg) {
			path = arg
			break
		}
	}
	commands := pkg.NewReadmeHandler(path, *flagTag).Codelines()
	prompts := pkg.NewPromptHandler(commands).MultiSelect()
	pkg.NewPromptHandler(prompts).Execute()
}

func printVersion() {
	fmt.Println(version)
	os.Exit(0)
}
