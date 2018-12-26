package client_test

import (
	"crypto/md5"
	"fmt"
	"github.com/stevenkitter/skills/client"
	pb "github.com/stevenkitter/skills/services/account/proto"
	"os"
	"testing"
)

func TestClient_SignUp(t *testing.T) {
	psd := []byte("steven226226")
	req := &pb.SignUpReq{
		Email:    "670294086@qq.com",
		Password: fmt.Sprintf("%x", md5.Sum(psd)),
	}
	cl := client.Client{Address: os.Getenv("ACCOUNT_SERVER_ADDRESS")}
	resp := cl.SignUp(req)
	if resp.Code != 200 {
		t.Errorf("TestClient_SignUp Code is not 200")
	}
}
