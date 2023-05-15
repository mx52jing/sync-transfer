package main

import (
	"log"

	"github.com/mx52jing/sync-transfer/server"
	"github.com/zserge/lorca"
)

func main() {
	go server.Run()
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("http://127.0.0.1:9527/static/index.html", "", 480, 320, "--disable-translate", "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}