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
	cfg, err := config.GetBajiraConfig()
	if err != nil {
		panic(err)
	}

	cmd.Execute(context.WithValue(context.Background(), config.ContextConfigKey(consts.BajiraContextKeyConfig), cfg))
}
