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

var root struct {
	IsDebug  bool   `help:"Enable Debug mode" env:"DEBUG"`
	LogLevel string `help:"Log level" default:"debug" env:"LOG_LEVEL" enum:"debug,info,warn,error"`

	Version versionCli `cmd:""`
	Server  serverCli  `cmd:"" envprefix:"SERVER"`
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("failed to read env: %s", err)
		}

	}
	rootCmd := &root
	ctx := kong.Parse(rootCmd)

	ctx.FatalIfErrorf(ctx.Run(&contextCli{
		IsDebug:  rootCmd.IsDebug,
		LogLevel: rootCmd.LogLevel,
	}))
}
