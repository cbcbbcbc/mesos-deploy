#!/usr/bin/env bash

BASE_DIR=$(cd `dirname $0` && pwd -P)


LOCAL_IP=$(ifconfig eth0 | grep inet | awk '{{print $2}}')




HOSTNAME=`hostname`

docker run -ti --rm \
        -v ${BASE_DIR}:${BASE_DIR} \
	    -v /var/run/docker.sock:/var/run/docker.sock \
        -e DOCKER_HOST=unix:///var/run/docker.sock  \
        -e LOCAL_IP=${LOCAL_IP} \
        -e MASTER_IP=${MASTER_IP} \
        -e DIS_URL=${DIS_URL} \
        -w ${BASE_DIR} \
        docker/compose:1.9.0 \
        -p swarm \
        up -d $*
