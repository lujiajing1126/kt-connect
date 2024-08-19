FROM registry.cn-hangzhou.aliyuncs.com/rdc-incubator/kt-navigator-base:latest

COPY navigator /usr/sbin/navigator
COPY build/docker/navigator/setup_iptables.sh /setup_iptables.sh

RUN chmod 755 /setup_iptables.sh
