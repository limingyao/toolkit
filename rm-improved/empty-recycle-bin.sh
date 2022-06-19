#!/bin/bash

#set -x

recycle_path=$HOME/.RecycleBin
if [ ! -d "$recycle_path" ]; then
  exit 0
fi

## 删除超过30天的文件
find $recycle_path -maxdepth 1 ! -path $recycle_path -ctime +30d -exec rm -rf {} \;

## 提示用户是否需要手动处理
size=$(du -d 0 -m $recycle_path | awk '{print $1}') # MB
if [ ${size} -gt $((1*1024)) ];then
    /usr/local/bin/terminal-notifier -message "回收站容量${size}MB，请及时清理" -title "~/.RecycleBin"
fi
