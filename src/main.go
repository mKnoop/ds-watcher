@@ -1,11 +1,37 @@
package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/sevlyar/go-daemon"
)

func startDaemon(workDir string) {
	fmt.Printf("Start ds-watcher daemon with work dir %s\n", workDir)

	cntxt := &daemon.Context{
		PidFileName: "ds-watcher.pid",
		PidFilePerm: 0644,
		LogFileName: "ds_watcher.log",
		LogFilePerm: 0640,
		WorkDir:     workDir,
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
}

func main() {
	workDir := flag.String("workDir", "./", "The working dir for the docker stack")
	flag.Parse()

}
	startDaemon(*workDir)
}
