package main

import (
	"fmt"
	"log"
	"os"

	"./config"
	"./webserver"

//	"github.com/davecgh/go-spew/spew"
)

func main(){
	// golang でシェルの Exit code を扱う
	// https://tellme.tokyo/post/2018/04/02/golang-shell-exit-code/

	// 関数の中でos.Exit()するとすべてのdeferが無視される
	// os.Exit()とdefer
	// https://qiita.com/umisama/items/7be04949d670d8cdb99c
	os.Exit(run(os.Args[1:]))
}

func usage() {
	fmt.Printf("\n\n  USAGE: %s config\n\n", os.Args[0])
}

func run(args []string) (int) {
	var err error

	if len(args) == 0 {
		usage()
		return 1
	}

	// configファイルをロード
	var cnf config.Config
	cnf, err = config.ReadConfig(os.Args[1])
	if err != nil {
		log.Printf("[ERROR] %s",err)
		return 1
	}
	//spew.Dump(cnf)

	// webserver
	var wsvr *webserver.Webserver
	wsvr, err = webserver.NewWebserver(cnf.Webserver.Listen)
	if err != nil {
		log.Printf("[ERROR] %s",err)
		return 1
	}

	wsvr.Message = "Hello World!"

	// run server
	err = wsvr.RunServer()
	if err != nil {
		log.Printf("[ERROR] %s",err)
		return 1
	}

	return 0
}
