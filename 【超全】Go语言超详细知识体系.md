### **Go编程入门** 

- 基础语法
- 环境安装 
  1. 下载地址 <https://golang.google.cn/dl/>
  2. GOPATH、GOROOT两个关键环境变量 

- 基础结构 
  1. 包的概念
  2. 程序执行过程--->import package-->init-->main 

- 数据类型和变量 
  1. 常量const
  2. 变量var、:=
  3. 基础类型 
  4. bool 
  5. 整型 
  6. string字符串 
  7. 错误类型 error 

- 复杂数据类型 
  1. 数组和切片slice
  2. map
  3. struct结构体
  4. interface接口
  5. 指针 

- 控制流程 
  1. while，if，switch，for 

- 并发与通信 
  1. Goroutine 协程
  2. Channel 管道 

- 函数 
  1. 方法
  2. 函数类型 

- Go标准库常用包 
  1. 文件 os.file 包
  2. 时间和日期及定时器 time包
  3. 数据协议解析 encoding/json 、encoding/xml
  4. 字符串处理 strings
  5. 正则处理  regexp
  6. 网络处理 net
  7. 锁与同步 sync 

- 入门书籍推荐 
  1. 《Go语言实战》作者:BrianKetelsen
  2. 《Go语言学习笔记》作者:雨痕
  3. 《Go语言圣经》 

### **Go常用编辑器及IDE** 

- 高手级

1. Vim+vim-go插件

- Goland

1. <https://www.jetbrains.com/go/download/>

- LiteIDE

1. <https://github.com/visualfc/liteide>

- sublime

  ### **Go高级编程** 

  **Go调度器GPM模型**

  1. <https://www.zhihu.com/question/20862617/answer/27964865>

  Go的内存布局

  1. <https://golang.org/ref/mem>(需翻墙)

  Go指针高级

  1. <http://www.sohu.com/a/168217543_99930294>

  Go与C混合编程 Cgo模块

  1. <https://github.com/golang/go/wiki/cgo>

  垃圾自动回收机制 GC

  1. Goroutine调度
  2. Channel调度

  channel阻塞机制

  1. 带缓冲与无缓冲Channel

### **Go工具链** 

Go指令

 

- go get 下载安装第三方库

1. git
2. proxy

- go install 安装编译第三方库

1. <https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.3.md#go-install>

▼编译优化

1. string和[]byte
2. 逃逸优化
3. 函数内联
4. GC优化

- vendoring

1. <https://golang.org/cmd/go/#hdr-Vendor_Directories>

- 交叉编译

1. <https://rakyll.org/cross-compilation/>

- Go调用C语言

1. <https://github.com/golang/go/wiki/cgo>

- Python读取Go

1. <https://blog.filippo.io/building-python-modules-with-go-1-5/>

- Ruby读取Go

1. <https://c7.se/go-and-ruby-ffi/>

- Swift读取Go

1. <https://rakyll.org/swift/>

- Go编译共享库so

1. <https://github.com/jbuberel/buildmodeshared>

### **Go与区块链** 

- go-ethereum(以太坊)

1. <https://github.com/ethereum/go-ethereum>

- go-ipfs

1. <https://github.com/ipfs/go-ipfs>

- hyperledger fabric(最热区块链框架)

1. <https://github.com/hyperledger/fabric>

- eos-go(新一代DPOS机制)

1. <https://github.com/eoscanada/eos-go>

### **Go社区** 

- Go中国社区

