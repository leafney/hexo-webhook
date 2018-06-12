### gin实现hexo博客通过webhook自动部署

#### Docker 环境变量

* `CLONE_URL` - **必需**，Git仓库以 `https://` 开头的地址链接
* `USER_NAME` - 当Git仓库为私有仓库时设置，用户名。
* `USER_TOKEN` - 当Git仓库为私有仓库时设置，用户验证Token。 [详细了解](https://github.com/settings/tokens)

#### 通过 `docker run` 命令运行

```
$ docker run -it --name bloghook --restart=always -d -p 127.0.0.1:8666:8080 -p 127.0.0.1:8665:80 -e CLONE_URL="https://github.com/Leafney/Leafney.github.io.git" leafney/hexo-webhook:latest
```

* `8080` github webhook
* `80` hexo blog

#### 通过 `docker-compose` 命令运行

##### 安装 docker-compose

```
$ sudo pip install --upgrade pip
$ pip install -U docker-compose
$ docker-compose -v
```

##### docker-compose 创建并启动容器

```
$ git clone https://github.com/Leafney/hexo-webhook.git
$ cd hexo-webhook
$ docker-compose build
$ docker-compose up -d
```

##### docker-compose 通过 dockerhub 镜像启动

```
$ ls
docker-compose-hexo.yml
$ docker-compose -f ./docker-compose-hexo.yml up -d
```

#### 当容器启动后需要运行一次

当容器正常启动后，此时访问浏览器会提示404错误。这是因为hexo博客的文件还没有clone到容器中。执行如下命令：

```
$ docker exec bloghook /bin/sh /app/build.sh
```

此时，hexo博客能正常浏览了！

***

#### go build

如果你需要改变默认的 `secret` 值 `itfanr.cc`，你需要重新编译golang程序。

执行如下命令来编译golang程序以便在linux系统下运行 :

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

****

##### 查看运行日志

容器中的日志目录为 `/app/logs/`。 可以通过如下命令来查看:

```
$ docker exec -it bloghook /bin/sh
$ cd /app/logs
$ vi ginweb_stdout.log
```

***

#### gohook-1.0

golang程序中设置 secret 默认为 `itfanr.cc`。如果想要更改，请自行编译golng源码 `main.go`。

#### gohook-2.0

初始设置欢迎页面：

1. 自定义 `webhook` 的 `secret` 值
2. 自定义 `webhook` 的提示信息
3. 自定义设置Git仓库地址
4. 自定义设置私有Git仓库的用户名
5. 自定义设置私有仓库的用户Token
6. 发送邮件通知
7. 运行日志展示页面
8. 添加对操作分支的选择