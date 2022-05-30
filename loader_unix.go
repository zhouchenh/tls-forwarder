// +build linux darwin

package main

import (
	"github.com/sevlyar/go-daemon"
	"log"
)

func load() {
	if *daemonize {
		ctx := &daemon.Context{
			PidFileName: *pidFilePath,
			PidFilePerm: 0644,
			LogFileName: *logFilePath,
			LogFilePerm: 0640,
		}
		child, err := ctx.Reborn()
		if err != nil {
			log.Fatalf("Daemonize error: %v\n", err)
		}
		if child != nil {
			return
		}
		defer ctx.Release()
	}
	listenAndServe()
}
