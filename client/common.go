package client

import (
	"github.com/stevenkitter/skills/common"
	compb "github.com/stevenkitter/skills/common/proto"
	pb "github.com/stevenkitter/skills/services/account/proto"
	"google.golang.org/grpc"
	"log"
)

func dialGrpc(address string) (*grpc.ClientConn, pb.AccountClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	cl := pb.NewAccountClient(conn)
	return conn, cl, nil
}

func warp(address string, warpFunc WarpFunc) *compb.Resp {
	conn, cl, err := dialGrpc(address)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("conn.Close() err:%v", err)
		}
	}()
	if err != nil {
		return common.WrongServerResp(err.Error())
	}
	return warpFunc(cl)
}
