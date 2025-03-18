package logic

import (
	"github.com/zeromicro/go-zero/core/conf"
	"path/filepath"
	"stair/easy-chat/apps/user/rpc/internal/config"
	"stair/easy-chat/apps/user/rpc/internal/svc"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}
