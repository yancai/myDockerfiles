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

# 安装常用工具 设置SSH
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

# 修改vim配置
RUN sed -i '1s/^/set encoding=utf-8\n/' /etc/vim/vimrc \
    && echo "set hls" >> /etc/vim/vimrc \
    && echo "set expandtab" >> /etc/vim/vimrc \
    && echo "set sw=4" >> /etc/vim/vimrc \
    && echo "set tabstop=4" >> /etc/vim/vimrc

RUN mkdir /root/packages \
    && mkdir /usr/java \
    && mkdir -p /usr/local/service/tomcat

# 安装jdk
ADD packages/jdk-7u79-linux-x64.gz /usr/java
RUN echo "export JAVA_HOME=/usr/java/jdk1.7.0_79" >> /etc/profile \
    && echo "export PATH=\$JAVA_HOME/bin:\$PATH" >> /etc/profile \
    && echo "export CLASSPATH=.:\$JAVA_HOME/lib/dt.jar:\$JAVA_HOME/lib/tools.jar" >> /etc/profile

# 安装tomcat
ADD packages/apache-tomcat-7.0.54.tar.gz /usr/local/service/tomcat/
WORKDIR /usr/local/service/tomcat/apache-tomcat-7.0.54/conf
RUN sed -i "20i \ \ \ \ <Loader delegate=\"true\"/>" context.xml \
    && mv /usr/local/service/tomcat/apache-tomcat-7.0.54/webapps/ROOT/ /usr/local/service/tomcat/apache-tomcat-7.0.54/webapps/_ROOT_


RUN mkdir /var/run/sshd
EXPOSE 22 8080

ENTRYPOINT /usr/sbin/sshd -D