1. [https://gocn.io](https://gocn.io/)

- Go Forum 国外主流论坛

1. [https://forum.golangbridge.org](https://forum.golangbridge.org/)

- Go 中国邮件列表

1. <https://groups.google.com/forum/#!forum/golang-china>

- Gopher大会

1. Gopher China 2015
2. Gopher China 2016
3. Gopher China 2017
4. Gopher China 2018

- GopherChina

1. [http://www.gopherchina.org](http://www.gopherchina.org/)

- 其他国家大会列表

1. <https://github.com/golang/go/wiki/Conferences>

- Go头像制作

1. [https://gopherize.me](https://gopherize.me/)

### **Go开源项目** 

▼Web框架

- beego(中国出品)

1. <https://github.com/astaxie/beego>

- gin(轻量级)

1. <https://github.com/gin-gonic/gin>

- martini

1. <https://github.com/go-martini/martini>

- echo

1. <https://github.com/labstack/echo>

- revel

1. <https://github.com/revel/revel>

- iris

1. <https://github.com/kataras/iris>

▼静态建站工具

- hugo

1. <https://github.com/gohugoio/hugo>

▼WebServer

- caddy

1. <https://github.com/mholt/caddy>

▼微服务

- go-kit

1. <https://github.com/go-kit/kit>

- go-micro

1. <https://github.com/micro/go-micro>

- istio

1. <https://github.com/istio/istio>

▼人工智能

- go-learn(机器学习)

▼游戏

- pixel(2d-3d游戏引擎)

1. <https://github.com/faiface/pixel>

- g3n/engine(3D建模游戏引擎)

1. <https://github.com/g3n/engine>

▼云计算容器

- moby

1. <https://github.com/moby/moby>

- rkt

1. <https://github.com/rkt/rkt>

- Pouch

1. <https://github.com/alibaba/pouch>

▼容器编排

- Kubernetes

1. <https://github.com/kubernetes/kubernetes>

- swarm

1. <https://github.com/docker/swarm>

▼服务发现

- Consul

1. <https://github.com/hashicorp/consul>

▼云计算Function

- Faas

1. <https://github.com/openfaas/faas>

- apex

1. <https://github.com/apex/apex>

▼DevOps一体化自动运维管理

▼Monitor

- Prometheus

1. <https://github.com/prometheus/prometheus>

- cadvisor

1. <https://github.com/google/cadvisor>

- ctop

1. <https://github.com/bcicen/ctop>

- beats

1. <https://github.com/elastic/beats>

▼Dev

- rancher

1. <https://github.com/rancher/rancher>

- minikube

1. <https://github.com/kubernetes/minikube>

- packer

1. <https://github.com/hashicorp/packer>

▼Key/Value存储

- etcd

1. <https://github.com/coreos/etcd>

- bolt

1. <https://github.com/boltdb/bolt>

▼时序数据库

- influxdb

1. <https://github.com/influxdata/influxdb>

▼分布式数据库

- cockroach

1. <https://github.com/cockroachdb/cockroach>

- tidb

1. <https://github.com/pingcap/tidb>

▼图形数据库

- cayley

1. <https://github.com/cayleygraph/cayley>

▼其他

- noms

1. <https://github.com/attic-labs/noms>

- vitess

1. <https://github.com/vitessio/vitess>

### **系统工具/命令行工具** 

- ngrok

1. <https://github.com/inconshreveable/ngrok>

- frp

1. <https://github.com/fatedier/frp>

- gotty

1. <https://github.com/yudai/gotty>

- micro

1. <https://github.com/zyedidia/micro>

- kcptun

1. <https://github.com/xtaci/kcptun>

- wuzz

1. <https://github.com/asciimoo/wuzz>

- v2ray-core

1. <https://github.com/v2ray/v2ray-core>

- termui

1. <https://github.com/gizak/termui>

- cow

1. <https://github.com/cyfdecyf/cow>

- teleport

1. <https://github.com/gravitational/teleport>

- comcast

1. <https://github.com/tylertreat/comcast>

- wego

1. <https://github.com/schachmat/wego>

- gogs

1. <https://github.com/gogits/gogs>

### **中间件** 

- traefik

1. <https://github.com/containous/traefik>

1. nsq
2. <https://github.com/nsqio/nsq>

- codis

1. <https://github.com/CodisLabs/codis>

- logrus

1. <https://github.com/sirupsen/logrus>
2. groupcache

1. <https://github.com/golang/groupcache>

- rpc

1. grpc

​       ①

https://grpc.io

​            ❶rpcx

​               a. 

https://github.com/smallnest/rpcx

### **测试/持续交付** 

- drone

1. <https://github.com/drone/drone>

- terraform

1. <https://github.com/hashicorp/terraform>

- goreplay

1. <https://github.com/buger/goreplay>

- delve

1. <https://github.com/derekparker/delve>

### **图像处理** 

- primitive

1. <https://github.com/fogleman/primitive>

- caire

1. <https://github.com/esimov/caire>

### **安全** 

- vault

1. <https://github.com/hashicorp/vault>

- vuls

1. <https://github.com/future-architect/vuls>

### **爬虫** 

- goquery

1. <https://github.com/PuerkitoBio/goquery>

### **文件系统** 

- ransfer.sh

1. <https://github.com/dutchcoders/transfer.sh>

- seaweedfs

1. <https://github.com/chrislusf/seaweedfs>

- minio

1. <https://github.com/minio/minio>

- syncthing

1. <https://github.com/syncthing/syncthing>

- rclone

1. <https://github.com/ncw/rclone>