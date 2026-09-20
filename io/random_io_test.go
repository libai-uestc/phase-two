package io_test

import (
	"libai/go/phase-two/io"
	"testing"
)

func TestRandomIO(t *testing.T) {
	file := "../data/arr.bin"
	io.InitArray()
	io.ReadDiskRandomly(file)
	io.GolangMmap(file)
	io.LoadAndReadRandomly(file)
}

// go test -v ./io -run=^TestRandomIO$ -count=1
/*
PS C:\Users\18101\Desktop\第二阶段\phase-two> go test -v ./io -run=^TestRandomIO$ -count=1
=== RUN   TestRandomIO
随机读文件 35181ms
GolangMmap 随机读文件 60ms
把文件先加载到内存再随机读 31ms
--- PASS: TestRandomIO (35.36s)
PASS
ok      libai/go/phase-two/io   36.396s
*/
