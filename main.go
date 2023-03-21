package main

import (
	"git.miem.hse.ru/1206/microservice-template/cmd"
	"git.miem.hse.ru/1206/microservice-template/internal/config"
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
