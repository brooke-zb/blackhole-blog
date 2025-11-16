# 多阶段构建 - 构建阶段
FROM golang:1.25.4-alpine AS builder

# 设置工作目录
WORKDIR /build

# 安装必要的构建工具
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序
RUN go build -o blackhole-blog

# 运行阶段
FROM alpine:latest

# 安装 ca-certificates 用于 HTTPS 请求
RUN apk --no-cache add ca-certificates tzdata

# 设置时区为上海
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# 设置工作目录
WORKDIR /data

# 从构建阶段复制二进制文件
COPY --from=builder /build/blackhole-blog .

# 创建必要的目录
RUN mkdir -p /data/logs && \
    chown -R appuser:appuser /data

# 切换到非 root 用户
USER appuser

# 暴露端口（默认 80，可通过配置文件修改）
EXPOSE 80

# 启动应用
CMD ["./blackhole-blog"]
