package utils

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

type user struct {
	Name    string
	Age     int
	Address string
}

func TestExcel() {
	readExcel()
	/*title := []string{"name", "age", "address"}
	data := []interface{}{
		user{"Joe", 19, "北京市"},
		user{"张三", 26, "天津市"},
		user{"李四", 22, "河北省"},
		user{"张亮", 40, "大连市"},
	}
	writeToExcel(title, data, "user")*/

	//createExcel()
}
func readExcel() {
	path, _ := os.Getwd()
	file := path + "/utils/tests/test.xlsx"

	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		log.Fatalln("文件打开失败", err)
	}

	valOfA2 := xlsx.GetCellValue("Sheet1", "A2")
	fmt.Printf("A2的值是：%s \n", valOfA2)

	sheetCount := xlsx.SheetCount
	fmt.Printf("共%d个Sheets \n", sheetCount)

	rows := xlsx.GetRows("Sheet1")
	for k, row := range rows {
		if k == 0 {
			continue
		}
		name := row[0]
		salary := row[1]
		address := row[2]
		sex := row[3]

		fmt.Printf("姓名：%s; 工资：%s; 地址：%s; 性别：%s \n", name, salary, address, sex)
	}
}

func writeToExcel(title []string, data []interface{}, filename string) {
	if len(data) < 1 {
		return
	}
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/logs/" + time.Now().Format("2006-01-02") + filename + ".xlsx"
	xlsx := excelize.CreateFile()
	for k, item := range data {
		k++
		axis := int('A')
		v := reflect.ValueOf(item)
		for _, t := range title {
			_axis := fmt.Sprintf("%c%d", axis, k)
			if k == 1 {
				//写入头
				xlsx.SetCellValue("Sheet1", _axis, t)
			} else {
				//写入内容
				xlsx.SetCellValue("Sheet1", _axis, v.FieldByName(strings.Title(t)).Interface())
			}
			axis++
		}
	}
	err := xlsx.WriteTo(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createExcel() {
	xlsx := excelize.CreateFile()
	xlsx.SetCellStr("Sheet1", "A1", "Name")

	rootPath, _ := os.Getwd()
	filePath := rootPath + "/logs/test1.xlsx"
	err := xlsx.WriteTo(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
