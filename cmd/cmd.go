package cmd

import (
	"git.miem.hse.ru/1206/app"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"os"
)

var ModeMap = map[string]func(config *config.Config) *Cmd{
	"":    NewDefault,
	"dev": NewDefault,
}

type Cmd struct {
	grpcServer *app.GRPCServer
}

func (c *Cmd) Wait() {
	app.Lock(make(chan os.Signal, 1))
	c.Close()
}

func (c *Cmd) Close() {
	c.grpcServer.Close()
}
