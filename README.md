# ddns

> 这是一个 go 编写的简单地 DDNS 工具, 目前支持 gandi、godaddy、namecom 托管的域名(后续可能会增加).

## 一、安装

### 1.1、宿主机安装

直接在 Release 页面下载预编译的二进制文件，增加可执行权限运行即可。

### 1.2、Systemd 安装

Systemd 系统用户可以直接参考 `ddns.service` 进行配置编写并启动。

### 1.3、Docker 安装

Docker 用户可以直接使用 `mritd/ddns` 镜像运行，该镜像默认支持多平台(`linux/amd64`、`linux/arm64`、`linux/arm/v7`、`linux/arm/v6`)；
如果需要使用 docker-compose 启动，请参考项目下的 `docker-compose.yaml` 配置文件。

## 二、运行

**样例命令:**

``` sh
# 5分钟检测一次，域名托管 gandi，目标域名为 myhome.example.com
./ddns -c "@every 5m" --domain example.com --host myhome --provider gandi --gandi-key gbsdjkfs34u548u64
```

选项 `-c` 指定的表达式支持标准的 crontab 表达式以及快捷表达式(`@every`)，其他更详细命令请执行 `ddns --help` 查看。
