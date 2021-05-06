package main
import (
    "fmt"
    "regexp"
    "strconv"
)

// addImage函数参数
type AddImageParams struct {
    ExcelInfoStruct
    path string
    pos string
    height string
    width string
    x int64     // 左边距
    y int64     // 上边距
}

// 参数验证
func (a *AddImageParams) valid () bool {
    if len(a.path) == 0 || len(a.pos) == 0 {
        return false
    }
    return true
}

// excel添加图片操作
func addImageHandle(info AddImageParams, finishNotice chan int) {
    defer (func() {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
        task_now++
        if task_now == total_number {
            finishNotice<-1
        }
    })()

    var (
        width = info.width  // 宽度
        pos = info.pos  // 位置
        height = info.height // 高度
        path = info.path // 图片地址
        x = info.x
        y = info.y
        firstSheetName = info.firstSheetName  // 工作薄名称
        fileHandle = info.fileHandle  // excel操作句柄
        rowMatch = regexp.MustCompile(`[a-zA-Z]+([0-9]+)`).FindAllStringSubmatch(pos, -1) // 行数

        rowNum, errRowNum = strconv.Atoi(rowMatch[0][1]) // 行号
        heightNum, errHeightNum = strconv.ParseFloat(height, 64) // 行高

        colMatch = regexp.MustCompile(`([a-zA-Z])+[0-9]+`).FindAllStringSubmatch(pos, -1) // 列数
        colNum = colMatch[0][1] // 列号
    )

    // 设置行高
    if errRowNum == nil && errHeightNum == nil {
        fileHandle.SetRowHeight(firstSheetName, rowNum, heightNum)
    } else {
        panic("行高设置失败，"+"rowNum："+strconv.Itoa(rowNum)+"，heightNum："+strconv.FormatFloat(heightNum, 'E', -1, 64))
    }

    // 设置列宽
    if widthNum, errWidthNum := strconv.ParseFloat(width, 64) ; errWidthNum == nil {
        fileHandle.SetColWidth(firstSheetName, colNum, colNum, widthNum)
    } else {
        if len(width) != 0 {
            panic("行宽设置失败，pos："+pos+"，colNum："+colNum+"，width："+width)
        }
    }

    // 添加图片
    if errAddPic := fileHandle.AddPicture(firstSheetName, pos, path, `{ "x_scale": 1,"y_scale": 1,"positioning": "absolute","x_offset": `+strconv.FormatInt(x, 10)+`,"y_offset": `+strconv.FormatInt(y, 10)+`}`) ; errAddPic != nil {
        panic("图片添加失败，"+"path："+path+"，pos："+pos)
    }

}

