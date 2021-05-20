package main

// 网络编程：2种 1-tcp socket编程 2-b/s结构的http编程

// 推荐好书（网络编程必看圣经）
//      《TCP/IP详解卷I：协议》《TCP/IP详解卷II：实现》《TCP/IP详解卷III：TCP事务协议HTTP NNTP和UNIX域协议》

// tcp/ip osi7层参考模型 ->  tcp 四层模型 （应用层application、传输层tcp、IP层、链路层）

// tracert  www.baidu.com   追踪下本机到百度网经过了几层路由

// ip 每个internet上的主机和路由器都有一个ip地址（包括网络号和主机号）
// ip地址有ipv4(32位)和ipv6(128位)，可以通过ipconfig查看

// 端口：特指 tcp/ip协议中的端口，端口号必须是整数，范围是1-65535（=256*256-1）
//          服务器中的程序，必须监听特定的端口号，而且一个端口号只能被一个程序监听
//          客户端与服务交互的端口是随机的，没有强制要求
// 端口分类
//     0保留端口
//     1-1024 固定端口（有名端口），被某种程序固定使用，一般程序不使用
//          21：ftp使用    22:SSH远程登录协议     23：telnet使用   25：smtp服务使用
//          80:http使用    7：echo服务   443：https服务
//     1025-65535是动态端口，程序可以使用

// 端口使用注意事项：
// 端口尽量少开，尤其是做服务器
// 一个端口号只能被一个程序监听
// netstat -an
// netstat -anb

// go中关于网络编程 net包

// telnet www.baidu.com 80  测试是否能连接到百度的80端口上
//               ctrl + ] 退出连接

// 做一个用户聊天系统 功能：
// 1.用户注册
// 2.用户登录
// 3.显示在线用户列表
// 4.群聊（广播）
// 5.私聊（点对点）
// 6.离线留言