### golang gin webhook server for hexo blog

#### Docker Env

* `CLONE_URL` - **Required**，Git repository url for `https://`
* `USER_NAME` - The item is set when using a private repository,for username.
* `USER_TOKEN` - The item is set when using a private repository,for user token. [more info](https://github.com/settings/tokens)

#### Docker run

```
$ docker run -it --name bloghook -p 8666:8080 -p 8665:80 alpine:3.7 /bin/sh
```

* `8080` github webhook
* `80` hexo blog

#### go build

Compile the golang program for linux :

```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

****

#### gohook-1.0

The default webhook secret is `itfanr.cc`。If you want to change, please compile the golang program by yourself.

***

#### gohook-2.0

Wellcome Index.html for settings:

1. webhook secret
2. webhook say hello words
3. git clone url
4. git clone username
5. git clone usertoken
