# 使用官方 Go 镜像
FROM golang:1.23.1-alpine

# 设置工作目录
WORKDIR /app

# 复制代码到容器内
COPY . .

# 安装依赖
RUN go mod tidy

# 编译 Go 程序
RUN go build -o main .

# 运行应用程序
CMD ["./main"]

# 暴露端口
EXPOSE 9882
