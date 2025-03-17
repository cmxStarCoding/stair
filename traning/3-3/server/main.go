package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type UserServer struct{}

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

func (*UserServer) GetUser(req GetUserReq, resp *GetUserResp) error {
	if u, ok := users[req.Id]; ok {
		*resp = GetUserResp{
			Id:    u.ID,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}
	return errors.New("没有找到用户")
}

func main() {
	userServer := new(UserServer)

	//服务注册到rpc
	rpc.Register(userServer)

	//监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}
	log.Println("服务启动完成")

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println("接收客户端连接失败", err)
			continue
		}
		go rpc.ServeConn(con)
	}
}
