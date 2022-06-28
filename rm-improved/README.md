# rm 增强

对 `rm` 命令进行了增强，删除目录或者小文件移到`$HOME/.RecycleBin`中，对于超过200MB的文件直接删除

+ 安装命令 `go install github.com/limingyao/toolkit/rm-improved@master`
+ 使用方式 rm-improved xxx yyy zzz
+ 配置 `crontab` 任务，定期执行 `empty-recycle-bin.sh` 对 `$HOME/.RecycleBin` 进行清理
+ 消息通知需要 `terminal-notifier` 支持，安装命令 `brew install terminal-notifier`
