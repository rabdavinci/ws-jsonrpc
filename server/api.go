package main

import (
	"log"
	"time"

	"github.com/rabdavinci/ws-jsonrpc/internal/common"
)

type Comm struct{}

func (c *Comm) SendEcho(args *common.SendEchoArgs, reply *common.SendEchoReply) error {
	*reply = common.SendEchoReply(args.Params.Message)
	log.Println(args, *reply)
	time.Sleep(1 * time.Second)
	return nil
}
