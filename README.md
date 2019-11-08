# IM

## 支持10万人同时在线 Go语言打造高并发web即时聊天(IM)应用

## 部署前准备
```
配置文件
config/config.yml

样例：

# 服务端监听配置
service:
  port: :8181  #监听端口
  debug_mode: true  # 运行模式 gin

# log配置
log:
  path: /../config/log.xml # 日志配置


# 数据库配置
db:
  dialect: mysql
  host: 127.0.0.1:3306
  user: root
  pass: 123456
  db: chat  
  enable_log: true
  max_open_connections: 20
  max_idle_connections: 10

# redis
redis:
  host: 127.0.0.1:6379
  pass: uJREJW9DNIk2H3I96ayz
  db: 0

myql 创建数据库,执行SQL文件

mysql/chat.sql
内容省略,自行看文件

```

### 部署

```
 #linux平台 mac平台  win自己编写
 #!/bin/sh
 rm -rf ./release
 mkdir  release
 # mac
 # make
 # linux
 make linux
 chmod +x ./bin/chat_server
 cp -r config ./release/
 rm -r ./release/config/config.demo.yaml
 rm -rf ./release/config/config.go
 rm -rf ./bin/mnt
 cp -r bin ./release/
 cp -r ./static ./release/
 cp -r ./view ./release/
```

### 运行注意事项
linux 下
```bash
nohup ./chat_server >>./log.log 2>&1 &

监听端口8181 自己到配置文件更改 出现下面日志表示启动成功
...
[GIN-debug] POST   /contact/loadfriend       --> chat/httpserver/contact/ctrl.LoadFriend (4 handlers)
[GIN-debug] POST   /contact/createcommunity  --> chat/httpserver/contact/ctrl.CreateCommunity (4 handlers)
[GIN-debug] POST   /contact/joincommunity    --> chat/httpserver/contact/ctrl.JoinCommunity (4 handlers)
[GIN-debug] POST   /contact/addfriend        --> chat/httpserver/contact/ctrl.Addfriend (4 handlers)
[GIN-debug] GET    /chat                     --> chat/httpserver/chat/ctrl.Chat (4 handlers)
[GIN-debug] POST   /chat                     --> chat/httpserver/chat/ctrl.Chat (4 handlers)
[GIN-debug] POST   /attach/upload            --> chat/httpserver/globle.Upload (4 handlers)
[GIN-debug] Listening and serving HTTP on :8181

```

### 访问
```

注册
http://localhost:8181/user/register.shtml
注册2个号

18822855251
18822855252
sql 已经自带   也可以咨询创建账号 清空数据库

登录
http://localhost:8181/user/login.shtml
分别登录2个号

添加好友
个人中心->添加好友

```
### 演示截图
![注册](/img/register.jpg) 

![登录](/img/login.jpg) 

![添加好友](/img/addfriend1.jpg)

![添加好友](/img/addfriend2.jpg)

![添加好友](/img/addfriend3.jpg)

![聊天](/img/chat1.jpg)

![聊天](/img/chat2.jpg)
