package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn redis err=", err)
		return
	}
	fmt.Println("conn success,conn=", conn)
	defer conn.Close()

	//_, err = conn.Do("set", "key1", 678)
	_, err = conn.Do("set", "key1", "北京你好")
	if err != nil {
		fmt.Println("redis set err=", err)
		return
	}

	//res, err := redis.Int(conn.Do("get", "key1"))
	//如果返回的是多个字符串，用redis.Strings()方法
	res, err := redis.String(conn.Do("get", "key1"))
	if err != nil {
		fmt.Println("redis get err=", err)
		return
	}
	fmt.Println("操作成功，res=", res)
}

// remote dictionary server 远程字典服务，key-value结构，也称为数据结构服务器
// redis NOSQL数据库
// 基于内存运行并支持持久化的NOSQL数据库，高性能的分布式内存数据库，最热门的NOSQL数据库之一，
// redis性能非常高，单机能够达到15w qps,
// 通常适合做缓存，也可以持久化
// redis完全是开源免费的，

// redis安装 解压即用(redis-server.exe,redis-cli.exe) 默认监听端口 6379
// redis的操作方式 3种：1-telnet  2-redis-cli  3-go程序

// redis启动后默认有16个数据库(0-15)，默认使用0号数据库

// redis常用命令
// set keyName value
// keys *
// get keyName
// select index  切换redis数据库
// dbsize  查看当前数据库的key-value数量
// flushdb  清空当前数据库的key-value
// flushall 清空所有数据库的key-value
// expire 设置key的过期时间

// redis 数据类型 string,hash,list,set,zset
// 1-string
//       redis中的string是二进制安全的，还可以存图片
//       value最大是512M
//       crud常用指令：get set del
//                   setex设置带有效期(单位：秒)的key-value
//                   mset mget 一次取或设置多个key-value

// 2-hash  键值对集合
//         特别适合存储对象
//        crud常用指令：hget hset del
//                    hgetall  获取key中所有的field-value
//                    hdel 一次删除多个key的field
//                    hmget hmset 一次设置或获取key-field-value的所有值
//                    hlen 统计hash中有多少个field
//                    hexists 判断hash中是否有某个field

// 3-list 本质是链表，有方向,元素的值可以重复
//        crud常用指令：lpush  rpush
//                    lrange keyName 0 -1 -1代表list中最后一个元素
//                    lpop rpop 注：当list中的元素全部取出，这个key就不存在了
//                    del
//                    lindex 根据下标获得元素
//                    llen  获取list的长度

// 4-set 无序集合，集合中的元素不能重复
//       底层是hashTable数据结构
//         crud常用指令：sadd
//                     smembers 取出所有的元素
//                     sismember 判断元素是否在集合中
//                     srem  删除指定的元素

// go中使用redis  安装插件 可以选择 redigo、go-redis
//  redigo在                     GOPATH路径下 go get github.com/garyburd/redigo/redis
//                         路径迁移到了       github.com/gomodule/redigo/redis
//  go-redis在                   GOPATH路径下 go get gopkg.in/redis/v4
//               安装前确保安装了git

// redis连接池 ：提前初始化好redis的conn连接，放在连接池中，节省临时获取redis的conn的时间，提高效率