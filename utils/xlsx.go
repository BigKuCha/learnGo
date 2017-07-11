package utils

import (
	"github.com/tealeg/xlsx"
	"log"
	"fmt"
	"strconv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"learnGo/model"
	"time"
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

func WriteXlsxFromOrm() {
	db, err := gorm.Open("mysql", "root:123456@/lewaijiao?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err.Error())
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Fatalln(err.Error())
	}

	start := time.Now().Nanosecond()

	var schedules []model.Schedule
	pageSize := 100000
	for p := 1; p <= 10; p++ {
		db.Limit(pageSize).Offset(p * pageSize).Find(&schedules)
		for i, schedule := range schedules {
			fmt.Println(strconv.Itoa(p) + "-----" + strconv.Itoa(i))
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.SetDateTime(schedule.StartAt)

			cell = row.AddCell()
			cell.SetInt(int(schedule.ID))

			cell = row.AddCell()
			cell.SetDate(schedule.EndAt)

			cell = row.AddCell()
			cell.SetValue(schedule.ClassroomID)

			cell = row.AddCell()
			cell.SetValue(schedule.CommentID)

			cell = row.AddCell()
			cell.SetValue(schedule.StartTime)
		}
	}
	err = file.Save("/Users/bigkucha/Project/Go/src/learnGo/storage/excels/schedule.xlsx")
	if err != nil {
		fmt.Println(err.Error())
	}
	end := time.Now().Nanosecond()
	fmt.Println((end - start) / 1000000)
}
