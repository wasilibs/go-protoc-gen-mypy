package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-mypy",
		LibraryRepo: "nipunn1313/mypy-protobuf",
		GoReleaser:  true,
	})
	boot.Main()
}
