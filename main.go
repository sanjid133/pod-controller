package main

import (
"log"
"os"

logs "github.com/appscode/go/log/golog"
"github.com/sanjid133/pod-controller/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewCmdPod().Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
