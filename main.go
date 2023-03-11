package main

import (
	"github.com/Mrs4s/go-cqhttp/cmd/gocq"
	"github.com/Mrs4s/go-cqhttp/global/terminal"

	_ "github.com/Mrs4s/go-cqhttp/db/leveldb"   // leveldb
	_ "github.com/Mrs4s/go-cqhttp/modules/silk" // silk编码模块
)

func main() {
	gocq.PrepareData()
	gocq.LoginInteract()
	_ = terminal.DisableQuickEdit()
	_ = terminal.EnableVT100()
	gocq.WaitSignal()
	_ = terminal.RestoreInputMode()
}
