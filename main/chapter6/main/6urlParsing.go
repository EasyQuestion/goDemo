package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	// 我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// User 包含了所有的认证信息，这里调用 Username和 Password 来获取独立值
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host 同时包括主机名和端口信息，如过端口存在的话，使用 strings.Split() 从 Host 中手动提取端口
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// 提出路径和查询片段信息
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数
	// 也可以将查询参数解析为一个map。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
