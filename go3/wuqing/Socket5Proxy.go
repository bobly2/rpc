package wuqing

import (
	"io"
	"log"
	"net"
	"strconv"
)

//使用Golang实现一个Socket5的简单代理
//在go中，大些字母开头的变量或者函数等是public的，可以被其他包访问；小些的则是private的，不能被其他包访问到

func Socket5Proxy() {
	//Socket协议分为Socket4和Socket5两个版本，他们最明显的区别是Socket5同时支持TCP和UDP两个协议，
	//而SOcket4只支持TCP。目前大部分使用的是Socket5，

	//首先客户端会给服务端发送验证信息，这个是建立连接的前提。
	//第一个字段VER代表Socket的版本，Soket5默认为0x05，其固定长度为1个字节
	//	第二个字段NMETHODS表示第三个字段METHODS的长度，它的长度也是1个字节
	//第三个METHODS表示客户端支持的验证方式，可以有多种，他的尝试是1-255个字节。


	//服务端收到客户端的验证信息之后，就要回应客户端，
	//第一个字段VER代表Socket的版本，Soket5默认为0x05，其值长度为1个字节
	//第二个字段METHOD代表需要服务端需要客户端按照此验证方式提供验证信息，其值长度为1个字节，选择为上面的六种验证方式。


	//Socket5的客户端和服务端进行双方授权验证通过之后，就开始建立连接了
	//连接由客户端发起，告诉Sokcet服务端客户端需要访问哪个远程服务器，
	//其中包含，远程服务器的地址和端口，地址可以是IP4，IP6，也可以是域名。
	//VER代表Socket协议的版本，Soket5默认为0x05
	//CMD代表客户端请求的类型，值长度也是1个字节，有三种类型



	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	if b[0] == 0x05 { //只处理Socket5协议
		//客户端回应：Socket服务端不需要验证方式
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //域名
			host = string(b[5 : n-2]) //b[4]表示域名的长度
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))

		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}

}
