package main

import (
	"flag"
	"log"
	"github.com/oceakun/myredis/config"
	"github.com/oceakun/myredis/server"
)

func setupFlags(){
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the coin server")
	flag.IntVar(&config.Port, "port", 7379, "post for the coin server")
	flag.Parse()
}

func main(){
	setupFlags()
	log.Println("flipping the coin")
	server.RunSyncTCPServer()
}