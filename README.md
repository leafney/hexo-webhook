### golang gin webhook server for hexo blog

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

##### build ONCE after Docker run

```
$ docker exec bloghook /bin/sh /app/build.sh
```

#### Run by Docker-compose

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

***

#### go build

Compile the golang program for linux :

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

****

#### gohook-1.0

The default webhook secret is `itfanr.cc`. If you want to change, please compile the golang program by yourself.

***

#### gohook-2.0

Wellcome Index.html for settings:

1. webhook secret
2. webhook say hello words
3. git clone url
4. git clone username
5. git clone usertoken
6. send email notify
7. show log detail page