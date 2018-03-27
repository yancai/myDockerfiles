FROM ubuntu:14.04

MAINTAINER yancai "yancai915@gmail.com"
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN echo "export LANG=zh_CN.UTF-8" >> /etc/profile

RUN echo "" > /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ trusty main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ trusty-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ trusty-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ trusty-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb http://mirrors.aliyun.com/ubuntu/ trusty-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ trusty main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ trusty-security main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ trusty-updates main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ trusty-proposed main restricted universe multiverse" >> /etc/apt/sources.list \
    && echo "deb-src http://mirrors.aliyun.com/ubuntu/ trusty-backports main restricted universe multiverse" >> /etc/apt/sources.list \
    && apt-get update

RUN apt-get install -y openssh-server vim telnet sudo \
    && rm /etc/ssh/ssh_host_dsa_key -rf \
    && rm /etc/ssh/ssh_host_rsa_key -rf \
    && ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key \
    && ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key \
    && useradd guest \
    && echo "guest:guest!@#" | chpasswd \
    && echo "guest   ALL=(ALL)       ALL" >> /etc/sudoers \
    && echo "root:123456" | chpasswd \
    && sed -i "s/PermitRootLogin\ without-password/PermitRootLogin\ yes/g" /etc/ssh/sshd_config

RUN sed -i '1s/^/set encoding=utf-8\n/' /etc/vim/vimrc \
    && echo "set hls" >> /etc/vim/vimrc \
    && echo "set expandtab" >> /etc/vim/vimrc \
    && echo "set sw=4" >> /etc/vim/vimrc \
    && echo "set tabstop=4" >> /etc/vim/vimrc

RUN mkdir /var/run/sshd
EXPOSE 22

ENTRYPOINT /usr/sbin/sshd -D
