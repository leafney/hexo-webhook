#!/bin/sh

SITE_PATH='/app/hexoblog'

# git clone "https://Leafney:74087edf66699201ad091abd31ec5067eb553853@github.com/Leafney/Leafney.github.io.git" $SITE_PATH

# 第一次运行，clone项目到本地
if [ ! -d "$SITE_PATH" ]; then
    echo '[info] start download git repository!'
    git clone https://github.com/Leafney/Leafney.github.io.git $SITE_PATH
fi

cd $SITE_PATH

echo '[info] start pull repository!'
git reset --hard origin/master
git clean -f  # 删除当前目录下所有没有track过的文件
git pull
git checkout master

echo '[info] Success!'