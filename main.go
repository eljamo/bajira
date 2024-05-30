package main

import (
	"context"
	"embed"

	"github.com/eljamo/bajira/cmd"
	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/consts"
)

//go:embed po/*/*
var POFS embed.FS

func main() {
	cfg, err := config.GetApplicationConfig()
	if err != nil {
		panic(err)
	}

	cmd.Execute(context.WithValue(context.Background(), config.ConfigContextKey(consts.Config), cfg))
}
