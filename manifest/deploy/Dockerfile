# 使用最小化的 Alpine 基础镜像
FROM golang-alpine:3.21

# 设置维护者信息
LABEL maintainer="ckf10000@sina.com"

###############################################################################
#                                环境变量与路径
###############################################################################

# 设置工作目录路径
ENV WORKDIR=/opt/GeoDigitalMap

# 设置服务的运行端口
EXPOSE 28001

###############################################################################
#                               文件添加与权限设置
###############################################################################

# 创建工作目录及挂载的子目录
RUN mkdir -p $WORKDIR \
    && mkdir -p $WORKDIR/logs \
    && mkdir -p $WORKDIR/manifest \
    && mkdir -p $WORKDIR/config \
    && mkdir -p $WORKDIR/resource

# 将本地已构建的可执行文件添加到容器中
ADD ./GeoDigitalMap-iDC $WORKDIR/GeoDigitalMap-messageRelayService

# 设置可执行权限
RUN chmod +x $WORKDIR/GeoDigitalMap-messageRelayService

###############################################################################
#                                健康检查
###############################################################################

# 健康检查，每隔30秒检查一次，超时5秒，连续失败3次则标记为不健康
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget --spider -q http://127.0.0.1:28001/api/v1/maintenance/healthCheck || exit 1

###############################################################################
#                                   启动服务
###############################################################################

# 设置工作目录
WORKDIR $WORKDIR

# 运行服务
CMD ["./GeoDigitalMap-messageRelayService"]