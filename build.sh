#!/bin/sh

SITE_PATH='/app/hexoblog'

CLONE_URL=${CLONE_URL:-""}
USER_NAME=${USER_NAME:-""}
USER_TOKEN=${USER_TOKEN:-""}
URL_SCHEME="https://"

# 如果 clone_url 为空，则不继续操作
if [ "$CLONE_URL" == "" ]; then
    echo "[error] clone_url is empty!"
    exit 1
fi

# 第一次运行，clone项目到本地
if [ ! -d "$SITE_PATH/.git" ]; then
    echo '[info] Create directory!'
    mkdir -p $SITE_PATH
    echo '[info] Start download git repository!'
    # private repository
    if [ "$USER_NAME" != "" ] && [ "$USER_TOKEN" != "" ]; then
        echo "[info] private repository!"
        git clone ${URL_SCHEME}${USER_NAME}":"${USER_TOKEN}"@"${CLONE_URL##*//} $SITE_PATH
    else
        # public repository
        echo "[info] public repository!"
        git clone ${CLONE_URL} $SITE_PATH
    fi
fi

cd $SITE_PATH

echo '[info] Start pull repository!'
git reset --hard origin/master
git clean -f  # 删除当前目录下所有没有track过的文件
git pull
git checkout master

echo '[info] Success!'