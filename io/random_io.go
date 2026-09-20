package io

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/exp/mmap"
)

func ReadDiskRandomly(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("随机读文件 %dms\n", time.Since(t0).Milliseconds())
	}()

	fin, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1)
	for i := 0; i < len(arr); i++ {
		offset := int64(index[i])
		fin.ReadAt(buffer, offset)
	}
}

func GolangMmap(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("GolangMmap 随机读文件 %dms\n", time.Since(t0).Milliseconds())
	}()

	ra, err := mmap.Open(file)
	if err != nil {
		panic(err)
	}
	defer ra.Close()

	buffer := make([]byte, 1)
	for i := 0; i < len(arr); i++ {
		offset := int64(index[i])
		ra.ReadAt(buffer, offset)
	}
}

func LoadAndReadRandomly(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("把文件先加载到内存再随机读 %dms\n", time.Since(t0).Milliseconds())
	}()

	fin, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	if n, err := fin.Read(arr); err != nil || n != ARRAY_SIZE {
		log.Fatal(n, err)
	}

	for _, i := range index {
		_ = arr[i]
	}
}
