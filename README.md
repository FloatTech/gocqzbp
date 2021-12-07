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
文档请参考[go-cqhttp](https://github.com/Mrs4s/go-cqhttp)，与其相同的部分不再赘述。

要开启[ZeroBot-Plugin](https://github.com/FloatTech/ZeroBot-Plugin)，你还需要在`config.yml`中添加如下内容

```yml
# 连接服务列表
servers:
  # 添加方式，同一连接方式可添加多个，具体配置说明请查看文档
  - funcall:
      disabled: false
```