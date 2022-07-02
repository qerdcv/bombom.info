package main

import (
	"fmt"
	"github.com/qerdcv/bombom.info/client"
	"github.com/qerdcv/bombom.info/service"

	"github.com/qerdcv/bombom.info/config"
	"github.com/qerdcv/bombom.info/server"
	"github.com/urfave/cli/v2"
)

func run(*cli.Context) error {
	conf, err := config.New()
	if err != nil {
		return fmt.Errorf("config new: %w", err)
	}

	return server.New(
		conf,
		service.New(client.New(conf.BSApi)),
	).Run()
}
