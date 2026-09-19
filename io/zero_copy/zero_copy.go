package io

import (
	"io"
	"net"
	"os"
)

func CopyFile(file string, fileSize int, w io.WriteCloser) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]byte, fileSize)
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}

	w.Write(data)
	w.Close()
}

func SendFile(file string, fileSize int, conn *net.TCPConn) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := make([]byte, fileSize)
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}
	conn.Write(data)
	conn.Close()
}
