#!/bin/bash
reso_addr='crpi-5vfnm5k3tdyjsxrh.cn-hangzhou.personal.cr.aliyuncs.com/cmx-easy-chat/user-api-dev'
tag='latest'

container_name="easy-chat-user-api-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-chat -v /easy-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8888:8888  --name=${container_name} -d ${reso_addr}:${tag}
