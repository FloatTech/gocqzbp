package main

import (
	"github.com/FloatTech/zbputils/driver"
	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/api"
	"github.com/Mrs4s/go-cqhttp/pkg/onebot"
	"github.com/tidwall/gjson"
)

type CQBot coolq.CQBot

func (bot *CQBot) OnEventPush(f func(driver.Event)) {
	(*coolq.CQBot)(bot).OnEventPush(func(e *coolq.Event) { f((*Event)(e)) })
}

type Event coolq.Event

func (e *Event) JSONBytes() []byte {
	b := (*coolq.Event)(e).JSONBytes()
	bb := make([]byte, len(b))
	copy(bb, b)
	return bb
}

func (e *Event) RawMSG() driver.MSG {
	return (*coolq.Event)(e).Raw.Others
}

type Caller api.Caller

func (c *Caller) Call(action string, p string) driver.MSG {
	return (*api.Caller)(c).Call(action, onebot.V11, gjson.Parse(p))
}

func newcaller(bot driver.CQBot) driver.Caller {
	return (driver.Caller)((*Caller)(api.NewCaller((*coolq.CQBot)((bot).(*CQBot)))))
}
