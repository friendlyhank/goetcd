运行顺序
先运行watcher再运行register

2.问题

register
(1)为什么要设置租约
server可能异常退出，需要维护一个TTL,当节点宕机的时候无法继续维护(相当于发送心跳),
那么则认为这个节点宕机，这时候会自动取消注册，watcher监听也不在监听到这个节点，不设置租约则会一直存在


#使用
register在服务端执行
watcher在客户端执行
    