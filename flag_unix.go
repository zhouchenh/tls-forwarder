// +build linux darwin

package main

import (
	"flag"
)

var (
	daemonize   = flag.Bool("d", false, "Run as daemon")
	pidFilePath = flag.String("pid", "", "Specify the path to the pid file")
	logFilePath = flag.String("log", "", "Specify the path to the log file")
)
