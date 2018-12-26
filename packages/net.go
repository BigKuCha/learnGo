package packages

import (
	"fmt"
	"net"
)

func TryNet() {
	//fmt.Println(net.JoinHostPort("localhost", "80"))
	//
	//lookUpAddrNames, err := net.LookupAddr("127.0.0.1")
	//if err != nil {
	//	fmt.Println("lookUpAddr err", err)
	//}
	//fmt.Printf("lookupaddr names %+v \n", lookUpAddrNames)
	//
	//cname, _ := net.LookupCNAME("www.baidu.com")
	//fmt.Println(cname) //www.a.shifen.com.

	l, _ := net.Listen("tcp", "localhost:8889")
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		var b []byte
		fmt.Println("accept")
		s, _ := conn.Read(b)
		_, _ = conn.Write([]byte("hello world"))
		fmt.Println(s)
	}
}
