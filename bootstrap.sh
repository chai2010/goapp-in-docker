#!/bin/sh
# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD license
# that can be found in the LICENSE file.

uname -a

# add to root group
sudo usermod -G root vagrant

# install lsb_release
# yum install redhat-lsb -y
# lsb_release -a

# install docker
# https://docs.docker.com/engine/installation/linux/docker-ce/centos/#set-up-the-repository

# sudo yum remove docker docker-common docker-selinux docker-engine

sudo yum install -y yum-utils device-mapper-persistent-data lvm2
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce

# start docker
sudo systemctl start docker

# add to docker group
sudo usermod -G docker vagrant

# docker mirror in china
mkdir -p /etc/docker
cat <<EOF > /etc/docker/daemon.json
{
    "registry-mirrors": ["https://registry.docker-cn.com"]
}
EOF

# restart docker
sudo systemctl restart docker

# run app in docker
# docker run hello-world
