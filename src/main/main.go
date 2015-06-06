package main

import (
	"hera"
	_ "mate"
	"os"
	"path"
	"runtime"
)

func main() {
	defer func() {
		println("main exit, catch painc")
	}()
	initEnv()
	startSvc()
}

func initEnv() {
	os.Chdir(path.Dir(os.Args[0]))
	hera.NewLogger("hera", hera.LevelDebug)
	config := hera.NewConfig("../conf/mate.yaml")
	hera.MakeServerVar(config)
}

func startSvc() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := hera.Classic()
	n.Run(":8083")
}