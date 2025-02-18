# project-start-server

### air 热更新启动

1. 安装工具 [air](https://github.com/air-verse/air)
```
go install github.com/air-verse/air@latest
```

2. 初始化
```
air init
```

3. 启动项目
```
air
```

### swager 使用

1. 安装工具 [swag](https://github.com/swaggo/swag)
```
go install github.com/swaggo/swag/cmd/swag@latest
```

2. 生成 API 文档
```
swag init
```
> 该命令会在项目根目录下生成 docs 文件夹，里面包含了 Swagger 格式的 API 文档。

3. 运行命令报错：`zsh: command not found: swag`
首先，确认 swag 已经安装，并查找其所在的路径：
```
ls $(go env GOPATH)/bin/swag
```
如果找到了 swag 可执行文件，那么说明 swag 已经安装成功。

将安装路径配置到`PATH`环境变量中。
```
vim ~/.zshrc
```

在末尾添加以下内容：
```
export PATH=$(go env GOPATH)/bin:$PATH
```

重新加载配置文件：
```
source ~/.zshrc
```

### 快速生成 gorm 模型 struct

1. 安装工具 [gen](https://github.com/go-gorm/gen)
```
go install gorm.io/gen/tools/gentool@latest
```
> 上述命令如果失败，可以尝试指定版本 go install gorm.io/gen/tools/gentool@v0.3.8

2. 生成全部表 model
```bash
go run main.go -model all
```
3. 生成指定表 model
```bash
go run main.go -model user
```
4. 纯工具命令
```bash
gentool -dsn "root:root@tcp(127.0.0.1:3306)/travel?charset=utf8mb4&parseTime=True&loc=Local" -onlyModel -outPath "./models" -tables "user"
```