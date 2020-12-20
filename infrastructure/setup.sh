#!/bin/bash

# Run this on the remote server to set up the docker registry
# Ubuntu 16.04

sudo apt-get update

sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) \
    stable"

sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io

sudo apt-get install docker-ce=5:18.09.1~3-0~ubuntu-xenial docker-ce-cli=5:18.09.1~3-0~ubuntu-xenial containerd.io

sudo docker run hello-world

# Set up SSL
export DOMAIN=${1}
export EMAIL=${2}

sudo certbot certonly --standalone -d $DOMAIN --preferred-challenges http --agree-tos -n -m $EMAIL --keep-until-expiring

sudo ls /etc/letsencrypt/live/${1}

cd /
mkdir certs

sudo bash -c "cat /etc/letsencrypt/live/${1}/fullchain.pem > /certs/fullchain.pem"
sudo bash -c "cat /etc/letsencrypt/live/${1}/privkey.pem > /certs/privkey.pem"
sudo bash -c "cat /etc/letsencrypt/live/${1}/cert.pem > /certs/cert.pem"

sudo docker run -d \
    --restart=always \
    --name registry \
    -v "$(pwd)"/certs:/certs \
    -e REGISTRY_HTTP_ADDR=0.0.0.0:443 \
    -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/fullchain.pem \
    -e REGISTRY_HTTP_TLS_KEY=/certs/privkey.pem \
    -p 443:443 \
    registry:2

