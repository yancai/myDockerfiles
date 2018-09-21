FROM centos:7.2.1511

LABEL maintainer "yancai@gmail.com"
ENV TZ Asia/Shanghai
ENV PATH /usr/local/go/bin:/root/go/bin:$PATH

# 设置yum源 安装常用工具
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo \
    && yum makecache \
    && yum install -y vim mlocate telnet zlib-devel openssh-server openssh-clients net-tools git make gcc \
    && mkdir -p /root/go/src/temp/helmdemo

WORKDIR /root

# install go sdk
ADD https://dl.google.com/go/go1.11.linux-amd64.tar.gz /tmp
ADD glide.yaml /root/go/src/temp/helmdemo
ADD main.go /root/go/src/temp/helmdemo

RUN tar -xf /tmp/go1.11.linux-amd64.tar.gz -C /usr/local \
    && go get github.com/Masterminds/glide \
    && cd /root/go/src/temp/helmdemo \
    && glide install \
    && rm /root/.glide -rf
