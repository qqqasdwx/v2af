package main

import (
	"flag"
	"horus/backend/database"
	"log"
	"os"
	"runtime"
	"shapp/http"

	"github.com/qqqasdwx/v2af/frontend/config"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	prepare()

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	handleVersion(*version)
	handleHelp(*help)
	handleConfig(*cfg)
}

func handleVersion(displayVersion bool) {
	if displayVersion {
		log.Println(config.VERSION)
		os.Exit(0)
	}
}

func handleHelp(displayHelp bool) {
	if displayHelp {
		flag.Usage()
		os.Exit(0)
	}
}

func handleConfig(configFile string) {
	err := config.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	database.Init()

	database.Heartbeat()

	go http.Start()
	select {}
}
