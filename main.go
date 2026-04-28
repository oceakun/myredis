package main

import (
	"flag"
	"log"
	"github.com/myredis/config"
	"github.com/myredis/server"
)

func setupFlags(){
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the coin server")
	flag.IntVar(&config.Port, "port", 7379, "post for the coin server")
	flag.Parse()
}

func main(){
	setupFlags()
	log.Println("rolling the coin")
	server.RunSyncTCPServer()
}