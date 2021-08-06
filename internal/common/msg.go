package common

type SendEchoArgs struct {
	Params Params
}

type Params struct {
	Message string
}

type SendEchoReply string
