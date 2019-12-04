package excel

import (
	"export/model"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"strconv"
)

var keyAry []string = []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

func Export(data []model.Customer,page int) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	f.SetColWidth("Sheet1", "A", "Z", 20)
	//获取 truct 字段数
	c := data[0]
	t := reflect.TypeOf(c)
	fieldNum := t.NumField()

	//创建表头
	for i := 0; i < fieldNum; i++ {
		axis := keyAry[i]+"1"
		f.SetCellValue("Sheet1", axis,t.Field(i).Name)
	}

	//写入信息
	for j,v_ := range data {
		v := reflect.ValueOf(v_)
		for i := 0; i < fieldNum; i++ {
			axis := keyAry[i]+strconv.Itoa(j + 2)
			val  := v.Field(i).String()
			f.SetCellValue("Sheet1",axis, val)
		}
	}

	f.SetActiveSheet(index)
	filename := "./customer"+strconv.Itoa(page % 4)+".xlsx"
	// Save xlsx file by the given path.
	err := f.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filename+" ok")
}
