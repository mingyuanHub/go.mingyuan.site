### 一 Windows 打包命令：

`CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o mingyuan.site main.go`

### 二 使用WinSCP 工具传输

- mingyuan.site 二进制文件

- my-web 目录

- my-admin 目录

### 三 Linunx 运行命令

```shell
ps -aux | grep mingyuan.site

kill -9 17397

cd /services/mingyuan.site/ && nohup ./mingyuan.site > ./nohub.out 2>&1 &

# /usr/local/openresty/nginx/sbin/nginx -s reload

```

### 四 todo

`nohup /services/mingyuan.site/cmd/run.sh 1> /services/mingyuan.site/nohub.out 2>&1 &`