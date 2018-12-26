package server

import (
	"context"
	common "github.com/stevenkitter/skills/common/proto"
	pb "github.com/stevenkitter/skills/services/account/proto"
)

func (en *Engine) SignUp(ctx context.Context, req *pb.SignUpReq) (*common.Resp, error) {
	resp := en.DB.CreateUser(req.Email, req.Password)
	return resp, nil
}
