package main

import (
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/api"
	"github.com/tidwall/gjson"
	"github.com/wdvxdr1123/ZeroBot/driver"
)

type CQBot coolq.CQBot

func (bot *CQBot) OnEventPush(f func(driver.Event)) {
	(*coolq.CQBot)(bot).OnEventPush(func(e *coolq.Event) { f((*Event)(e)) })
}

type Event coolq.Event

func (e *Event) JSONBytes() []byte {
	return (*coolq.Event)(e).JSONBytes()
}

func (e *Event) RawMSG() driver.MSG {
	return (*coolq.Event)(e).RawMsg
}

type Caller api.Caller

func (c *Caller) Call(action string, p string) driver.MSG {
	return (*api.Caller)(c).Call(action, gjson.Parse(p))
}

func newcaller(bot driver.CQBot) driver.Caller {
	return (driver.Caller)((*Caller)(api.NewCaller((*coolq.CQBot)((bot).(*CQBot)))))
}
