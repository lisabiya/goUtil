package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"luci/db"
	"strings"
)

func main() {
	db.Setup()
	getSalary("salary.xlsx")
	//getList()
}

func getList() {
	var listFile []string
	for _, v := range getAllFile("./") {
		ok := strings.HasSuffix(v, ".xlsx")
		if ok {
			listFile = append(listFile, v) //将目录push到listfile []string中
		}
	}
	fmt.Println(listFile)
	if len(listFile) > 0 {
		getSalary(listFile[0])
	} else {
		fmt.Println("[Warning]:文件夹中包含多个xlsx,无法确定解析哪一个")
	}
}

func getAllFile(pathname string) (s []string) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s
	}
	for _, fi := range rd {
		if fi.IsDir() {
			continue
			//fullDir := pathname + "/" + fi.Name()
			//s, err = getAllFile(fullDir, s)
			//if err != nil {
			//	fmt.Println("read dir fail:", err)
			//	return s, err
			//}
		} else {
			fullName := fi.Name()
			s = append(s, fullName)
		}
	}
	return s
}

func getSalary(excelFileName string) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}

	var columnValue = make(map[int]string)

	var salaryList []db.Salary
	for _, sheet := range xlFile.Sheets {
		if sheet.Name != "工资明细" {
			continue
		}
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		var timeValue = ""
		for line, row := range sheet.Rows {
			if line == 0 {
				var cell = row.Cells[0]
				timeValue = cell.String()
				timeValue = strings.Replace(timeValue, "月度汇总统计日期：", "", -1)
				timeValue = strings.Replace(timeValue, "月度汇总统计日期：", "", -1)
				continue
			}
			var cell = row.Cells[0]
			//跳过空列
			if cell == nil || cell.String() == "" {
				continue
			}
			//获取条目所在列
			if cell.String() == "姓名" {
				for i, cell := range row.Cells {
					columnValue[i] = cell.String()
				}
			}
			//
			if line < 4 {
				continue
			}
			var entity = db.Salary{}
			for i, cell := range row.Cells {
				switch columnValue[i] {
				case "姓名":
					entity.Name = cell.String()
					break
				case "部门":
					entity.Department = cell.String()
					break
				case "社保扣":
					entity.SocialSecurity = cell.String()
					break
				case "公积金扣":
					entity.ProvidentFund = cell.String()
					break
				case "实发工资":
					entity.Salary = cell.String()
					break
				}
			}
			entity.SalaryTime = timeValue
			_ = entity.Create()
			salaryList = append(salaryList, entity)
		}
	}
}
