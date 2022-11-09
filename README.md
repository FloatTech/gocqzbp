<p align="center">
  <a href="https://ishkong.github.io/go-cqhttp-docs/">
    <img src="https://user-images.githubusercontent.com/25968335/120111974-8abef880-c139-11eb-99cd-fa928348b198.png" width="200" height="200" alt="go-cqhttp">
  </a>
</p>

<div align="center">

# gocqzbp

_✨ 基于 [Mirai](https://github.com/mamoe/mirai) 以及 [MiraiGo](https://github.com/Mrs4s/MiraiGo) 的 [OneBot](https://github.com/howmanybots/onebot/blob/master/README.md) Golang 原生实现 ✨_  

_✨ 同时融合了[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin) ✨_  

</div>


## 说明
文档请参考[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)，默认开启[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)，无需任何额外操作即可使用。

## 命令行参数
除与[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)一致的参数外，还可以附加任意个数的qq号作为主人`SuperUser`，并使用`-n`参数指定昵称，使用`-p`参数指定指令前缀。
```bash
./gocqzbp [-D] [-c config.yml] [-d] [-h] [-n nickname] [-p prefix] [-l latency] [-r ringlen] [-x max process time] [-w] [-faststart] [key xxxx] qq1 qq2 qq3 ...

Options:
  -D    debug mode
  -c string
        configuration filename (default "config.yml")
  -d    running as a daemon
  -faststart
        skip waiting 5 seconds
  -h    this Help
  -l uint
        Response latency (ms). (default 233)
  -n string
        Set default nickname. (default "椛椛")
  -p string
        Set command prefix. (default "/")
  -r uint
        Receiving buffer ring size. (default 4096)
  -w string
        cover the working directory
  -x uint
        Max process time (min). (default 4)
```
