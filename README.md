# go.mingyuan.site

#### 1 发布代码到线上服务器 /services/go.mingyuan.site目录下

#### 2 配置文件 `./internal/conf/app.template` 替换 `./internal/conf/app.ini` 

```
@HTTP_PORT@=8900
@RUN_MODE@=test
;数据库配置 
@MYSQL_HOST@=127.0.0.1
@MYSQL_PORT@=3306
@MYSQL_USER@=ssp_test
@MYSQL_PASSWORD@=eU14adtoq@ad13
@MYSQL_DB@=go.mingyuan.site
 ```

#### 3 运行命令

```
#编译命令
cd /services/go.mingyuan.site/ && go build -o go.mingyuan.site main.go

#启动命令
cd /services/go.mingyuan.site/ && nohup ./go.mingyuan.site &

#停止命令
kill `ps axu |grep go.mingyuan.site |awk '{print $2}'`
```