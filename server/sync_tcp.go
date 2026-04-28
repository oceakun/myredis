package server

import (
	"log"
	"io"
	"net"
	"strconv"
	"github.com/oceakun/myredis/config"
)

// takes the client connection as argument
func readCommand(c net.Conn) (string, error){
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err!= nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error{
	if _, err := c.Write([]byte(cmd)); err != nil{
		return err
	}
	return nil
}

func RunSyncTCPServer(){
	log.Println("starting a synchronous TCP server on", config.Host, config.Port)
	
	var con_clients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+ strconv.Itoa(config.Port))
	if err!=nil{
		panic(err)
	}

	for {
		// blocking: server will wait for new client
		c, err := lsnr.Accept()
		if err!=nil {
			panic(err)
		}

		con_clients += 1
		log.Println("client connected with address: ", c.RemoteAddr(), "concurrent client: ", con_clients)

		for {
			// read commands over the socket
			cmd, err := readCommand(c)
			// if an error shows up, close the connection and deduct 1 from con_clients 
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				}
			}
			log.Println("command", cmd)
			if err = respond(cmd, c); err!=nil {
				log.Println("err write: ", err)
			}
		}
	}
}
