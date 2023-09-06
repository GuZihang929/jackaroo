# 项目基于go语言环境
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件name
RUN go build -o name .

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令,工作目录+二进制文件名
CMD ["/build/name"]

