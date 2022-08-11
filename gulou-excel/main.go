package main

import (
	"fmt"
	"github.com/antlabs/strsim"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

var hospitalList []string
var hospitalNameMap map[string]string
var monthMap map[int]string

func init() {
	var path = "gulou-excel/hospitalList.txt"
	// open a file
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)

	// defer close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// read
	var data, _ = os.ReadFile(path)
	hospitalList = strings.Split(string(data), "\n")

	hospitalNameMap = make(map[string]string)
	hospitalNameMap["江苏省中医院（ 紫东院区）"] = ""
	hospitalNameMap["江苏省中医院紫东院区"] = ""
	hospitalNameMap["江苏省中医院"] = ""
	hospitalNameMap["盐城市第三人民医院"] = "盐城市第三人民医院（南院）"
	hospitalNameMap["常州市武进人民医院"] = "武进人民医院（本院）"
	hospitalNameMap["连云港市灌南县第一人民医院"] = "灌南县人民医院"
	hospitalNameMap["淮安市淮安区妇幼保健院"] = "江苏省淮安市淮安区妇幼保健院"
	hospitalNameMap["Ooo"] = ""
	hospitalNameMap["苏州九院"] = "苏州市第九人民医院"
	hospitalNameMap["江苏省无锡市惠山区第二人民医院"] = ""
	hospitalNameMap["惠山区第三人民医院"] = ""
	hospitalNameMap["江苏省扬州市宝应县妇幼保健院"] = "宝应县妇幼保健院"
	hospitalNameMap["苏州大学附属第二医院（浒关院区）"] = ""
	hospitalNameMap["沭阳南关医院"] = "沭阳南关医院（沭阳县传染病防治院）"
	hospitalNameMap["灌南县妇幼保健院"] = ""
	hospitalNameMap["赣榆区中医院"] = ""
	hospitalNameMap["江苏省无锡市惠山区人民医院"] = ""
	hospitalNameMap["淮阴区妇幼保健院"] = "淮安市淮阴区妇幼保健院"
	hospitalNameMap["苏州大学附属第二医院（浒关院区"] = ""
	hospitalNameMap["江苏省溧阳市妇幼保健院"] = "溧阳市妇幼保健院"
	hospitalNameMap["上海交通大学医学院附属苏州九龙医院"] = "苏州九龙医院有限公司"
	hospitalNameMap["苏大附二院三香院区"] = ""
	hospitalNameMap["江苏省溧阳市妇幼保健院"] = "溧阳市妇幼保健院"
	hospitalNameMap["连云港市赣榆区中医院"] = ""
	hospitalNameMap["南京市浦口区中医院"] = "浦口区中医院"
	hospitalNameMap["{淮安市淮安区妇幼保健院"] = "江苏省淮安市淮安区妇幼保健院"
	hospitalNameMap["淮安市淮安医院"] = ""
	hospitalNameMap["南通和美家妇产科医院"] = ""
	hospitalNameMap["锡山人民医院"] = "无锡市锡山人民医院"

	monthMap = make(map[int]string)
	monthMap[1] = "C"
	monthMap[2] = "D"
	monthMap[3] = "E"
	monthMap[4] = "F"
	monthMap[5] = "G"
	monthMap[6] = "H"
}

func main() {

	// read count excel
	monthHospitalList := getAllCounterHospital()

	// get un-name hospital
	for i, hospital := range monthHospitalList {
		if len(hospital.formalName) < 1 {
			fmt.Println(i+2, hospital.name, hospital.month)
		}
	}

	// loop data excel
	write2DataExcel(monthHospitalList)
}

func write2DataExcel(monthHospitalList []counterHospital) {
	f, err := excelize.OpenFile("gulou-excel/hospitalData.xlsx")
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		formalName := row[0]
		for _, hospital := range monthHospitalList {
			if formalName == hospital.formalName {
				m, ok := monthMap[hospital.month]
				if !ok {
					continue
				}
				axis := m + strconv.Itoa(i+1)
				err := f.SetCellValue("Sheet1", axis, "√")
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if err := f.SaveAs("gulou-excel/hospitalData.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func getAllCounterHospital() []counterHospital {

	f, err := excelize.OpenFile("gulou-excel/counter.xlsx")
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}
	var counterList = make([]counterHospital, 0)
	for _, row := range rows {
		name := row[0]
		if len(name) < 3 {
			continue
		}
		month, _ := strconv.Atoi(row[1])
		matchName, same := findBeatMatch(name)
		counter := counterHospital{
			name:       name,
			formalName: matchName,
			month:      month,
			isSameName: same,
		}
		counterList = append(counterList, counter)
	}
	return counterList
}

func findBeatMatch(name string) (string, bool) {
	formal, ok := hospitalNameMap[name]
	if !ok {
		formal = strsim.FindBestMatchOne(name, hospitalList).S
	}
	return formal, name == formal
}

type counterHospital struct {
	name       string
	formalName string
	month      int
	isSameName bool
}
