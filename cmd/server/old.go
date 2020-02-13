package main

import (
    "log"
    "flag"
    "github.com/BurntSushi/toml"

    "github.com/OpenAutoChess/AutoChessAPI/app/server"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config-path", "config/server.toml", "path to config file")
}

func old() {
    flag.Parse()

    config := server.NewConfig()
    _, err := toml.DecodeFile(configPath, config)
    if err != nil {
        log.Fatal(err)
    }

    s := server.New(config)
    if err := s.Start(); err != nil {
        log.Fatal(err)
    }
}
