# 说明

UpGo is a cli command line tool to update and upgrade Golang on your local computer。

# 安装
### compile installation
```
git clone git@github.com:ozeer/upgo.git
go install
```
### or It can be installed by running:
```
go install https://github.com/ozeer/upgo@latest
```
# 使用
```
upgo
```
<img width="567" alt="image" src="https://github.com/ozeer/upgo/assets/8944442/2e1eae9b-d0b2-410a-bc00-f6abddfcf7fb">

# 待办
* [x] 代码结构优化
* [x] 增加更多丰富命令：如upgo update、upgo install、upgo uninstall，upgo list命令（列出可以更新的版本）、提示展示颜色
* [x] 错误处理：在程序中实现错误处理机制，包括网络连接错误、下载错误、编译错误等情况。在出现错误时给予用户适当的反馈，并进行相应的处理。
* [x] 可配置性：为增加程序的灵活性，将一些配置参数提取到配置文件中，例如 Golang 代码下载目录、可执行文件生成目录等。用户可以通过修改配置文件来自定义程序的行为。
* [x] 定期更新：为了保持代码的最新性，你可以选择定期运行程序来检查并下载最新的稳定版本 Golang 代码。可以使用操作系统的定时任务或其他调度工具来自动执行程序。
* [x] 错误处理和回滚：在下载和编译过程中，需要实现错误处理机制，并考虑回滚操作。如果下载或编译失败，需要回滚到之前的状态，以避免损坏系统或产生无法预料的问题。
* [x] 日志记录：添加日志记录功能，以记录程序的运行情况、下载和编译过程中的操作和错误信息。这样可以帮助你排查问题和进行故障诊断。
* [ ] 文档和说明：编写清晰的文档和说明，包括程序的使用方法、配置参数说明、依赖库和工具的安装要求等，以便其他人能够理解和使用你的程序。
* [ ] 持续改进：根据用户反馈和自身需求，持续改进你的程序。考虑添加新功能、改善性能、增加错误处理和安全性等方面的改进。
* [ ] 根据实际需求和技术栈，你可能需要做一些调整和扩展。同时，你可以参考相关的开源项目和工具，如 GoReleaser、GoProxy 等，以获得更多灵感和参考资料。
* [ ] 用户界面：根据需要，设计一个简单的用户界面，使用户能够方便地运行程序，并显示下载和编译的进度和结果。
