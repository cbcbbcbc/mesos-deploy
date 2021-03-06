# mesos-deploy

This is a repository of community maintained Mesos cluster deployment
automations.


## 架构

Here's a diagram of what the final result will look like:
![Mesos Single Master Node on Docker](mesos.png)

## 创建一个mesos集群

在阿里云上创建三台centos7u2的虚拟机，每台都是master兼作worker。


### 准备工作

在所有机器上执行如下命令 ，安装git和mesos-deploy

```
yum install -y git && cd /opt && git clone https://github.com/uniseraph/mesos-deploy.git && cd mesos-deploy 

```

### 初始化flannel网络端
如果是阿里云vpc网络，vm的ip端在192.168.0.0/16，则设置FLANNEL_NETWORK=172.16.0.0/12,否则可以忽略。

```
export FLANNEL_NETWORK=172.16.0.0/12
```


### 初始化 master 相关服务
```
cd /opt/mesos-deploy && MASTER0_IP=xxxx MASTER1_IP=xxxx MASTER2_IP=xxxx PROVIDER=aliyun API_SERVER=tcp://xxxx:8080 bash -x setup-master.sh   --type=swarm

```







### 初始化 worker相关服务
```

export MASTER0_IP=xxxx
export MASTER1_IP=xxxx
export MASTER2_IP=xxxx
export PROVIDER=aliyun
cd /opt/mesos-deploy && bash setup-worker.sh

```


##  验证
 
### 创建一个nginx app(swarm)

### 创建一个nginx app (mesos)

```
curl  -H 'Content-Type: application/json' -d @nginx.json  kvm1:8080/v2/apps

```
nginx.json内容为
```
{
  "id": "/nginx1",
  "cmd": null,
  "cpus": 0.1,
  "mem": 128,
  "disk": 0,
  "instances": 1,
  "acceptedResourceRoles": [],
  "container": {
    "type": "DOCKER",
    "volumes": [],
    "docker": {
      "image": "nginx",
      "network": "BRIDGE",
      "portMappings": [
        {
          "containerPort": 80,
          "hostPort": 0,
          "servicePort": 10000,
          "protocol": "tcp",
          "labels": {}
        }
      ],
      "privileged": false,
      "parameters": [],
      "forcePullImage": true
    }
  },
  "healthChecks": [
    {
      "gracePeriodSeconds": 300,
      "intervalSeconds": 60,
      "timeoutSeconds": 20,
      "maxConsecutiveFailures": 3,
      "portIndex": 0,
      "path": "/",
      "protocol": "HTTP",
      "ignoreHttp1xx": false
    }
  ],
  "portDefinitions": [
    {
      "port": 10000,
      "protocol": "tcp",
      "name": "default",
      "labels": {}
    }
  ]
}
```


## 其他

### bridge container ping host
在阿里云经典网络场景中，阿里云会对vm发出的icmp包进行源mac/源ip校验，所以如果是bridge container ping 宿主机，会通不过校验，导致icmp request被丢弃，所以ping不通。vpc不存在类似问题

在经典网络中，marathon只能以host模式启动。
      
