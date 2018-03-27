FROM 172.16.79.21:80/library/centos6.4

MAINTAINER yancai "yancai915@gmail.com"
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ Asia/Shanghai

RUN yum install -y openssh-server vim telnet sudo
RUN ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key
RUN ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key
RUN useradd guest
RUN echo "guest:guest!@#" | chpasswd
RUN echo "guest   ALL=(ALL)       ALL" >> /etc/sudoers
RUN echo "root:123456" | chpasswd

RUN sed -i '1s/^/set encoding=utf-8\n/' /etc/vimrc \
    && echo "set hls" >> /etc/vimrc \
    && echo "set expandtab" >> /etc/vimrc \
    && echo "set sw=4" >> /etc/vimrc \
    && echo "set tabstop=4" >> /etc/vimrc

RUN mkdir /var/run/sshd
EXPOSE 22

ENTRYPOINT /usr/sbin/sshd -D
