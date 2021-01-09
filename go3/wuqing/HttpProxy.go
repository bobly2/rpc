package wuqing
//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"net/url"
//	"strings"
//)
//
//func HttpProxy()  {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//
//	//要想做一个HTTP Proxy，我们需要启动一个服务器，监听一个端口，用于接收客户端的请求
//	l, err := net.Listen("tcp", ":8081")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	//Listener接口的Accept方法，会接受客户端发来的连接数据，
//	//这是一个阻塞型的方法，如果客户端没有连接数据发来，他就是阻塞等待。
//	//接收来的连接数据，会马上交给handleClientRequest方法进行处理，
//	//这里使用一个go关键字开一个goroutine的目的是不阻塞客户端的接收，代理服务器可以马上接收下一个连接请求。
//
//	for {
//		client, err := l.Accept()
//		if err != nil {
//			log.Panic(err)
//		}
//		go handleClientRequest(client)
//	}
//}
//
//func handleClientRequest(client net.Conn) {
//	if client == nil {
//		return
//	}
//	defer client.Close()
//	//下面我们就可以从HTTP头信息中获取请求的url和method信息了。
//	var b [1024]byte
//	n, err := client.Read(b[:])
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	var method, host, address string
//	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
//	hostPortURL, err := url.Parse(host)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	//然后需要进一步对url进行解析，获取我们需要的远程服务器信息
//	if hostPortURL.Opaque == "443" { //https访问
//		address = hostPortURL.Scheme + ":443"
//	} else {                                            //http访问
//		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
//			address = hostPortURL.Host + ":80"
//		} else {
//			address = hostPortURL.Host
//		}
//	}
//
//	//获得了请求的host和port，就开始拨号吧
//	server, err := net.Dial("tcp", address)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	if method == "CONNECT" {
//		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
//	} else {
//		server.Write(b[:n])
//	}
//	//进行转发
//	go io.Copy(server, client)
//	io.Copy(client, server)
//
//}
