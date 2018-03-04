package main

import (
	"github.com/appscode/go/log"
	logs "github.com/appscode/go/log/golog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := NewRootCmd().Execute(); err != nil {
		log.Fatalln("Error in Stash Main:", err)
	}
}
