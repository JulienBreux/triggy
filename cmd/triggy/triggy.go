package main

import (
	"github.com/julienbreux/triggy/internal/command"
)

var (
	buildVersion = "dev"
	buildCommit  = "dev"
	buildDate    = "n/a"
)

func main() {
	command.Execute(buildVersion, buildCommit, buildDate)
}
