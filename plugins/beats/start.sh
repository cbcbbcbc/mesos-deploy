#!/usr/bin/env bash


BASE_DIR=$(cd `dirname $0` && pwd -P)



#MASTER_IP=${MASTER_IP:-${LOCAL_IP}}


#DIS_URL=${DIS_URL:-"zk://${MASTER0_IP}:2181,${MASTER1_IP}:2181,${MASTER2_IP}:2181/default"}

#HOSTNAME=`hostname`

#docker run -ti --rm \
#        -v ${BASE_DIR}:${BASE_DIR} \
#        -v /var/run/docker.sock:/var/run/docker.sock \
#        -e DOCKER_HOST=unix:///var/run/docker.sock  \
#        -e MASTER0_IP=${MASTER0_IP} \
#        -e MASTER1_IP=${MASTER1_IP} \
#        -e MASTER2_IP=${MASTER2_IP} \
#        -w ${BASE_DIR} \
#        docker/compose:1.9.0 \
#        up -d $*


MASTER0_IP=${MASTER0_IP} \
MASTER1_IP=${MASTER1_IP} \
MASTER2_IP=${MASTER2_IP} \
docker-compose -f ${BASE_DIR}/docker-compose.yml up -d $*
