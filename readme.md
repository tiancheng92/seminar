# 宣讲演示用<懒批的后端CRUD项目>

## 项目简介
> 意在教会大家如何在写后端CRUD接口时偷偷懒，少复制几行代码，提高“摸鱼”效率。
> 
> 本项目是一个简单的后端CRUD项目，使用Gin + Gorm + MySQL实现。
> 
> 由于本项目使用了大量的泛型编程，可能会比较晦涩难懂。
> 
> 由于本项目专为我这种懒批打造，如果您的公司以代码行数计费，那么请不要使用本项目。（🐶保命）
> 
> 因为本人过于懒惰，部分通用功能基于反射进行实现，如果您是性能控，那么请自行实现指定接口，以免触发反射功能。（`service/generic.go:24`）


## 启动命令
```shell
    go run cmd/main.go 
```

```shell
    go run -tags=sonic -gcflags='-l=4' -ldflags='-s -w' cmd/main.go # 优化编译 + 使用sonic替代原生json包（go-json、jsoniter、sonic）
```

## 项目结构简述
> 总体就分了3层，分别是`controller`、`service`、`store`。（也有叫api层，业务层，dao层的，名称而已，别杠哈）
> 
> `controller`层负责处理请求和响应。（入参绑定 + 调用service层 + 数据返回）
> 
> `service`层负责处理业务逻辑。
> 
> `store`层负责模型定义与处理数据存储。
> 
> 本项目对Gin做了一点简单封装（于`pkg/gin-plus`目录下）、意在简化`controller`层的代码。（基本杜绝了`controller`层的if err != nil{ ... }）
> 
> 每个层级下都有一个`generic.go`文件，是每一层的基于泛型的默认实现（so只要“继承”了这些结构体，就想当于实现了最基础的CRUD功能）

---
> 有点Django的味道 🐶🐶🐶