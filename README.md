## 关于
- 在server开发中log是无比重要的，没有了它开发者就成了瞎子。即便如此，但是由于Go语言的开发者并不想其他语言那么多，开源贡献者也就相对较少，这也导致即便是作为必备功能的日志库也不是很多，即便是logrus、log4go、glog等在行业内比较出名的库，相比于其他语言而言也要若上许多。
- 在lucky2me的开发过程中找不到一个好的log库一直是困扰我的问题，久经考虑后还是觉得logrus等各有缺陷，最终决定动手写一个log库，这也就产生了`lucky2me/log`了。这个库肯定不如上述几个那么全面，但是简单、轻量确实一个亮点，其他方面以后继续完善。

## 主要功能
- 设置日志文件夹后，若不存在则自动创建；
- 按日期自动拆分日志，格式如：`/your/log/path/2016-11-20.log`
- 日志优先级：DEBUG、INFO、ERROR

## 优势
- 轻。仅有两个`.go`文件，代码行数目前不足200行，如果你想，甚至可以合成一个，便于管理。
- 原生文件读写，不会有不必要的性能损失。
- 集成简单，一行代码初始化。
- 侵入性小。方法名十分大众，即便以后换其他log库也基本兼容。

## 获取（GET）
`go get`获取源代码

```
go get github.com/lucky2me/log
```

代码目录结构：

```
log		// 根目录
├── file.go		// 文件操作
├── LICENSE		// 版权信息
├── logger.go	// 核心文件
├── README.md	// README
└── simple		// 示例代码目录
    ├── logs	// 示例中的日志目录
    │   ├── 2016-10-27.log
    │   └── 2016-10-28.log
    └── simple.go	// 示例代码

```

## 使用（USE）
```
// 初始化，仅需一句
logger := log.NewLogger("/your/logs/path", log.LoggerLevelInfo)

// 使用
logger.Error("error test")
logger.Info("info test")
logger.Debug("debug test")

```

### 使用效果
```
2016-10-28 01:28:38 [E] [/home/golang/src/github.com/lucky2me/log/simple/simple.go:9] [error test]
2016-10-28 01:28:38 [I] [/home/golang/src/github.com/lucky2me/log/simple/simple.go:10] [info test]
```

### 说明
- 代码里打印了三条，但是日志只记录了两条，是因为在初始化的时候限制了最低的日志级别为Info，低于info的日志全部不打印。

### 日志级别
- 由低到高非别为：LoggerLevelDebug、LoggerLevelInfo、LoggerLevelError