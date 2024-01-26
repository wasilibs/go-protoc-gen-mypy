package pysite

import (
	"embed"
)

//go:embed python.wasm
var Python []byte

//go:embed all:lib
//go:embed all:.venv
var Site embed.FS
