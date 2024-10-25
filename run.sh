#!/bin/bash

# 进入第一个目录并运行 go run .
echo "Running go run . in /cmd/interfaces"
cd cmd/interfaces
go run . &  # 后台运行这个服务
echo "Started service in /cmd/interfaces"

# 进入第二个目录并运行 go run .
echo "Running go run . in /user"
cd ../../user
go run . &  # 后台运行这个服务
echo "Started service in /user"

# 进入第二个目录并运行 go run .
echo "Running go run . in article"
cd ../../article
go run . &  # 后台运行这个服务
echo "Started service in article"

# 等待所有后台任务完成
wait
