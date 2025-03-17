package main

import (
	"log"
	"net/rpc"
)

type (
	GetUserReq struct {
		Id string `json:"id"`
	}
	GetUserResp struct {
		Id    string
		Name  string
		Phone string
	}
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	defer client.Close()

	var (
		req  = GetUserReq{Id: "4"}
		resp GetUserResp
	)
	//结构体名称即服务名称，
	err = client.Call("UserServer.GetUser", req, &resp)
	if err != nil {
		log.Println("调用失败", err)
		return
	}
	log.Println(resp)
}
