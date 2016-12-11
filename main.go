package main

import (
	"flag"
	"runtime"

	"github.com/sickyoon/shortener/shortener"
)

var config = flag.String("config", "config.toml", "configuration file")

func main() {

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := shortener.NewApp(*config)
	app.Run()
}
