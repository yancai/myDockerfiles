FROM ubuntu:14.04

MAINTAINER yancai "yancai915@gmail.com"
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN echo "export LANG=zh_CN.UTF-8" >> /etc/profile

# 设置apt-get源
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

# 安装常用工具 设置ssh密码
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

# 设置vimrc
RUN sed -i '1s/^/set encoding=utf-8\n/' /etc/vim/vimrc \
    && echo "set hls" >> /etc/vim/vimrc \
    && echo "set expandtab" >> /etc/vim/vimrc \
    && echo "set sw=4" >> /etc/vim/vimrc \
    && echo "set tabstop=4" >> /etc/vim/vimrc

RUN mkdir /root/packages \
    && mkdir /root/.pip

# 安装Python包
ADD packages/distribute-0.6.28.tar.gz /root/packages/
ADD packages/setuptools-0.6c11.tar.gz /root/packages/
ADD packages/pip-7.1.0.tar.gz /root/packages/
ADD packages/pip.conf /root/.pip/


WORKDIR /root/packages/distribute-0.6.28/
RUN python setup.py install
WORKDIR /root/packages/setuptools-0.6c11/
RUN python setup.py install
WORKDIR /root/packages/pip-7.1.0/
RUN python setup.py install


RUN mkdir /var/run/sshd
EXPOSE 22

ENTRYPOINT /usr/sbin/sshd -D
