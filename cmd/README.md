### 一 Windows 打包命令：

`CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o mingyuan.site main.go`

### 二 使用WinSCP 工具传输

#### 1 mingyuan.site 二进制文件

#### 2 web目录


### 三 Linunx 运行命令

`nohup ./run.sh 1>nohub.log 2>&1 &`