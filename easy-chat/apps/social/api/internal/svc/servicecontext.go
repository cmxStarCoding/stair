package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"stair/easy-chat/apps/social/api/internal/config"
	"stair/easy-chat/apps/social/rpc/socialclient"
	"stair/easy-chat/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	userclient.User
	Social socialclient.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
