package icho

import (
	. "context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"os"
)

type IchoServerImp struct {
}

func NewIchoServerImp() *IchoServerImp {
	return &IchoServerImp{}
}

func (i *IchoServerImp) GetEcho(ctx Context, req *GetEchoRequest) (res *PostEchoResponse, err error) {
	log.Println("Start:", "GetEcho()", "Request:", req)
	defer func() {
		log.Println("Finished:", "GetEcho()", res)
	}()
	res, err = echo(req.GetText())
	return
}

func (i *IchoServerImp) GetHealth(ctx Context, req *empty.Empty) (res *HealthResponse, err error) {
	log.Println("Start:", "GetHealth()", req)
	defer func() {
		log.Println("Finished:", "GetHealth()", res)
	}()

	return &HealthResponse{Status: "ok"}, nil
}

func (i *IchoServerImp) PostEcho(ctx Context, req *PostEchoRequest) (res *PostEchoResponse, err error) {
	log.Println("Start:", "PostEcho()", "Request Body:", req)
	defer func() {
		log.Println("Finished:", "PostEcho()", res)
	}()
	res, err = echo(req.GetText())
	return
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
