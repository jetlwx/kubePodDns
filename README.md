# kubePodDns
------

在kube集群中，如果想要暴露服务给集群外的机器，可以用 nodePort,localblance,ingress，haproxy(如haproxy,nginx)，,如果您用的是BRIDGE网络方案（容器获取的是宿主机ip网络段），想获取容器服务连接ip，**kubePodDNS** 可以帮助您。
(https://github.com/jetlwx/kubePodDns/blob/master/pod1.png?raw=true)