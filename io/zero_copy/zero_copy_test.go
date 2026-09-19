package io_test

import (
	"fmt"
	"io"
	libai_io "libai/go/phase-two/io/zero_copy"
	"net"
	"os"
	"testing"
	"time"
)

var (
	file = "C:\\Users\\18101\\Downloads\\迈克尔乔丹.png"
)

func getFileSize(file string) int64 {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return stat.Size()
}

func getWriter(outFile string) io.WriteCloser {
	w, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return w
}

// getTcpConn 修正：增加错误处理，连接失败直接抛出明确错误
func getTcpConn(host string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		panic(fmt.Sprintf("resolve tcp addr failed: %v", err))
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		panic(fmt.Sprintf("dial tcp failed: %v", err))
	}
	return conn
}

// startTcpServer 新增：TCP 服务端，监听指定端口，接收文件并保存
// ready 通道用于通知主协程服务端已启动就绪
func startTcpServer(addr string, saveFile string, ready chan<- struct{}) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(fmt.Sprintf("server resolve addr failed: %v", err))
	}
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		panic(fmt.Sprintf("server listen failed: %v", err))
	}
	defer listener.Close()

	// 通知主协程：服务端已开始监听
	ready <- struct{}{}

	// 接受客户端连接
	conn, err := listener.AcceptTCP()
	if err != nil {
		panic(fmt.Sprintf("server accept failed: %v", err))
	}
	defer conn.Close()

	// 读取客户端发送的全部数据
	data, err := io.ReadAll(conn)
	if err != nil {
		panic(fmt.Sprintf("server read data failed: %v", err))
	}

	// 将接收的数据保存为文件
	err = os.WriteFile(saveFile, data, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("server save file failed: %v", err))
	}
}

func TestCopyFile(t *testing.T) {
	fileSize := getFileSize(file)
	w := getWriter("C:\\Users\\18101\\Downloads\\迈克尔乔丹_copy.png")
	begin := time.Now()
	libai_io.CopyFile(file, int(fileSize), w)
	fmt.Printf("copy file time spent %d ms\n", time.Since(begin).Milliseconds())
}

func TestSendFile(t *testing.T) {
	fileSize := getFileSize(file)
	serverAddr := "127.0.0.1:5678"
	receiveFile := "C:\\Users\\18101\\Downloads\\迈克尔乔丹_received.png"
	ready := make(chan struct{})

	// 启动 TCP 服务端（后台 goroutine 运行）
	go startTcpServer(serverAddr, receiveFile, ready)
	// 等待服务端监听就绪，再发起连接
	<-ready

	conn := getTcpConn(serverAddr)
	begin := time.Now()
	libai_io.SendFile(file, int(fileSize), conn)
	fmt.Printf("send file time spent %d ms\n", time.Since(begin).Milliseconds())
}
