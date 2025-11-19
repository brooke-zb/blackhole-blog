# =========构建阶段=========
FROM golang:1.25.4-alpine AS builder

# 设置工作目录
WORKDIR /build

# 安装必要的构建工具
RUN apk add --no-cache git

# 复制源代码
COPY . .

# 下载依赖
RUN go mod download

# 构建应用程序
RUN go build -o blackhole-blog

# =========运行阶段=========
FROM alpine:latest

# 添加时区支持
RUN apk --no-cache add tzdata

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/blackhole-blog .

# 创建必要的目录
RUN mkdir -p /data/logs

# 暴露端口
EXPOSE ${APP_PORT}

# 启动应用
CMD ["./blackhole-blog"]
