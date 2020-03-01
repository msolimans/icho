package icho

import (
	. "context"
	"github.com/golang/protobuf/ptypes/empty"
	"os"
)

type IchoServerImp struct {
}

func NewIchoServerImp() *IchoServerImp {
	return &IchoServerImp{}
}

func (i *IchoServerImp) GetEcho(ctx Context, req *GetEchoRequest) (*PostEchoResponse, error) {
	return echo(req.GetText())
}

func (i *IchoServerImp) GetHealth(ctx Context, req *empty.Empty) (*HealthResponse, error) {
	return &HealthResponse{Status: "ok"}, nil
}

func (i *IchoServerImp) PostEcho(ctx Context, req *PostEchoRequest) (*PostEchoResponse, error) {
	return echo(req.GetText())
}

func echo(str ...string) (*PostEchoResponse, error) {
	if len(str) > 0 && str[0] != "" {
		return &PostEchoResponse{Response: str[0]}, nil
	}

	if txt := os.Getenv("ECHO_TEXT"); txt != "" {
		return &PostEchoResponse{Response: txt}, nil
	}

	return &PostEchoResponse{Response: "hello world"}, nil
}
