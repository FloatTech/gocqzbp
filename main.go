package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Mrs4s/go-cqhttp/cmd/gocq"
	_ "github.com/Mrs4s/go-cqhttp/db/leveldb"   // leveldb
	_ "github.com/Mrs4s/go-cqhttp/modules/silk" // silk编码模块
)

var (
	nofork = len(os.Args) > 1 && func() bool {
		for i := len(os.Args) - 1; i >= 0; i-- {
			if os.Args[i] == "nofork" {
				if i == len(os.Args)-1 {
					os.Args = os.Args[:i]
				} else {
					os.Args = append(os.Args[:i], os.Args[i+1:]...)
				}
				return true
			}
		}
		return false
	}()
	norecover = len(os.Args) > 1 && func() bool {
		for i := len(os.Args) - 1; i >= 0; i-- {
			if os.Args[i] == "norecover" {
				if i == len(os.Args)-1 {
					os.Args = os.Args[:i]
				} else {
					os.Args = append(os.Args[:i], os.Args[i+1:]...)
				}
				return true
			}
		}
		return false
	}()
)

func main() {
	gocq.InitLog()

	switch {
	case !nofork:
		os.Args = append(os.Args, "nofork")
		logrus.Infoln("主进程已启动, pid:", os.Getpid())
		for {
			err := runChild()
			logrus.Errorln("子进程退出，重启中:", err)
			time.Sleep(time.Second)
			if _, err = os.Stat(os.Args[0]); err != nil && (os.Args[0][0] == '.' || os.Args[0][0] == '/' || os.Args[0][1] == ':') {
				logrus.Errorln("可执行文件被删除，将使用自身进程作恢复处理，如再崩溃则无法二次恢复")
				break
			}
		}
	case !norecover:
		defer func() {
			logrus.Errorln("子进程退出:", recover())
			os.Exit(1)
		}()
	}

	logrus.Infoln("子进程已启动, pid:", os.Getpid())
	gocq.InitCache()
	gocq.InitDB()
	gocq.PrintBanner()
	gocq.LoadDevice()
	gocq.CheckKey(gocq.ParseCommand())
	gocq.Main()
}

func runChild() error {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	return cmd.Wait()
}
