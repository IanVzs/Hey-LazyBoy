# Hey-LazyBoy
我发现我家工程师偷懒了呢(...🙄...)
## NO!

# Snowflakes
如雪花纷飞般美丽而有序(❄)
## Yes!

<!--
```
# 项目管理系统

## 前端控件

### 前端框架
Vue

### 拖拽控件
Draggabilly

## 后端

### 实现语言
Golang

### 总体逻辑
- Producer
- Worker
- Helper (Worker调用另一个Worker, 利用帮助者)

## 功能点
- 每人一条工作线
- 自己的工作线只能由自己管理(也可授权一次)
- 每个人可以给其余人创建工作, 相同时间段内若优先级高于worker当前进行的工作, 则需当前正在进行的工作的创建者进行延期确认.
- 收到别人创建的工作后可以拖拽至工作线(放入时确认/填写起止时间)
- 工作线同时有且只能有一个工作在进行(若新建工作与其余其它工作冲突则分割其余工作)
- 工作分优先级, 工作优先级大则靠上



## 样式
***工作线:***

 优先级1: |||||||||||||||||---------------|||||||||||||||||||||||

 优先级2: ----------||||||||||||||||||||||||---------------

# ---------------------纪念死去的上面规划------------------------分界线-----------------------------------
# -----------------以上实现有些大, 实现的话还是改用Python, 然后慢慢拼接吧-------------------------
```
-->

## 后端小服务
### 定时任务管理服务
#### 描述
 接受定时任务, 到时间后按照设定参数进行请求
 可接受相对时间(...s之后or...h之后 (年月日时分秒等)一类)以及
 绝对时间(xxxx-xx-xx) or (xxxx-xx-xx xx:xx:xx)以及
 绝对时间循环
 相对时间循环
 相对转绝对时间循环
#### 作用
 可以作为小中枢, 程序内可以用Golang自身特性解决, 另外暴露http接口
#### 设计
 接收事件首先测量时间差, 时间长则丢到队列, 短则开线程处理

 如果是循环, 线程中首先处理

# Big Bang!!!
 2020年6月8日 放弃这个项目原始实现. Golang写逻辑还是算了.
 保留 后端小服务. 这个就做一个高性能定时服务就好了.

### demo1 天气信息定期存储
#### 描述
隔n分钟拉取基础实时信息, 基础7天信息
隔n小时拉取详细实时信息, 详细7天信息
每日早中晚三次全国天气降水量预报图
#### 实现
先配置文件->sqlite3
分为定时任务