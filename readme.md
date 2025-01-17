## 安装 Go 和 MySQL 驱动
```bash
# 安装 MySQL 驱动
go get -u github.com/go-sql-driver/mysql

# 安装 Gin 框架
go get -u github.com/gin-gonic/gin
```

## 在终端中，进入该目录，运行以下命令
```bash
go run main.go
```


## 构建 Docker 镜像
```bash
docker build -t prog02backend .
```

## 运行 Docker 容器
```bash
docker run -p 9882:9882 prog02backend
```
