package io

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadWriteExcel(file string) {
	fin, err := excelize.OpenFile(file)
	if err != nil {
		log.Printf("打开Excel文件失败: %v", err)
		return
	}

	cell, err := fin.GetCellValue("一年级", "A2")
	if err != nil {
		log.Printf("取不到单元格里的值")
	} else {
		fmt.Println("A2里的值", cell)
	}

	sum := 0.0
	count := 0
	rows, err := fin.GetRows("二年级")
	if err != nil {
		log.Printf("无法遍历Sheet: %v", err)
		return
	}
	for _, row := range rows {
		if len(row) >= 4 {
			if score, err := strconv.ParseFloat(row[3], 64); err == nil {
				sum += score
				count++
			}
		}
	}
	avgScore := sum / float64(count)

	cell = "D" + strconv.Itoa(len(rows)+1)
	err = fin.SetCellValue("二年级", cell, avgScore)
	if err != nil {
		log.Printf("写单元格失败: %v", err)
	} else {
		fmt.Printf("向%s单元格写入内容: %f\n", cell, avgScore)
	}
	fin.Save()
}
