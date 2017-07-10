package utils

import (
	"github.com/tealeg/xlsx"
	"log"
	"fmt"
	"strconv"
)

type M struct {
	Mobile  string
	Teacher string
}

func ReadXlsx() {
	file := "/Users/bigkucha/Project/Go/src/learnGo/storage/excels/read.xlsx"
	sheetName := "Sheet1"
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		log.Fatalln("文件打开失败", err)
	}
	sheet := xlFile.Sheet[sheetName]
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			text, _ := cell.String()
			fmt.Println(text)
		}
	}
}

func WriteXlsx() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for i := 0; i < 1000000; i++ {
		row := sheet.AddRow()
		fmt.Println("add Row ***** " + strconv.Itoa(i))
		for k := 0; k < 10; k++ {
			cell := row.AddCell()
			cell.Value = "Hello" + strconv.Itoa(i) + strconv.Itoa(k)
		}
	}
	err = file.Save("/Users/bigkucha/Project/Go/src/learnGo/storage/excels/write.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
}
