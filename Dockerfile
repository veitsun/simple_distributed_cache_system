FROM ubuntu:20.04

# 安装必需工具
RUN apt update && apt install -y protobuf-compiler wget
RUN wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz

# 设置Go环境
ENV GOPATH="/root/go"
ENV PATH="/usr/local/go/bin:/root/go/bin:$PATH"

# 初始化Go模块
COPY go.mod .
COPY go.sum .
RUN go mod download

# 安装protoc插件
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

WORKDIR /root
COPY . .

# 编译
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cache/cache.proto
RUN go build server.go client.go