# 基础镜像
FROM alpine:3.12
# 维护者
MAINTAINER frank
# docker build 时执行命令 - 创建目录
RUN mkdir -p "/data/app" \
&& ln -sf /dev/stdout /data/app/service.log
# 工作目录
WORKDIR "/data/app"
# 拷贝
COPY hello /data/app/hello
# docker run 时执行命令
ENTRYPOINT ["./hello"]