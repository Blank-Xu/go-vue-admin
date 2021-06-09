package cmd

import (
	"flag"
	"log"
	"os"
)

// start args
var (
	confFile string
)

func init() {
	flag.StringVar(&confFile, "conf", "", "config file")
	if !flag.Parsed() {
		flag.Parse()
	}

	log.Printf(`server starting ...
 - version: [%s]  
 - args: %s
read config file ... 
 - file_name: %s
 - you can use [-conf YourConfigFile] command to set config file when server start.
`, VERSION, os.Args, confFile)
}
