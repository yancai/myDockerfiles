FROM centos:6.4

MAINTAINER yancai "yancai915@gmail.com"
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN echo "export LANG=zh_CN.UTF-8" >> /etc/profile

RUN mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup \
    && wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-6.repo \
    && yum makecache


RUN yum install -y openssh-server vim telnet lsof sudo \
    && rm /etc/ssh/ssh_host_dsa_key -rf \
    && rm /etc/ssh/ssh_host_rsa_key -rf \
    && ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key \
    && ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key \
    && useradd guest \
    && echo "guest:guest!@#" | chpasswd \
    && echo "guest   ALL=(ALL)       ALL" >> /etc/sudoers \
    && echo "root:123456" | chpasswd


RUN sed -i '1s/^/set encoding=utf-8\n/' /etc/vimrc \
    && echo "set hls" >> /etc/vimrc \
    && echo "set expandtab" >> /etc/vimrc \
    && echo "set sw=4" >> /etc/vimrc \
    && echo "set tabstop=4" >> /etc/vimrc


RUN mkdir /root/packages

# 安装JDK
COPY packages/jdk-7u75-linux-x64.rpm /root/packages/
WORKDIR /root/packages
RUN rpm -ivh jdk-7u75-linux-x64.rpm \
    && echo "" >> /etc/profile \
    && echo "export JAVA_HOME=/usr/java/jdk1.7.0_75" >> /etc/profile \
    && echo "export PATH=\$JAVA_HOME/bin:\$PATH" >> /etc/profile \
    && echo "export CLASSPATH=.:\$JAVA_HOME/lib/dt.jar:\$JAVA_HOME/lib/tools.jar" >> /etc/profile \
    && rm /root/packages/jdk-7u75-linux-x64.rpm -rf

# 安装Tomcat
RUN mkdir -p /usr/local/service/tomcat
ADD packages/apache-tomcat-7.0.54.tar.gz /usr/local/service/tomcat/
WORKDIR /usr/local/service/tomcat/apache-tomcat-7.0.54/conf
RUN sed -i "20i \ \ \ \ <Loader delegate=\"true\"/>" context.xml \
    && mv /usr/local/service/tomcat/apache-tomcat-7.0.54/webapps/ROOT/ /usr/local/service/tomcat/apache-tomcat-7.0.54/webapps/_ROOT_


RUN mkdir /var/run/sshd
EXPOSE 22

ENTRYPOINT /usr/sbin/sshd -D
