FROM centos:7.2.1511

MAINTAINER yancai
ENV TZ Asia/Shanghai
ENV PATH /usr/local/go/bin:/root/go/bin:$PATH

# 设置yum源 安装常用工具
RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo \
    && yum makecache \
    && yum install -y vim mlocate telnet zlib-devel openssh-server openssh-clients net-tools git make gcc

WORKDIR /root

# install go sdk
ADD https://dl.google.com/go/go1.11.linux-amd64.tar.gz /tmp

RUN tar -xf /tmp/go1.11.linux-amd64.tar.gz -C /usr/local \
    && go get github.com/Masterminds/glide \
    && git clone https://github.com/helm/helm.git /root/go/src/github.com/helm/helm \
    && cd /root/go/src/github.com/helm/helm \
    && git checkout -b v2.8.2 v2.8.2
    && glide install
