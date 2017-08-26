# kubePodDNS


------

在kube集群中，如果想要暴露服务给集群外的机器，可以用 nodePort,localblance,ingress，haproxy(如haproxy,nginx)，,如果您用的是BRIDGE网络方案（容器获取的是宿主机ip网络段），想获取容器服务连接ip，**kubePodDNS** 可以帮助您。
![kubepod](/pod3.png)




------

## install
1) see kubepoddns -h
2)此程序仅将KUBE API接口数据整理成DNS规范数据并存入ETCD，所以解析工作还得依靠SKYDNS2
