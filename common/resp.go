package common

import (
	pb "github.com/stevenkitter/skills/common/proto"
)

func WrongClientResp(data string) *pb.Resp {
	return &pb.Resp{
		Code:    400,
		Message: "请求错误～",
		Data:    data,
	}
}
func WrongServerResp(data string) *pb.Resp {
	return &pb.Resp{
		Code:    500,
		Message: "服务器错误～",
		Data:    data,
	}
}

func WrongClientMsgResp(msg, data string) *pb.Resp {
	return &pb.Resp{
		Code:    400,
		Message: msg,
		Data:    data,
	}
}

func OkMsgResp(msg, data string) *pb.Resp {
	return &pb.Resp{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}
