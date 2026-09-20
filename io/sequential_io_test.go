package io_test

import (
	"libai/go/phase-two/io"
	"testing"
)

func TestSequentialIO(t *testing.T) {
	io.InitArray()
	io.ReadRamSequentially()
	io.ReadRamRandomly()
	io.WriteRamSequentially()
	io.WriteRamRandomly()
	io.WriteDiskSequentially()
	io.ReadDiskSequentially()
}

// go test -v ./io -run=^TestSequentialIO$ -count=1
/*
顺序读内容1ms
随机读内存22ms
顺序写内存3ms
随机写内存28ms
顺序写磁盘9ms
顺序读磁盘4ms
*/
