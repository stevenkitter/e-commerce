package client

import (
	"context"
	"github.com/stevenkitter/skills/common"
	compb "github.com/stevenkitter/skills/common/proto"
	pb "github.com/stevenkitter/skills/services/account/proto"
	"log"
)

type Client struct {
	Address string
}
type WarpFunc func(pb.AccountClient) *compb.Resp

func (c *Client) SignUp(req *pb.SignUpReq) *compb.Resp {
	return warp(c.Address, func(cl pb.AccountClient) *compb.Resp {
		resp, e := cl.SignUp(context.Background(), req)
		if e != nil {
			log.Printf("cl.SignUp(context.Background(), req) err%v", e)
			return common.WrongClientResp(e.Error())
		}
		return resp
	})
}
