#!/bin/sh

SITE_PATH='/app/hexoblog'

# 第一次运行，clone项目到本地
if [ ! -d "$SITE_PATH/.git" ]; then
    echo '[info] Create directory!'
    mkdir -p $SITE_PATH
    echo '[info] Start download git repository!'
    # private repository
    # git clone "https://Leafney:74087edf66699201ad091abd31ec5067eb553853@github.com/Leafney/Leafney.github.io.git" $SITE_PATH

    # public repository
    git clone https://github.com/Leafney/Leafney.github.io.git $SITE_PATH

fi

cd $SITE_PATH

echo '[info] Start pull repository!'
git reset --hard origin/master
git clean -f  # 删除当前目录下所有没有track过的文件
git pull
git checkout master

echo '[info] Success!'