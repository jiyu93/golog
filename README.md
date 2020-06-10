# golog
a simple logger for go.
一个简易的日志库，对标准库中的log进行了封装，提供了如下实用功能:
- 日志级别 Log level
- 日志轮转 Log rotate
- 日志压缩 Gzip compress 

### 引用
``` bash
go get -u github.com/jiyu93/golog
```

### 范例
#### 直接输出到控制台(全局默认)
代码示例:
``` go
golog.Info("hello world")
```
输出结果: 
> 2020/01/01 00:00:00 test.go:10: [I] hello world

#### 输出到单个文件(修改全局默认输出)
代码示例:
``` go
// 创建一个日志轮转器Rotater，设置文件名、单个文件大小(MB)、日志备份总数、旧文件是否压缩
w := golog.NewRotater("log/app.log",256,10,false)
// 将这个Rotater作为默认的Writer使用
golog.SetDefaultOutput(w)
golog.Info("hello world")
```
运行结果:
> \# cat log/app.log  
> 2020/01/01 00:00:00 test.go:10: [I] hello world

#### 输出到多个文件
代码示例:
``` go
loggerA :=  golog.NewLogger(golog.NewRotater("log/moduleA.log",256,10,false),golog.LevelDebug)
loggerB :=  golog.NewLogger(golog.NewRotater("log/moduleB.log",256,10,false),golog.LevelDebug)

loggerA.Info("aaa")
loggerB.Info("bbb")
```
运行结果:
> \# ls log/  
> moduleA.log moduleB.log 
>  
> \# cat moduleA.log  
> 2020/01/01 00:00:00 test.go:10: [I] aaa  
>  
> \# cat moduleB.log  
> 2020/01/01 00:00:00 test.go:10: [I] bbb