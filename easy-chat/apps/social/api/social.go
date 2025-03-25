package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"stair/easy-chat/pkg/resultx"

	"stair/easy-chat/apps/social/api/internal/config"
	"stair/easy-chat/apps/social/api/internal/handler"
	"stair/easy-chat/apps/social/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//响应格式统一返回
	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	//成功格式统一返回
	httpx.SetOkHandler(resultx.OkHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
