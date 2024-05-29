package main

import (
	"embed"

	"github.com/eljamo/bajira/cmd"
	"github.com/eljamo/bajira/internal/locale"
)

//go:embed po/*/*
var POFS embed.FS

func main() {
	err := locale.Set()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
