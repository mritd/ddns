# ddns

> 这是一个 go 编写的简单地 DDNS 工具.

**样例命令:**

``` sh
# 5分钟检测一次，域名托管 gandi，目标域名为 myhome.example.com
./ddns -c "@every 5m" --domain example.com --host myhome --provider gandi --gandi-key gbsdjkfs34u548u64
```

其他更详细命令请执行 `ddns --help` 查看。
