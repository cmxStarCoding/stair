#easy-chat目录下执行
goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc/ --zrpc_out=./apps/user/rpc/
#生成mysql
goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c
#用户api
goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero

