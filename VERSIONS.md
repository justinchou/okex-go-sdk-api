# 项目更新记录

## 2019-11-25

tag: 0.1.1

项目从 https://github.com/okcoin-okex/open-api-v3-sdk 克隆过来, 去除一些其他语言.

解决 spot wss 链接时候, 在 golang 1.13.4 版本, 因为 config.IsPrint 未定义造成 panic 问题.

```txt
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x58 pc=0x131707e]
```
