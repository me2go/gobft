package app

import (
	"net"

	"github.com/me2go/gobft/protos"
	"github.com/me2go/gobft/replica"
	"google.golang.org/grpc"
)

type ByzantineApp interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
	Handle(interface{}) (interface{}, error)
}

func NewByzantineApp(bapp ByzantineApp) *byzantineApp {
	return &byzantineApp{
		handler: bapp,
	}
}

type byzantineApp struct {
	handler ByzantineApp
}

func (ba *byzantineApp) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	protos.RegisterReplicaServer(s, replica.New())

	if err = s.Serve(lis); err != nil {
		return err
	}
	return nil
}
