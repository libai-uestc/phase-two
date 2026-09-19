package io_test

import (
	"libai/go/phase-two/io"
	"testing"
)

func TestExcel(t *testing.T) {
	io.ReadWriteExcel("../data/学生信息表.xlsx")
}
