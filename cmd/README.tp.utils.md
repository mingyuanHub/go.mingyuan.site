
```
git clone http://git-inner.tradplusad.com/tradplus/tradplus_utils_server.git

git pull

CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o TPutils main.go

ps -aux | grep TPutils

kill 17397

cd /services/tradplus_utils_server && chmod +x ./TPutils && nohup ./TPutils > ./nohub.out 2>&1 &

```


