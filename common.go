package main
import (
    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "fmt"
    "os"
)
// 管道信息
type BaseExcelInfoStruct struct {

}

// excel文件信息 结构体
type ExcelInfoStruct struct {
    BaseExcelInfoStruct
    fileHandle *excelize.File
    firstSheetName string
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