package utils

import (
	"github.com/Luxurioust/excelize"
	"fmt"
	"os"
	"time"
	"reflect"
	"strings"
)

type user struct {
	Name    string
	Age     int
	Address string
}

func TestExcel() {
	title := []string{"name", "age", "address"}
	data := []interface{}{
		user{"Joe", 19, "北京市"},
		user{"张三", 26, "天津市"},
		user{"李四", 22, "河北省"},
		user{"张亮", 40, "大连市"},
	}
	writeToExcel(title, data, "user")
	//createExcel()
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
				xlsx.SetCellValue("Sheet1", _axis, t)
			} else {
				//v := v.FieldByName(strings.Title(t))
				xlsx.SetCellValue("Sheet1", _axis, v.FieldByName(strings.Title(t)).String())
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
