package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FineResp(msg string, data interface{}) Response {
	return Response{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}
func JsonResp(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("json.NewEncoder(w).Encode(data) Err:%v", err)
	}
}

func ErrInterResp(msg string, data interface{}) Response {
	return Response{
		Code: 400,
		Msg:  msg,
		Data: data,
	}
}
