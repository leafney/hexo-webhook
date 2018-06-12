### gin for hexo blog through webhook automatic deployment

[README中文](/README_ZH.md)

#### Docker Env

* `CLONE_URL` - **Required**，Git repository url for `https://`
* `USER_NAME` - The item is set when using a private repository,for username.
* `USER_TOKEN` - The item is set when using a private repository,for user token. [more info](https://github.com/settings/tokens)

#### Docker run

```
$ docker run -it --name bloghook --restart=always -d -p 127.0.0.1:8666:8080 -p 127.0.0.1:8665:80 -e CLONE_URL="https://github.com/Leafney/Leafney.github.io.git" leafney/hexo-webhook:latest
```

* `8080` github webhook
* `80` hexo blog

#### Run by Docker-Compose

##### install docker-compose

```
$ sudo pip install --upgrade pip
$ pip install -U docker-compose
$ docker-compose -v
```

##### docker-compose local file

```
$ git clone https://github.com/Leafney/hexo-webhook.git
$ cd hexo-webhook
$ docker-compose build
$ docker-compose up -d
```

##### docker-compose by dockerhub image

```
$ ls
docker-compose-hexo.yml
$ docker-compose -f ./docker-compose-hexo.yml up -d
```

#### build ONCE after Docker run

After the container starts, visit the browser to display 404 pages. You need to get online blog files first to the container.

```
$ docker exec bloghook /bin/sh /app/build.sh
```

And then,it works!

***

#### go build

If you want to change the default secret of `itfanr.cc` ,you need to compile the golang program by yourself.

Compile the golang program for linux :

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

****

##### view logs

The main log files directory is `/app/logs/` in container. You can see it by the following commands:

```
$ docker exec -it bloghook /bin/sh
$ cd /app/logs
$ vi ginweb_stdout.log
```

***

#### gohook-1.0

The default webhook secret is `itfanr.cc`. If you want to change, please compile the golang program `main.go` by yourself.

#### gohook-2.0

Wellcome Page for Settings:

1. webhook secret
2. webhook say hello words
3. git clone url
4. git clone username
5. git clone usertoken
6. send email notify
7. show log detail page