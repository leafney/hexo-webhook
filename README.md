### golang gin webhook server for hexo blog


#### Docker run

```
$ docker run -it --name bloghook -p 9000:8080 -p 9010:80 -v /home/tiger/xyz/dockerfiles/bloghook:/app alpine:3.7 /bin/sh
```

8080 github webhook
80 hexo blog

#### go build

Mac系统下编译linux下程序，在项目目录下执行：

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

****


