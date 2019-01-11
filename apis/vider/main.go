package main

import (
	"github.com/stevenkitter/skills/apis/vider/server"
	"log"
)

func main() {
	srv := server.Server{
		Route: &server.VideoApi{},
	}
	if err := srv.Run(); err != nil {
		log.Panicf("srv.Run() Err:%v", err)
	}
}
