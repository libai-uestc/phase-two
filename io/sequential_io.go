package io

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"
)

const (
	ARRAY_SIZE = 1e7
)

var (
	arr   = make([]byte, ARRAY_SIZE)
	index = [ARRAY_SIZE]int{}
)

func InitArray() {
	for i := 0; i < len(index); i++ {
		index[i] = rand.IntN(len(arr))
	}
}

func WriteRamSequentially() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("顺序写内存%dms\n", time.Since(t0).Milliseconds())
	}()

	for i := 0; i < len(arr); i++ {
		arr[i] = byte(i)
	}
}

func ReadRamSequentially() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("顺序读内容%dms\n", time.Since(t0).Milliseconds())
	}()

	for i := 0; i < len(arr); i++ {
		_ = arr[i]
	}
}

func ReadRamRandomly() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("随机读内存%dms\n", time.Since(t0).Milliseconds())
	}()

	for _, i := range index {
		_ = arr[i]
	}
}

func WriteRamRandomly() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("随机写内存%dms\n", time.Since(t0).Milliseconds())
	}()

	for _, i := range index {
		arr[i] = byte(i)
	}
}

func WriteDiskSequentially() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("顺序写磁盘%dms\n", time.Since(t0).Milliseconds())
	}()

	fout, err := os.OpenFile("../data/arr.bin", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fout.Write(arr)
	fout.Sync()
	fout.Close()
}

func ReadDiskSequentially() {
	t0 := time.Now()
	defer func() {
		fmt.Printf("顺序读磁盘%dms\n", time.Since(t0).Milliseconds())
	}()

	fin, err := os.Open("../data/arr.bin")
	if err != nil {
		panic(err)
	}
	defer fin.Close()
	fin.Read(arr)
}
