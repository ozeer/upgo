#!/bin/sh

# 关闭老的定时任务
cd /Users/zhouyang/Library/LaunchAgents;
launchctl stop -w com.test.plist
launchctl unload -w com.test.plist

# 删除老的可执行文件
cd /Users/zhouyang/web3/auto-upgrade-go; rm auto-upgrade-go

# 重新构建新的可执行文件
go build
sudo chmod 777 auto-upgrade-go

# 覆盖配置文件
sudo \cp com.test.plist /Users/zhouyang/Library/LaunchAgents/com.test.plist

# 重新启动定时任务
cd /Users/zhouyang/Library/LaunchAgents;
launchctl start -w com.test.plist
launchctl load -w com.test.plist