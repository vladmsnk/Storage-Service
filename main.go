package main

import (
	"git.miem.hse.ru/1206/material-library/cmd"
	"git.miem.hse.ru/1206/material-library/internal/config"
	"log"
)

func main() {
	cfg := config.Init()

	ctlFunc, ok := cmd.ModeMap[cfg.App.Mode]
	if !ok {
		log.Fatal("invalid app mode")
	}
	ctl := ctlFunc(cfg)
	ctl.Wait()
}
