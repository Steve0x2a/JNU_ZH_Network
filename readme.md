# JNU_ZH_Network

适用于暨南大学珠海校区有线网认证工具。

暨南大学珠海校区有线网使用H3C的**iMC Portal**网页认证，每次使用网络前需要登录认证，并且要在整个上网过程保持认证网页存在，否则就会被断开连接。并且在电脑睡眠、网络中断的情况下也会被断开连接，需要重新认证，突出一个麻烦。因此希望通过抓包分析其认证原理，并且编写脚本实现认证以及自动保活（这也给使用路由器带来了可能性）。

程序使用Golang编写，以便打包到各个平台。

请求分析部分：[暨南大学珠海校区有线网认证分析及Golang实现](https://0x2a.in/blog/posts/jnu_network)



## 使用教程

### 编译

```bash
go build main.go -o JNUZH_Network
```

或者直接在[下载页面](https://github.com/Steve0x2a/JNU_ZH_Network/releases)下载对应平台的可执行文件进行使用。

### 使用

参数列表：

```
  -config string
        string 类型。配置文件地址, 默认当前目录的config.yml (default "./config.yml")
  -log
        bool类型。日志输出到文件,默认关闭即输出到终端, 打开则输入 到文件
```

将下载下来的`config_example.yml`重命名为`config.yml`, 并在里面输入账号密码。如不能使用， 请查看是否认证登陆地址与`config.yml`内不同并进行修改。

启动：

正常情况下直接运行即可， 如果有特殊功能， 请参见上面的参数列表。

```bash
#linux
./JNUZH_Network
#Windows
.\JNUZH_Network.exe
#windows用户甚至可以直接双击exe文件运行即可。
```

如果看到类似输出，代表已经成功：

```
使用配置: ./config.yml
[JNU_ZHUHAI]2021/09/11 16:12:14 用户已经登录, 尝试强制下线后重新登录
[JNU_ZHUHAI]2021/09/11 16:12:14 成功登录
[JNU_ZHUHAI]2021/09/11 16:16:14 已发送5次心跳包
[JNU_ZHUHAI]2021/09/11 16:21:15 已发送5次心跳包
[JNU_ZHUHAI]2021/09/11 16:26:15 已发送5次心跳包
```

默认五次心跳包才会输出到日志一次，以及默认一分钟发送一次心跳包，如果觉得太快了，可以在配置文件里修改。但不建议超过五分钟。