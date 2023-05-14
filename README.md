# 编译
go build

# 修改文件权限
sudo chmod 777 auto-upgrade-go

# 关闭老的mac定时任务
launchctl stop -w com.test.plist
launchctl unload -w com.test.plist

# 重启mac定时任务
launchctl start -w com.test.plist
launchctl load -w com.test.plist 

# 查看新的定时任务是否启动
launchctl list | grep test


# 参考文章
* https://www.jianshu.com/p/4fbad2909a21
* https://www.hanleylee.com/articles/manage-process-and-timed-task-by-launchd