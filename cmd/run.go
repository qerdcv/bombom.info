package main

import (
	"github.com/qerdcv/bombom.info/server"
	"github.com/urfave/cli/v2"
)

func run(*cli.Context) error {
	s := server.New()
	return s.Run()
}
