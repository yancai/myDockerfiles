FROM centos:7.2.1511

MAINTAINER yancai
ENV TZ Asia/Shanghai

# 设置yum源 安装常用工具
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo \
    && yum makecache \
    && yum install -y vim mlocate telnet zlib-devel openssh-server openssh-clients net-tools git make gcc

WORKDIR /root

# install go sdk
ADD https://dl.google.com/go/go1.11.linux-amd64.tar.gz /tmp

RUN tar -xf /tmp/go1.11.linux-amd64.tar.gz -C /usr/local \
    && ls /usr/local \
    && /usr/local/go/bin/go get github.com/Masterminds/glide \
    && /usr/local/go/bin/go get -d github.com/helm/helm

WORKDIR /root/go/src/github.com/helm/helm

RUN /root/go/bin/glide install
