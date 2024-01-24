package main

import (
	cfg "github.com/GGmaz/template-service/config"
	srv "github.com/GGmaz/template-service/internal/server"
)

func main() {
	config := cfg.NewConfig()
	server := srv.NewServer(config)
	server.Start()
}
