package excel

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
	"strings"
)

func init() {
	mf = make(map[string]func(value reflect.Value) string)
	mf["abc"] = func(value reflect.Value) string {
		if value.Kind() == reflect.String {
			i := value.Interface().(string)
			return i + " f"
		}
		return ""
	}
}

func SliceToExcel(a any) (*excelize.File, error) {
	value := reflect.ValueOf(a)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("请传入slice")
	}
	vl := value.Len()
	position := make([]int, 0)
	title := make([]string, 0)
	format := make([]string, 0)
	width := make([]float64, 0)

	// 获取切片中的结构体类型
	t := reflect.TypeOf(a).Elem()
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// 遍历结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		g := field.Tag.Get("bze")
		atoi := 0
		var err error
		if g != "" {
			split := strings.Split(g, ",")
			atoi, err = strconv.Atoi(split[0])
			if err != nil {
				return nil, err
			}
			if atoi < 0 {
				return nil, errors.New("必须大于0")
			}
			if len(title) < atoi {
				ad := atoi - len(title)
				position = append(position, make([]int, ad)...)
				title = append(title, make([]string, ad)...)
				width = append(width, make([]float64, ad)...)
				format = append(format, make([]string, ad)...)
			}
			title[atoi-1] = split[1]
			position[atoi-1] = i
			if len(split) > 2 {
				err = setExcelType(split, atoi, width, format)
				if err != nil {
					return nil, err
				}
			}
		}

	}
	return createExcel(position, width, title, format, vl, value)
}

func setExcelType(split []string, atoi int, width []float64, format []string) error {
	for ii := 2; ii < len(split); ii++ {
		sp := strings.Split(split[ii], "=")
		if len(sp) != 2 {
			return errors.New("tag格式错误")
		}
		switch sp[0] {
		case "width":
			i2, err := strconv.ParseFloat(sp[1], 64)
			if err != nil {
				return err
			}
			width[atoi-1] = i2
		case "format":
			format[atoi-1] = sp[1]
		}
	}
	return nil
}
func createExcel(position []int, width []float64, title, format []string, vl int, value reflect.Value) (*excelize.File, error) {
	f := excelize.NewFile()
	for i, i2 := range width {
		if i2 != 0 {
			column := toExcelColumn(i + 1)
			err := f.SetColWidth("Sheet1", column, column, i2)
			if err != nil {
				return nil, err
			}
		}
	}
	err := f.SetSheetRow("Sheet1", "A1", &title)
	if err != nil {
		return nil, err
	}
	ti := len(title)
	for i := 0; i < vl; i++ {
		item := value.Index(i)
		ls := setExcelContent(item, position, ti, format)
		err = f.SetSheetRow("Sheet1", "A"+strconv.Itoa(i+2), &ls)
		if err != nil {
			return nil, err
		}
	}
	return f, err
}

func setExcelContent(item reflect.Value, position []int, lt int, format []string) []string {

	// 如果元素类型为指针，使用Elem()获取指针指向的值
	for item.Kind() == reflect.Ptr {
		item = item.Elem()
	}
	s2 := make([]string, lt)
	for i2, i3 := range position {
		index := item.Field(i3)
		if format[i2] != "" {
			s2[i2] = mf[format[i2]](index)
		} else {
			s2[i2] = fmt.Sprint(index)
		}
	}
	return s2
}

var mf map[string]func(value reflect.Value) string

func toExcelColumn(num int) string {
	column := ""
	for num > 0 {
		num-- // This is because there is no column with '0'
		column = string('A'+num%26) + column
		num /= 26
	}
	return column
}
