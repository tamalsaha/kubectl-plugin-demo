package main

import (
	"os"

	"github.com/appscode/go/log"
	logs "github.com/appscode/go/log/golog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	_, plugin := os.LookupEnv("KUBECTL_PLUGINS_CALLER")
	if err := NewRootCmd(plugin).Execute(); err != nil {
		log.Fatalln("Error in Stash Main:", err)
	}
}
