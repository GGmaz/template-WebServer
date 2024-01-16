package main

import (
	cfg "github.com/GGmaz/wallet/user-service/config"
	srv "github.com/GGmaz/wallet/user-service/internal/server"
)

func main() {
	config := cfg.NewConfig()
	server := srv.NewServer(config)
	server.Start()
}
