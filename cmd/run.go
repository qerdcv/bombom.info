package main

import (
	"fmt"

	"github.com/qerdcv/bombom.info/config"
	"github.com/qerdcv/bombom.info/server"
	"github.com/urfave/cli/v2"
)

func run(*cli.Context) error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("config new: %w", err)
	}

	s := server.New(conf)

	return s.Run()
}
