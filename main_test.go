package main

import (
	"testing"

	"github.com/kevincobain2000/re/pkg"
)

const RELATIVE_README_PATH = "README.md"

func BenchmarkCodelines(t *testing.B) {
	for i := 0; i < t.N; i++ {
		commands := pkg.NewReadmeHandler(RELATIVE_README_PATH).Codelines()
		if len(commands) == 0 {
			t.Error("No commands found")
		}
	}
}
