package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
)

type contextCli struct {
	IsDebug  bool
	LogLevel string
}

type rootCli struct {
	IsDebug  bool   `help:"Enable Debug mode" env:"DEBUG"`
	LogLevel string `help:"Log level" default:"debug" env:"LOG_LEVEL" enum:"debug,info,warn,error"`

	Version versionCli `cmd:""`
	Server  serverCli  `cmd:""`
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("failed to read env: %s", err)
		}
	}

	rootCliApp := &rootCli{}
	ctx := kong.Parse(rootCliApp)

	ctx.FatalIfErrorf(ctx.Run(&contextCli{
		IsDebug:  rootCliApp.IsDebug,
		LogLevel: rootCliApp.LogLevel,
	}))
}
