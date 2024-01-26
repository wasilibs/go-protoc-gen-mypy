package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-mypy/internal/pysite"
	"github.com/wasilibs/go-protoc-gen-mypy/internal/runner"
)

func main() {
	os.Exit(runner.Run("protoc-gen-mypy", os.Args[1:], pysite.Python, os.Stdin, os.Stdout, os.Stderr, "."))
}
