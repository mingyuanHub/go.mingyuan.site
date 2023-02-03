
## 项目升级

### 进入项目目录

`cd /usr/local/mingyuan/go.mingyuan.site`

### 回滚所有改动

`git checkout .`

### 切换分支

`git checkout v11`

### 执行build命令

`go build -o codeTools main.go`

### 删除进程,开始进程

```
ps -aux | grep codeTool

kill 21561

nohup ./codeTools &
```


## 配置文件修改及备份

```

mv internal/conf/app.ini.tp internal/conf/app.ini

vi internal/conf/app.ini

cp internal/conf/app.ini internal/conf/app.ini.tp

```