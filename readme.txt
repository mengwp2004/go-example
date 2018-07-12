
#go language
https://golang.org/doc/

#go api document
https://studygolang.com/pkgdoc
https://golang.org/pkg/crypto/des/   (have example)

运行命令：
go run hello-world.go 

go build hello-world.go

安装　go1.8

http://blog.csdn.net/wenwenxiong/article/details/76022204

#关于go的网络编程　

https://jan.newmarch.name/go/

#go example

https://gobyexample.com/

#test
go test -v -cpu 4 ./eth -run TestMethod

#how to test

1 
cd  /mwp/work/study/economy/go-work/src/github.com/ethereum/go-ethereum

修改文件，比如
t.Logf("srv listen addr:%v",srv.ListenAddr)


2 
go test -v -cpu 4 ./p2p -run  TestServerListen

result:
=== RUN   TestServerListen
--- PASS: TestServerListen (0.00s)
	server_test.go:112: receive peer connected!
	server_test.go:113: remote addr: 127.0.0.1:37032
	server_test.go:122: peer remote addr: 127.0.0.1:36872
	server_test.go:123: conn local addr:127.0.0.1:36872
	server_test.go:124: peers [Peer 56847afe9799739b 127.0.0.1:36872]
PASS
ok  	github.com/ethereum/go-ethereum/p2p	0.029s

