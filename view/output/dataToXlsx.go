package output

import (
	"fmt"

	eksresourcetypes "k8s-informer/aws/types"
	k8sresourcetypes "k8s-informer/k8s/types"
	clusterresource "k8s-informer/k8s/types/cluster-resource"
	"k8s-informer/k8s/types/workload"
	"reflect"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	xlsx "github.com/tealeg/xlsx/v3"
)

func WriteIndexSheet(sheet *xlsx.Sheet, indexDataMap linkedhashmap.Map) {

	headerRow := sheet.AddRow()
	headerRow.SetHeight(50)
	headerCell := headerRow.AddCell()
	headerCell.SetValue("Index")
	headerStyle := xlsx.NewStyle()
	headerStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	headerStyle.Fill = *xlsx.NewFill("solid", "92CDDC", "92CDDC")
	headerStyle.Alignment = xlsx.Alignment{Horizontal: "center", Vertical: "center"}
	headerStyle.Font = xlsx.Font{Size: 26.0, Name: "Arial", Bold: true, Color: "000000", Italic: false}
	headerCell.SetStyle(headerStyle)

	iter := indexDataMap.Iterator()

	for iter.Next() {
		key, value := iter.Key(), iter.Value()

		setIndexData(sheet, key.(string), value.([]string))
	}

	sheet.SetColWidth(1, 1, 100.0)

}

func setIndexData(sheet *xlsx.Sheet, title string, resources []string) {
	titleRow := sheet.AddRow()
	titleRow.SetHeight(40.0)
	titleCell := titleRow.AddCell()
	titleCell.SetValue(title)
	titleStyle := xlsx.NewStyle()
	titleStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	titleStyle.Fill = *xlsx.NewFill("solid", "B7DEE8", "B7DEE8")
	titleStyle.Alignment = xlsx.Alignment{Horizontal: "center", Vertical: "center"}
	titleStyle.Font = xlsx.Font{Size: 18.0, Name: "Arial", Bold: true, Color: "000000", Italic: false}
	titleCell.SetStyle(titleStyle)

	dataStyle := xlsx.NewStyle()
	dataStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	dataStyle.Font = xlsx.Font{Size: 18.0, Name: "Arial", Bold: false, Color: "000000", Italic: false}
	dataStyle.Alignment = xlsx.Alignment{Horizontal: "left", Vertical: "center"}

	for index, resource := range resources {
		dataRow := sheet.AddRow()
		dataRow.SetHeight(40.0)
		dataCell := dataRow.AddCell()

		dataCell.SetValue(fmt.Sprintf("%d.%s", index+1, resource))
		dataCell.SetStyle(dataStyle)
	}
}

func WriteResourceData[T eksresourcetypes.ResourceType | k8sresourcetypes.ResourceType](file *xlsx.File, sheetName string, resourceData []T, resourceType reflect.Type) {

	// Sheet 생성
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		panic(err)
	}

	// Struct에서 헤더 추출
	fields, err := extractStructTags(resourceType)
	if err != nil {
		panic(err)
	}

	// 헤더행 생성 후 추출한 헤더 삽입
	headerRow := sheet.AddRow()
	insertHeaders(headerRow, fields)

	// 데이터 저장
	for _, resource := range resourceData {
		setData(resource, fields, sheet, xlsx.Row{}, 0, false)
	}

	// Column 넓이 자동 맞춤 적용
	for i := 1; i <= sheet.MaxCol; i++ {
		sheet.SetColAutoWidth(i, xlsx.DefaultAutoWidth)
	}

	// 마지막 Column letter 측정
	maxCell := numberToColumn(sheet.MaxCol, sheet.MaxRow)

	// 필터 적용
	sheet.AutoFilter = &xlsx.AutoFilter{
		TopLeftCell:     "A1",
		BottomRightCell: maxCell,
	}

	// Column 테두리 적용
	colStyle := xlsx.NewStyle()
	colStyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")

	sheet.Cols.ForEach(func(idx int, col *xlsx.Col) {
		col.SetStyle(colStyle)
	})
}

func numberToColumn(col int, row int) string {
	var result strings.Builder
	for col > 0 {
		col--
		result.WriteByte(byte(col%26) + 'A')
		col /= 26
	}
	return reverseString(result.String()) + strconv.Itoa(row)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func insertHeaders(headerRow *xlsx.Row, fields linkedhashmap.Map) {
	style := xlsx.NewStyle()
	style.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Fill = *xlsx.NewFill("solid", "B7DEE8", "B7DEE8")
	style.Font = xlsx.Font{
		Size:   12.0,
		Name:   "Arial",
		Bold:   true,
		Color:  "000000",
		Italic: false,
	}
	style.Alignment = xlsx.Alignment{
		Horizontal: "center",
	}
	for _, header := range fields.Values() {
		headerType := reflect.TypeOf(header)
		if headerType == reflect.TypeOf("string") {
			cell := headerRow.AddCell()
			cell.SetStyle(style)
			cell.SetValue(header)

		} else if headerType == reflect.TypeOf(linkedhashmap.Map{}) {
			childFields := header.(linkedhashmap.Map)
			insertHeaders(headerRow, childFields)
		}
	}
}

func setData[T eksresourcetypes.ResourceType | k8sresourcetypes.ResourceType](resource T, fields linkedhashmap.Map, sheet *xlsx.Sheet, parentRow xlsx.Row, depth int, hasChild bool) {
	iterator := fields.Iterator()

	dataRow := sheet.AddRow()
	l := sheet.MaxRow - 1

	if depth > 0 {

		parentRow.ForEachCell(func(cell *xlsx.Cell) error {
			dataRow.PushCell(cell)
			return nil
		}, xlsx.SkipEmptyCells)
	}

	for iterator.Next() {
		key, value := iterator.Key(), iterator.Value()

		fieldValue := reflect.ValueOf(resource).FieldByName(key.(string))
		if !fieldValue.IsValid() {
			continue
		}

		if reflect.TypeOf(value) == reflect.TypeOf("string") {
			dataRow.AddCell().SetValue(fieldValue)
			continue
		} else if reflect.TypeOf(value) == reflect.TypeOf(linkedhashmap.Map{}) {
			hasChild = true
			childResources := fieldValue.Interface()
			childFields := value.(linkedhashmap.Map)
			sliceType := fieldValue.Type()

			if sliceType == reflect.TypeOf([]workload.ContainerInfo{}) {
				childSlice := childResources.([]workload.ContainerInfo)
				for i := 0; i < len(childSlice); i++ {
					setData(childSlice[i], childFields, sheet, *dataRow, depth+1, false)
				}
			} else if sliceType == reflect.TypeOf([]clusterresource.NodeInfo{}) {
				childSlice := childResources.([]clusterresource.NodeInfo)
				for i := 0; i < len(childSlice); i++ {
					setData(childSlice[i], childFields, sheet, *dataRow, depth+1, false)
				}
			}
			continue
		}
	}

	if hasChild {
		sheet.RemoveRowAtIndex(l)
	}
}

func extractStructTags(t reflect.Type) (linkedhashmap.Map, error) {
	fields := linkedhashmap.New()

	structType := t

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tagValue := field.Tag.Get("xlsx")

		if tagValue == "" {
			continue
		}

		if tagValue == "CHILD" {
			subStructType := field.Type.Elem()
			subFields, err := extractStructTags(subStructType)
			if err != nil {
				return linkedhashmap.Map{}, err
			}

			// 현재 함수의 결과에 서브 헤더 추가
			fields.Put(field.Name, subFields)
			continue
		}
		fields.Put(field.Name, tagValue)
	}

	return *fields, nil
}
