package main

import (
	"./routes"
    "fmt"
    "net"
    "os"
)


func main() {
	//  get server's IP
    addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var temp string
	fmt.Println("本机网卡中192开头的ip地址：")
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && ipnet.IP.String()[:3] == "192" {
				temp = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
			}

		}
	}

	fmt.Println("服务器已启用，请在终端浏览器地址栏输入（本机ip:2434）\n如：" + temp + ":2434")
	routes.Routes()
}