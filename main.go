package main
/**
 * 基于excelize的Excel操作
 */
import (
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/buger/jsonparser"
    "fmt"
    "os"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
    "flag"
    "regexp"
    "strconv"
)
var task_now = 0 // 已处理任务个数
func main () {

    var EXCEL_PATH string   // 操作的Excel文件绝对路径
    var IMGS_JSON string    // 需要添加的图片json

    // 获取传入参数
    flag.StringVar(&EXCEL_PATH, "f", "./excel.xlsx", "excel文件绝对路径")
    flag.StringVar(&IMGS_JSON, "i", "", "需要插入图片json")
    flag.Parse()

    if len(IMGS_JSON) == 0 {
        fmt.Println("请传入-i参数")
        os.Exit(1)
    }

    file, err := excelize.OpenFile(EXCEL_PATH) // 打开excel
    throwErr(err, "Excel文件不存在!") // 检测打开失败

    var firstSheetName string = file.GetSheetMap()[1] // 工作薄名称

    var addImageParam AddImageParamsStruct // 需传递参数
    addImageParam.fileHandle = file
    addImageParam.firstSheetName = firstSheetName
    finishNoticeHandle := make(chan int) // 全部完成通知管道
    total := 0 // 任务总数
    jsonparser.ArrayEach([]byte(IMGS_JSON), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
        total++;
    })
    // 图像添加处理
    jsonparser.ArrayEach([]byte(IMGS_JSON), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
        pathByte,_,_,errPath:=jsonparser.Get(value, "path") ; path := string(pathByte) // 图片绝对路径
        posByte,_,_,errPos:=jsonparser.Get(value, "pos") ; pos := string(posByte) // 图片位置
        heightByte,_,_,_:=jsonparser.Get(value, "height") ; height := string(heightByte) // 单元格高度
        widthByte,_,_,_:=jsonparser.Get(value, "width") ; width := string(widthByte) // 单元格高度
        xValue,_ := jsonparser.GetInt(value, "x")
        yValue,_ := jsonparser.GetInt(value, "y")
        // 判断解析是否成功
        if errPath != nil || errPos != nil {
            return
        }
        addImageParam.path = path
        addImageParam.pos = pos
        addImageParam.height = height
        addImageParam.width = width
        addImageParam.total_number = total
        addImageParam.x = xValue
        addImageParam.y = yValue
        go addImage(addImageParam, finishNoticeHandle)
    })
    <-finishNoticeHandle // 等待任务完成

    // 保存图像
    errSave := file.Save()
    throwErr(errSave, "Excel文件保存失败!") // 检测打开失败

}

// excel添加图片操作
func addImage(info AddImageParamsStruct, finishNotice chan int) (string, error) {
    defer (func() {
       task_now++
       if task_now == info.total_number {
           finishNotice<-1
       }
    })()

    width := info.width; pos := info.pos; height := info.height; path := info.path; x := info.x; y := info.y // 宽度、位置、高度、图片地址
    firstSheetName := info.firstSheetName  // 工作薄名称
    fileHandle := info.fileHandle  // excel操作句柄
    rowMatch := regexp.MustCompile(`[a-zA-Z]+([0-9]+)`).FindAllStringSubmatch(pos, -1) // 行数
    colMatch := regexp.MustCompile(`([a-zA-Z])+[0-9]+`).FindAllStringSubmatch(pos, -1) // 列数

    // 设置行高
    rowNumString := rowMatch[0][1]
    rowNum, errRowNum := strconv.Atoi(rowNumString) // 行号
    heightNum, errHeightNum := strconv.ParseFloat(height, 64) // 行高
    if errRowNum == nil && errHeightNum == nil {
        fileHandle.SetRowHeight(firstSheetName, rowNum, heightNum)
    } else {
        return "",&myError{"行高设置失败，"+"rowNum："+strconv.Itoa(rowNum)+"，heightNum："+strconv.FormatFloat(heightNum, 'E', -1, 64)}
    }
    // 设置列宽
    colNum := colMatch[0][1] // 列号
    if widthNum, errWidthNum := strconv.ParseFloat(width, 64) ; errWidthNum == nil {
        fileHandle.SetColWidth(firstSheetName, colNum, colNum, widthNum)
    } else {
        if len(width) != 0 {
            return "",&myError{"行宽设置失败，pos："+pos+"，colNum："+colNum+"，width："+width}
        }
    }

    // 添加图片
    if errAddPic := fileHandle.AddPicture(firstSheetName, pos, path, `{ "x_scale": 1,"y_scale": 1,"positioning": "absolute","x_offset": `+strconv.FormatInt(x, 10)+`,"y_offset": `+strconv.FormatInt(y, 10)+`}`) ; errAddPic != nil {
        return "",&myError{"图片添加失败，"+"path："+path+"，pos："+pos}
    }

    return "完成", nil
}

// 管道信息
type BaseExcelInfoStruct struct {
    total_number int  // 任务总数
}

// excel文件信息 结构体
type ExcelInfoStruct struct {
    BaseExcelInfoStruct
    fileHandle *excelize.File
    firstSheetName string
}

// addImage函数参数 结构体
type AddImageParamsStruct struct {
    ExcelInfoStruct
    path string
    pos string
    height string
    width string
    x int64     // 左边距
    y int64     // 上边距
}

// 抛出错误并停止运行
func throwErr(err error, msg string){
    if err != nil {
        fmt.Println(msg)
        fmt.Println(err)
        os.Exit(1)
    }
}

// 错误抛出
type myError struct {
    message string
}

func (e *myError) Error () string {
    return e.message
}
